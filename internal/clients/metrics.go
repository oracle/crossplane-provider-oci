/*
Copyright 2026 Oracle and/or its affiliates.
*/

package clients

import (
	"github.com/prometheus/client_golang/prometheus"
	crmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	metricsNamespace = "provider_oci"
	metricsSubsystem = "nofork_provider_meta_cache"
)

type providerMetaCacheMetrics struct {
	hits                   prometheus.Counter
	misses                 *prometheus.CounterVec
	waits                  *prometheus.CounterVec
	initializations        *prometheus.CounterVec
	initializationDuration *prometheus.HistogramVec
	evictions              prometheus.Counter
	entries                prometheus.Gauge
	inflight               prometheus.Gauge
	waiting                prometheus.Gauge
	capacity               prometheus.Gauge
}

func newProviderMetaCacheMetrics(reg prometheus.Registerer) *providerMetaCacheMetrics {
	m := &providerMetaCacheMetrics{
		hits: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "hits_total",
			Help:      "Number of configured in-process provider metadata cache hits.",
		}),
		misses: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "misses_total",
			Help:      "Number of configured in-process provider metadata cache misses.",
		}, []string{"reason"}),
		waits: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "waits_total",
			Help:      "Number of callers that waited for an in-flight provider initialization.",
		}, []string{"result"}),
		initializations: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "initializations_total",
			Help:      "Number of in-process provider metadata initialization attempts.",
		}, []string{"result"}),
		initializationDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "initialization_duration_seconds",
			Help:      "Time spent initializing in-process provider metadata.",
			Buckets:   prometheus.DefBuckets,
		}, []string{"result"}),
		evictions: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "evictions_total",
			Help:      "Number of least-recently-used provider metadata cache evictions.",
		}),
		entries: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "entries",
			Help:      "Current number of configured provider metadata cache entries.",
		}),
		inflight: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "inflight_initializations",
			Help:      "Current number of in-process provider metadata initializations.",
		}),
		waiting: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "waiting_callers",
			Help:      "Current number of callers waiting for an in-flight provider initialization.",
		}),
		capacity: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: metricsNamespace,
			Subsystem: metricsSubsystem,
			Name:      "capacity",
			Help:      "Configured maximum number of provider metadata cache entries.",
		}),
	}
	reg.MustRegister(
		m.hits,
		m.misses,
		m.waits,
		m.initializations,
		m.initializationDuration,
		m.evictions,
		m.entries,
		m.inflight,
		m.waiting,
		m.capacity,
	)
	return m
}

var defaultProviderMetaCacheMetrics = newProviderMetaCacheMetrics(crmetrics.Registry)
