package metrics

import (
	"github.com/MagalixCorp/magalix-agent/scanner"
	"time"
)

// Deprecated: MetricsSource interface is deprecated and will be removed
// in future releases. Consider using Source interface instead.
// MetricsSource interface for metrics source
type MetricsSource interface {
	GetMetrics(scanner *scanner.Scanner) ([]*Metrics, map[string]interface{}, error)
}

// Source interface is to be implemented by metrics sources
type Source interface {
	GetMetrics(time time.Time) (map[string]*MetricFamily, error)
}
