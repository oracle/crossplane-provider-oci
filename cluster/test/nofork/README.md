# No-fork metrics collector

This directory contains a reusable test collector for correctness,
compatibility, and scale validation of the in-process OCI provider runtime.
It does not create or delete OCI resources.

## Prerequisites

- A current `kubectl` context for the test cluster.
- Metrics Server for pod CPU and memory (`kubectl top pod` must work).
- Provider service pods exposing their default metrics endpoint on port 8080.
- Optional: Prometheus Operator when using `podmonitor.yaml`.

Apply `runtime-config.yaml` and reference `oci-nofork-metrics` from every OCI
family and service `Provider` under test. Merge any authentication-specific
runtime configuration, such as Workload Identity environment variables or
service accounts, into the same `DeploymentRuntimeConfig`.

## Run the collector

From the repository root:

```bash
chmod +x cluster/test/nofork/*.sh

RUN_DIR="$PWD/test-results/nofork-smoke-$(date -u +%Y%m%dT%H%M%SZ)"
RUN_DIR="$RUN_DIR" \
  cluster/test/nofork/collect.sh \
  --run-name nofork-smoke \
  --namespace crossplane-system \
  --interval 10 &
COLLECTOR_PID=$!

cluster/test/nofork/mark-phase.sh "$RUN_DIR" create
# Apply and wait for the test resources.

cluster/test/nofork/mark-phase.sh "$RUN_DIR" observe
# Leave resources unchanged for at least one poll interval.

cluster/test/nofork/mark-phase.sh "$RUN_DIR" update
# Apply an in-place update and wait for Ready/Synced.

cluster/test/nofork/mark-phase.sh "$RUN_DIR" provider-restart
# Restart the service provider and verify recovery.

cluster/test/nofork/mark-phase.sh "$RUN_DIR" delete
# Delete resources and verify both Kubernetes and OCI cleanup.

kill -INT "$COLLECTOR_PID"
wait "$COLLECTOR_PID"
```

Set `--duration` for a bounded unattended run. `--pod-pattern` can narrow
collection to specific family or service provider pod names.

## Built-in provider metrics

The collector samples Crossplane lifecycle, state, Upjet, controller-runtime,
Go process, and no-fork provider-meta cache metrics. Cache metrics include:

- hits and misses, with absent/configuration-changed miss reasons;
- callers waiting on a shared initialization and their outcomes;
- initialization counts and durations by success/error result;
- current entries, configured capacity, in-flight initializations and LRU
  evictions.

No ProviderConfig UID, name, configuration hash, or credential is used as a
metric label.

## Prometheus

`podmonitor.yaml` is optional. Apply it only when the Prometheus Operator CRDs
are installed and configure the Prometheus selector to discover the PodMonitor.
The shell collector works without Prometheus and stores filtered text-format
samples in the run directory.

## Sharing results

The collector intentionally avoids dumping ProviderConfig and Secret objects.
Provider logs, Kubernetes events, context names, node names, and OCI resource
identifiers may still be environment-sensitive. Review and sanitize the run
directory before attaching it to a public pull request.
