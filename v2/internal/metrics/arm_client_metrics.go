/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type HTTPMethod string

const (
	HttpPut    HTTPMethod = http.MethodPut
	HttpDelete HTTPMethod = http.MethodDelete
	HttpGet    HTTPMethod = http.MethodGet
)

type ARMClientMetrics struct {
	azureRequestsTotal       *prometheus.CounterVec
	azureFailedRequestsTotal *prometheus.CounterVec
	azureRequestsTime        *prometheus.HistogramVec
}

var _ Metrics = &ARMClientMetrics{}

func NewARMClientMetrics() ARMClientMetrics {

	azureRequestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "azure_requests_total",
		Help: "Total number of requests to azure",
	}, []string{"resource", "requestType", "responseCode"})

	azureFailedRequestsTotal := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "azure_failed_requests_total",
		Help: "Total number of failed requests to azure",
	}, []string{"resource", "requestType"})

	azureRequestsTime := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "azure_requests_time_seconds",
		Help: "Length of time per ARM request",
	}, []string{"resource", "requestType"})

	return ARMClientMetrics{
		azureRequestsTotal:       azureRequestsTotal,
		azureFailedRequestsTotal: azureFailedRequestsTotal,
		azureRequestsTime:        azureRequestsTime,
	}
}

// RegisterMetrics registers the collectors with prometheus server.
func (a ARMClientMetrics) RegisterMetrics() {
	metrics.Registry.MustRegister(a.azureRequestsTime, a.azureRequestsTotal, a.azureFailedRequestsTotal)
}

// RecordAzureRequestsTotal records the total(failed + successful) number requests to ARM by increasing the counter.
func (a ARMClientMetrics) RecordAzureRequestsTotal(resourceName string, statusCode int, method HTTPMethod) {
	a.azureRequestsTotal.WithLabelValues(resourceName, string(method), strconv.Itoa(statusCode)).Inc()
}

// RecordAzureFailedRequestsTotal records the number of failed requests to ARM.
func (a ARMClientMetrics) RecordAzureFailedRequestsTotal(resourceName string, method HTTPMethod) {
	a.azureFailedRequestsTotal.WithLabelValues(resourceName, string(method)).Inc()
}

// RecordAzureRequestsTime records the round-trip time taken by the request to ARM.
func (a ARMClientMetrics) RecordAzureRequestsTime(resourceName string, requestTime time.Duration, method HTTPMethod) {
	a.azureRequestsTime.WithLabelValues(resourceName, string(method)).Observe(requestTime.Seconds())
}
