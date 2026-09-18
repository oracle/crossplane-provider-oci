/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"errors"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestProviderMetaCacheMetrics(t *testing.T) {
	registry := prometheus.NewRegistry()
	metrics := newProviderMetaCacheMetrics(registry)
	cache := newProviderMetaCacheWithMetrics(1, metrics)

	create := func(value string) func() (any, error) {
		return func() (any, error) { return value, nil }
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create("a")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", create("unused")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-a", "hash-b", create("a-updated")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-b", "hash-b", create("b")); err != nil {
		t.Fatal(err)
	}
	if _, err := cache.getOrCreate(t.Context(), "uid-c", "hash-c", func() (any, error) {
		return nil, errors.New("configure failed")
	}); err == nil {
		t.Fatal("getOrCreate() error = nil, want initialization error")
	}

	assertMetricValue(t, metrics.hits, 1)
	assertMetricValue(t, metrics.misses.WithLabelValues("absent"), 3)
	assertMetricValue(t, metrics.misses.WithLabelValues("configuration_changed"), 1)
	assertMetricValue(t, metrics.initializations.WithLabelValues("success"), 3)
	assertMetricValue(t, metrics.initializations.WithLabelValues("error"), 1)
	assertMetricValue(t, metrics.evictions, 1)
	assertMetricValue(t, metrics.entries, 1)
	assertMetricValue(t, metrics.inflight, 0)
	assertMetricValue(t, metrics.waiting, 0)
	assertMetricValue(t, metrics.capacity, 1)

	if got := testutil.CollectAndCount(metrics.initializationDuration); got != 2 {
		t.Fatalf("initialization duration metric count = %d, want 2 result series", got)
	}
}

func TestProviderMetaCacheWaitMetrics(t *testing.T) {
	registry := prometheus.NewRegistry()
	metrics := newProviderMetaCacheMetrics(registry)
	cache := newProviderMetaCacheWithMetrics(1, metrics)

	started := make(chan struct{})
	release := make(chan struct{})
	creatorDone := make(chan error, 1)
	go func() {
		_, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", func() (any, error) {
			close(started)
			<-release
			return "meta", nil
		})
		creatorDone <- err
	}()
	<-started

	waiterDone := make(chan error, 1)
	go func() {
		_, err := cache.getOrCreate(t.Context(), "uid-a", "hash-a", func() (any, error) {
			return "duplicate", nil
		})
		waiterDone <- err
	}()

	deadline := time.Now().Add(time.Second)
	for testutil.ToFloat64(metrics.waiting) != 1 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	assertMetricValue(t, metrics.waiting, 1)
	assertMetricValue(t, metrics.inflight, 1)

	close(release)
	if err := <-creatorDone; err != nil {
		t.Fatalf("creator error: %v", err)
	}
	if err := <-waiterDone; err != nil {
		t.Fatalf("waiter error: %v", err)
	}

	assertMetricValue(t, metrics.waits.WithLabelValues("completed"), 1)
	assertMetricValue(t, metrics.waiting, 0)
	assertMetricValue(t, metrics.inflight, 0)
}

func assertMetricValue(t *testing.T, collector prometheus.Collector, want float64) {
	t.Helper()
	if got := testutil.ToFloat64(collector); got != want {
		t.Fatalf("metric value = %v, want %v", got, want)
	}
}
