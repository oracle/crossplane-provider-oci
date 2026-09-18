#!/usr/bin/env bash

set -euo pipefail

if (($# != 2)); then
  echo "usage: mark-phase.sh RUN_DIR PHASE" >&2
  exit 2
fi

run_dir=$1
phase=$(printf '%s' "$2" | tr -cs '[:alnum:]._-' '-')
if [[ ! -f $run_dir/current-phase ]]; then
  echo "collector run directory not found: $run_dir" >&2
  exit 1
fi

printf '%s\n' "$phase" >"$run_dir/current-phase"
printf '%s\t%s\n' "$(date -u +%Y-%m-%dT%H:%M:%SZ)" "$phase" >>"$run_dir/phases.tsv"
echo "phase: $phase"
