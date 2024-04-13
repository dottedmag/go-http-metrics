// Code generated by mockery v1.0.0. DO NOT EDIT.

package metrics

import (
	context "context"

	metrics "github.com/dottedmag/go-http-metrics/metrics"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Recorder is an autogenerated mock type for the Recorder type
type Recorder struct {
	mock.Mock
}

// AddInflightRequests provides a mock function with given fields: ctx, props, quantity
func (_m *Recorder) AddInflightRequests(ctx context.Context, props metrics.HTTPProperties, quantity int) {
	_m.Called(ctx, props, quantity)
}

// ObserveHTTPRequestDuration provides a mock function with given fields: ctx, props, duration
func (_m *Recorder) ObserveHTTPRequestDuration(ctx context.Context, props metrics.HTTPReqProperties, duration time.Duration) {
	_m.Called(ctx, props, duration)
}

// ObserveHTTPResponseSize provides a mock function with given fields: ctx, props, sizeBytes
func (_m *Recorder) ObserveHTTPResponseSize(ctx context.Context, props metrics.HTTPReqProperties, sizeBytes int64) {
	_m.Called(ctx, props, sizeBytes)
}
