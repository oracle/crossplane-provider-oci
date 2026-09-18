#!/usr/bin/env bash

set -euo pipefail

KUBECTL=${KUBECTL:-kubectl}
NAMESPACE=${NAMESPACE:-crossplane-system}
INTERVAL_SECONDS=${INTERVAL_SECONDS:-10}
DURATION_SECONDS=${DURATION_SECONDS:-0}
POD_PATTERN=${POD_PATTERN:-'provider.*oci|oci.*provider'}
OUTPUT_ROOT=${OUTPUT_ROOT:-test-results}
RUN_NAME=${RUN_NAME:-nofork}

usage() {
  cat <<'EOF'
Usage: collect.sh [options]

Collect provider resource usage, lifecycle state, Prometheus metrics, events,
logs, and no-fork execution evidence until interrupted.

Options:
  --run-name NAME       Run name used in the output directory.
  --output-root DIR     Parent directory for results (default: test-results).
  --namespace NAME      Provider pod namespace (default: crossplane-system).
  --interval SECONDS    Sample interval (default: 10).
  --duration SECONDS    Stop after this duration; 0 waits for a signal.
  --pod-pattern REGEX   Extended regex used to select provider pod names.
  -h, --help            Show this help.

Environment variables with the uppercase option names are also supported.
Use mark-phase.sh with the printed run directory to label test phases.
EOF
}

while (($# > 0)); do
  case "$1" in
    --run-name) RUN_NAME=$2; shift 2 ;;
    --output-root) OUTPUT_ROOT=$2; shift 2 ;;
    --namespace) NAMESPACE=$2; shift 2 ;;
    --interval) INTERVAL_SECONDS=$2; shift 2 ;;
    --duration) DURATION_SECONDS=$2; shift 2 ;;
    --pod-pattern) POD_PATTERN=$2; shift 2 ;;
    -h|--help) usage; exit 0 ;;
    *) echo "unknown argument: $1" >&2; usage >&2; exit 2 ;;
  esac
done

if ! [[ $INTERVAL_SECONDS =~ ^[1-9][0-9]*$ ]]; then
  echo "--interval must be a positive integer" >&2
  exit 2
fi
if ! [[ $DURATION_SECONDS =~ ^[0-9]+$ ]]; then
  echo "--duration must be a non-negative integer" >&2
  exit 2
fi
command -v "$KUBECTL" >/dev/null 2>&1 || {
  echo "kubectl command not found: $KUBECTL" >&2
  exit 1
}

timestamp=$(date -u +%Y%m%dT%H%M%SZ)
safe_run_name=$(printf '%s' "$RUN_NAME" | tr -cs '[:alnum:]._-' '-')
RUN_DIR=${RUN_DIR:-"$OUTPUT_ROOT/${safe_run_name}-${timestamp}"}
mkdir -p "$RUN_DIR/metrics" "$RUN_DIR/logs" "$RUN_DIR/snapshots"
PHASE_FILE="$RUN_DIR/current-phase"
printf 'baseline\n' >"$PHASE_FILE"

RESOURCE_SAMPLES="$RUN_DIR/pod-resources.tsv"
STATUS_SAMPLES="$RUN_DIR/pod-status.tsv"
METRIC_SAMPLES="$RUN_DIR/metrics/provider-metrics.prom"
PHASES="$RUN_DIR/phases.tsv"
ERRORS="$RUN_DIR/collector-errors.log"

printf 'timestamp\tphase\tpod\tcontainer\tcpu\tmemory\n' >"$RESOURCE_SAMPLES"
printf 'timestamp\tphase\tpod\tready\tstatus\trestarts\twaiting_reasons\tterminated_reasons\tnode\n' >"$STATUS_SAMPLES"
printf 'timestamp\tphase\n' >"$PHASES"
printf '%s\tbaseline\n' "$(date -u +%Y-%m-%dT%H:%M:%SZ)" >>"$PHASES"
: >"$METRIC_SAMPLES"
: >"$ERRORS"

list_provider_pods() {
  "$KUBECTL" -n "$NAMESPACE" get pods --no-headers \
    -o custom-columns='NAME:.metadata.name' 2>>"$ERRORS" | grep -E "$POD_PATTERN" || true
}

record_metadata() {
  {
    echo "run_name=$RUN_NAME"
    echo "started_at=$(date -u +%Y-%m-%dT%H:%M:%SZ)"
    echo "namespace=$NAMESPACE"
    echo "interval_seconds=$INTERVAL_SECONDS"
    echo "duration_seconds=$DURATION_SECONDS"
    echo "pod_pattern=$POD_PATTERN"
    echo "kubectl_context=$($KUBECTL config current-context 2>/dev/null || echo unavailable)"
    if git rev-parse --is-inside-work-tree >/dev/null 2>&1; then
      echo "git_commit=$(git rev-parse HEAD)"
      echo "git_branch=$(git branch --show-current)"
      echo "git_dirty=$(test -n "$(git status --porcelain)" && echo true || echo false)"
    fi
  } >"$RUN_DIR/run.env"
  "$KUBECTL" version -o yaml >"$RUN_DIR/kubernetes-version.yaml" 2>>"$ERRORS" || true
  "$KUBECTL" get providers.pkg.crossplane.io -o wide >"$RUN_DIR/providers-start.txt" 2>>"$ERRORS" || true
  "$KUBECTL" -n "$NAMESPACE" get pods -o wide >"$RUN_DIR/pods-start.txt" 2>>"$ERRORS" || true
}

sample_resources() {
	local now=$1 phase=$2
	"$KUBECTL" -n "$NAMESPACE" top pod --containers --no-headers 2>>"$ERRORS" |
		grep -E "$POD_PATTERN" |
		awk -v ts="$now" -v phase="$phase" 'BEGIN {OFS="\t"} {print ts,phase,$1,$2,$3,$4}' \
      >>"$RESOURCE_SAMPLES" || true
}

sample_status() {
  local now=$1 phase=$2
  "$KUBECTL" -n "$NAMESPACE" get pods --no-headers \
    -o custom-columns='NAME:.metadata.name,READY:.status.containerStatuses[*].ready,STATUS:.status.phase,RESTARTS:.status.containerStatuses[*].restartCount,WAITING:.status.containerStatuses[*].state.waiting.reason,TERMINATED:.status.containerStatuses[*].lastState.terminated.reason,NODE:.spec.nodeName' \
    2>>"$ERRORS" | grep -E "$POD_PATTERN" |
    awk -v ts="$now" -v phase="$phase" 'BEGIN {OFS="\t"} {print ts,phase,$1,$2,$3,$4,$5,$6,$7}' \
      >>"$STATUS_SAMPLES" || true
}

sample_metrics() {
  local now=$1 phase=$2 pod metrics
  while IFS= read -r pod; do
    [[ -n $pod ]] || continue
    if metrics=$($KUBECTL get --raw "/api/v1/namespaces/$NAMESPACE/pods/$pod:8080/proxy/metrics" 2>>"$ERRORS"); then
      {
        printf '# sample timestamp=%s phase=%s pod=%s\n' "$now" "$phase" "$pod"
        printf '%s\n' "$metrics" | grep -E \
          '^(provider_oci_nofork_|crossplane_|upjet_|controller_runtime_reconcile_|workqueue_|go_memstats_|go_goroutines|process_)' || true
      } >>"$METRIC_SAMPLES"
    else
      printf '%s metrics scrape failed for %s/%s\n' "$now" "$NAMESPACE" "$pod" >>"$ERRORS"
    fi
  done < <(list_provider_pods)
}

sample_managed_resources() {
	local now=$1 phase=$2 output
	{
		printf '# sample timestamp=%s phase=%s\n' "$now" "$phase"
		if output=$("$KUBECTL" get managed -A 2>&1); then
			printf '%s\n' "$output"
		else
			printf '%s managed-resource snapshot failed: %s\n' "$now" "$output" >>"$ERRORS"
		fi
	} >>"$RUN_DIR/snapshots/managed-resources.log"
}

sample_processes() {
  local now=$1 phase=$2 pod
  while IFS= read -r pod; do
    [[ -n $pod ]] || continue
    {
      printf '# sample timestamp=%s phase=%s pod=%s\n' "$now" "$phase" "$pod"
      "$KUBECTL" -n "$NAMESPACE" exec "$pod" -- sh -c \
        'for f in /proc/[0-9]*/cmdline; do tr "\000" " " < "$f" 2>/dev/null; echo; done' 2>&1 || true
    } >>"$RUN_DIR/snapshots/processes.log"
  done < <(list_provider_pods)
}

capture_final_evidence() {
	local pod restart_counts output
	"$KUBECTL" get providers.pkg.crossplane.io -o wide >"$RUN_DIR/providers-final.txt" 2>>"$ERRORS" || true
	"$KUBECTL" -n "$NAMESPACE" get pods -o wide >"$RUN_DIR/pods-final.txt" 2>>"$ERRORS" || true
	"$KUBECTL" -n "$NAMESPACE" get events --sort-by=.lastTimestamp >"$RUN_DIR/events.log" 2>/dev/null || true
	if output=$("$KUBECTL" get managed -A 2>&1); then
		printf '%s\n' "$output" >"$RUN_DIR/managed-resources-final.txt"
	else
		printf '%s final managed-resource snapshot failed: %s\n' "$(date -u +%Y-%m-%dT%H:%M:%SZ)" "$output" >>"$ERRORS"
	fi

  while IFS= read -r pod; do
    [[ -n $pod ]] || continue
		"$KUBECTL" -n "$NAMESPACE" describe pod "$pod" >"$RUN_DIR/logs/$pod.describe.log" 2>>"$ERRORS" || true
		"$KUBECTL" -n "$NAMESPACE" logs "$pod" --all-containers >"$RUN_DIR/logs/$pod.log" 2>>"$ERRORS" || true
		restart_counts=$("$KUBECTL" -n "$NAMESPACE" get pod "$pod" -o jsonpath='{.status.containerStatuses[*].restartCount}' 2>>"$ERRORS" || true)
		if [[ $restart_counts =~ (^|[[:space:]])[1-9][0-9]*($|[[:space:]]) ]]; then
			"$KUBECTL" -n "$NAMESPACE" logs "$pod" --all-containers --previous >"$RUN_DIR/logs/$pod.previous.log" 2>>"$ERRORS" || true
		fi
	done < <(list_provider_pods)

  if [[ -x $(dirname "$0")/summarize.sh ]]; then
    "$(dirname "$0")/summarize.sh" "$RUN_DIR" >"$RUN_DIR/summary.md" 2>>"$ERRORS" || true
  fi
}

stopping=false
cleanup() {
  local exit_status=${1:-$?}
  if [[ $stopping == true ]]; then
    exit "$exit_status"
  fi
  stopping=true
  trap - INT TERM EXIT
  echo "capturing final evidence in $RUN_DIR" >&2
  capture_final_evidence
  echo "$RUN_DIR"
  exit "$exit_status"
}
trap 'cleanup 0' INT TERM
trap 'cleanup $?' EXIT

record_metadata
echo "collector output: $RUN_DIR" >&2
echo "mark phases with: $(dirname "$0")/mark-phase.sh '$RUN_DIR' <phase>" >&2

start_epoch=$(date +%s)
sample_number=0
while true; do
  now=$(date -u +%Y-%m-%dT%H:%M:%SZ)
  phase=$(tr -d '\r\n' <"$PHASE_FILE")
  sample_resources "$now" "$phase"
  sample_status "$now" "$phase"
  sample_metrics "$now" "$phase"

  if ((sample_number % 6 == 0)); then
    sample_managed_resources "$now" "$phase"
    sample_processes "$now" "$phase"
  fi
  sample_number=$((sample_number + 1))

  if ((DURATION_SECONDS > 0 && $(date +%s) - start_epoch >= DURATION_SECONDS)); then
    break
  fi
  sleep "$INTERVAL_SECONDS"
done
