package victoria

import (
	"context"
	"fmt"
	"strings"
	"time"

	victoriametrics "github.com/VictoriaMetrics/metrics"
	"github.com/dottedmag/go-http-metrics/metrics"
)

type Config struct {
	set *victoriametrics.Set
}

type recorder struct {
	set *victoriametrics.Set
}

func NewRecorder(cfg Config) metrics.Recorder {
	return &recorder{set: cfg.set}
}

func name(s string, kv ...string) string {
	var out []string
	for i := 0; i < len(kv); i += 2 {
		out = append(out, fmt.Sprintf("%s=%q", kv[i], kv[i+1]))
	}
	return s + "{" + strings.Join(out, ",") + "}"
}

func (r recorder) ObserveHTTPRequestDuration(_ context.Context, p metrics.HTTPReqProperties, duration time.Duration) {
	k := fmt.Sprintf("http_request_duration_seconds{code=%q,handler=%q,method=%q,service=%q}",
		p.Code, p.ID, p.Method, p.Service)

	h := r.set.GetOrCreateHistogram(k)
	h.Update(duration.Seconds())
}

func (r recorder) ObserveHTTPResponseSize(_ context.Context, p metrics.HTTPReqProperties, sizeBytes int64) {
	k := fmt.Sprintf("http_response_size_bytes{code=%q,handler=%q,method=%q,service=%q}",
		p.Code, p.ID, p.Method, p.Service)
	h := r.set.GetOrCreateHistogram(k)
	h.Update(float64(sizeBytes))
}

func (r recorder) AddInflightRequests(_ context.Context, p metrics.HTTPProperties, quantity int) {
	k := fmt.Sprintf("http_requests_inflight{handler=%q,service=%q}", p.ID, p.Service)
	g := r.set.GetOrCreateGauge(k, nil)
	g.Add(float64(quantity))
}
