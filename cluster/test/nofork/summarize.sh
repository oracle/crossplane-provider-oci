#!/usr/bin/env bash

set -euo pipefail

if (($# != 1)); then
  echo "usage: summarize.sh RUN_DIR" >&2
  exit 2
fi

run_dir=$1
resources="$run_dir/pod-resources.tsv"
status="$run_dir/pod-status.tsv"
metrics="$run_dir/metrics/provider-metrics.prom"

cat <<EOF
# No-fork test summary

- Run directory: \`$run_dir\`
- Generated: $(date -u +%Y-%m-%dT%H:%M:%SZ)

## Peak provider container usage by phase

| Phase | Provider | Peak CPU (cores) | Peak memory (MiB) |
|---|---|---:|---:|
EOF

if [[ -s $resources ]]; then
  awk -F '\t' '
    function cpu(v) {
      if (v ~ /n$/) {sub(/n$/, "", v); return v / 1000000000}
      if (v ~ /u$/) {sub(/u$/, "", v); return v / 1000000}
      if (v ~ /m$/) {sub(/m$/, "", v); return v / 1000}
      return v + 0
    }
    function mem(v) {
      if (v ~ /Ki$/) {sub(/Ki$/, "", v); return v / 1024}
      if (v ~ /Mi$/) {sub(/Mi$/, "", v); return v + 0}
      if (v ~ /Gi$/) {sub(/Gi$/, "", v); return v * 1024}
      if (v ~ /K$/)  {sub(/K$/, "", v); return v / 1024}
      if (v ~ /M$/)  {sub(/M$/, "", v); return v + 0}
      if (v ~ /G$/)  {sub(/G$/, "", v); return v * 1024}
      return v / 1048576
    }
    function provider(pod) {
      if (pod ~ /provider-family-oci/) return "family"
      if (pod ~ /provider-oci-/) {
        sub(/^.*provider-oci-/, "", pod)
        sub(/(-[0-9a-f]+)+-[0-9a-z]+$/, "", pod)
        return pod
      }
      return "other"
    }
    NR > 1 {
      key=$2 SUBSEP provider($3)
      c=cpu($5); m=mem($6)
      if (c > maxcpu[key]) maxcpu[key]=c
      if (m > maxmem[key]) maxmem[key]=m
      seen[key]=1
    }
    END {
      for (key in seen) {
        split(key, parts, SUBSEP)
        printf "| %s | %s | %.3f | %.1f |\n", parts[1], parts[2], maxcpu[key], maxmem[key]
      }
    }
  ' "$resources" | sort
else
  echo "| No Metrics Server samples collected | - | - | - |"
fi

cat <<'EOF'

## Peak provider container usage

| Pod/container | Peak CPU (cores) | Peak memory (MiB) |
|---|---:|---:|
EOF

if [[ -s $resources ]]; then
  awk -F '\t' '
    function cpu(v) {
      if (v ~ /n$/) {sub(/n$/, "", v); return v / 1000000000}
      if (v ~ /u$/) {sub(/u$/, "", v); return v / 1000000}
      if (v ~ /m$/) {sub(/m$/, "", v); return v / 1000}
      return v + 0
    }
    function mem(v) {
      if (v ~ /Ki$/) {sub(/Ki$/, "", v); return v / 1024}
      if (v ~ /Mi$/) {sub(/Mi$/, "", v); return v + 0}
      if (v ~ /Gi$/) {sub(/Gi$/, "", v); return v * 1024}
      if (v ~ /K$/)  {sub(/K$/, "", v); return v / 1024}
      if (v ~ /M$/)  {sub(/M$/, "", v); return v + 0}
      if (v ~ /G$/)  {sub(/G$/, "", v); return v * 1024}
      return v / 1048576
    }
    NR > 1 {
      key=$3 "/" $4
      c=cpu($5); m=mem($6)
      if (c > maxcpu[key]) maxcpu[key]=c
      if (m > maxmem[key]) maxmem[key]=m
      seen[key]=1
    }
    END {
      for (key in seen) printf "| %s | %.3f | %.1f |\n", key, maxcpu[key], maxmem[key]
    }
  ' "$resources" | sort
else
  echo "| No Metrics Server samples collected | - | - |"
fi

cat <<'EOF'

## Pod stability

| Pod | Maximum observed restarts |
|---|---:|
EOF

if [[ -s $status ]]; then
  awk -F '\t' '
    NR > 1 {
      total=0
      value=$6
      gsub(/[\[\]]/, "", value)
      count=split(value, parts, /,/)
      for (i=1; i<=count; i++) total += parts[i] + 0
      if (total > restarts[$3]) restarts[$3]=total
      seen[$3]=1
    }
    END {for (pod in seen) printf "| %s | %d |\n", pod, restarts[pod]}
  ' "$status" | sort
else
  echo "| No pod status samples collected | - |"
fi

cat <<'EOF'

## Final no-fork cache metrics

The following values are the last collected sample for each provider pod. They
exclude ProviderConfig names and credential data.

```text
EOF

if [[ -s $metrics ]] && grep -q '^provider_oci_nofork_' "$metrics"; then
  awk '
    /^# sample / {
      pod="unknown"
      for (i=1; i<=NF; i++) if ($i ~ /^pod=/) {pod=$i; sub(/^pod=/, "", pod)}
      next
    }
    /^provider_oci_nofork_/ {last[pod SUBSEP $1]=pod " " $0}
    END {for (key in last) print last[key]}
  ' "$metrics" | sort
else
  echo "No no-fork cache metrics were collected. The running provider image may predate these metrics."
fi

cat <<'EOF'
```

## Evidence files

- `phases.tsv`: test phase timeline
- `pod-resources.tsv`: timestamped CPU and memory samples
- `pod-status.tsv`: readiness, restarts and termination reasons
- `metrics/provider-metrics.prom`: filtered provider Prometheus samples
- `snapshots/managed-resources.log`: managed-resource status snapshots
- `snapshots/processes.log`: best-effort process evidence
- `events.log`: namespace events
- `logs/`: provider logs, previous logs and pod descriptions

Review and sanitize all files before sharing outside the test environment.
EOF
