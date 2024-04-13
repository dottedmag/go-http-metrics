package victoria

import (
	"bytes"
	"context"
	"net/http"
	"testing"
	"time"

	libvictoria "github.com/VictoriaMetrics/metrics"
	"github.com/dottedmag/go-http-metrics/metrics"
	"github.com/stretchr/testify/assert"
)

func TestVictoriaRecorder(t *testing.T) {
	tests := []struct {
		name          string
		recordMetrics func(r metrics.Recorder)
		expMetrics    []string
	}{
		{
			name: "Default configuration should measure with the default metric style.",
			recordMetrics: func(r metrics.Recorder) {
				r.ObserveHTTPRequestDuration(context.TODO(), metrics.HTTPReqProperties{Service: "svc1", ID: "test1", Method: http.MethodGet, Code: "200"}, 5*time.Second)
				r.ObserveHTTPRequestDuration(context.TODO(), metrics.HTTPReqProperties{Service: "svc1", ID: "test1", Method: http.MethodGet, Code: "200"}, 175*time.Millisecond)
				r.ObserveHTTPRequestDuration(context.TODO(), metrics.HTTPReqProperties{Service: "svc1", ID: "test2", Method: http.MethodGet, Code: "201"}, 50*time.Millisecond)
				r.ObserveHTTPRequestDuration(context.TODO(), metrics.HTTPReqProperties{Service: "svc2", ID: "test3", Method: http.MethodPost, Code: "500"}, 700*time.Millisecond)
				r.ObserveHTTPResponseSize(context.TODO(), metrics.HTTPReqProperties{Service: "svc1", ID: "test4", Method: http.MethodPost, Code: "500"}, 529930)
				r.ObserveHTTPResponseSize(context.TODO(), metrics.HTTPReqProperties{Service: "svc1", ID: "test4", Method: http.MethodPost, Code: "500"}, 231)
				r.ObserveHTTPResponseSize(context.TODO(), metrics.HTTPReqProperties{Service: "svc2", ID: "test4", Method: http.MethodPatch, Code: "429"}, 99999999)
				r.AddInflightRequests(context.TODO(), metrics.HTTPProperties{Service: "svc1", ID: "test1"}, 5)
				r.AddInflightRequests(context.TODO(), metrics.HTTPProperties{Service: "svc1", ID: "test1"}, -3)
				r.AddInflightRequests(context.TODO(), metrics.HTTPProperties{Service: "svc2", ID: "test2"}, 9)
			},
			expMetrics: []string{
				`http_request_duration_seconds_bucket{code="200",handler="test1",method="GET",service="svc1",vmrange=`,
				`http_request_duration_seconds_count{code="200",handler="test1",method="GET",service="svc1"} 2`,

				`http_request_duration_seconds_bucket{code="201",handler="test2",method="GET",service="svc1",vmrange=`,
				`http_request_duration_seconds_count{code="201",handler="test2",method="GET",service="svc1"} 1`,

				`http_request_duration_seconds_bucket{code="500",handler="test3",method="POST",service="svc2",vmrange=`,
				`http_request_duration_seconds_count{code="500",handler="test3",method="POST",service="svc2"} 1`,

				`http_response_size_bytes_bucket{code="429",handler="test4",method="PATCH",service="svc2",vmrange=`,
				`http_response_size_bytes_count{code="429",handler="test4",method="PATCH",service="svc2"} 1`,

				`http_response_size_bytes_bucket{code="500",handler="test4",method="POST",service="svc1",vmrange=`,
				`http_response_size_bytes_count{code="500",handler="test4",method="POST",service="svc1"} 2`,

				`http_requests_inflight{handler="test1",service="svc1"} 2`,

				`http_requests_inflight{handler="test2",service="svc2"} 9`,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			set := libvictoria.NewSet()
			mrecorder := NewRecorder(Config{set: set})

			test.recordMetrics(mrecorder)

			buf := bytes.NewBuffer(nil)
			set.WritePrometheus(buf)

			out := buf.String()

			// Check all metrics are present.
			for _, expMetric := range test.expMetrics {
				assert.Contains(out, expMetric, "metric not present on the result")
			}
		})
	}
}
