package config

type Metrics struct {
	GoTechBookFrameworkMetricsConstTags                map[string]string  `json:"go-tech-book-framework-metrics-constTags" yaml:"go_tech_book_framework_metrics_const_tags"`
	GoTechBookFrameworkMetricsCustom                   *CustomMetricsSpec `json:"go-tech-book-framework-metrics-custom" yaml:"go_tech_book_framework_metrics_custom"`
	GoTechBookFrameworkMetricsPrometheusAdditionalTags map[string]string  `json:"go-tech-book-framework-metrics-prometheus-additionalTags" yaml:"go_tech_book_framework_metrics_prometheus_additional_tags"`
	GoTechBookFrameworkMetricsPrometheusEnabled        bool               `json:"go-tech-book-framework-metrics-prometheus-enabled" yaml:"go_tech_book_framework_metrics_prometheus_enabled"`
	GoTechBookFrameworkMetricsPrometheusPort           int                `json:"go-tech-book-framework-metrics-prometheus-port" yaml:"go_tech_book_framework_metrics_prometheus_port"`
	GoTechBookFrameworkMetricsStatsdEnabled            bool               `json:"go-tech-book-framework-metrics-statsd-enabled" yaml:"go_tech_book_framework_metrics_statsd_enabled"`
	GoTechBookFrameworkMetricsStatsdHost               string             `json:"go-tech-book-framework-metrics-statsd-host" yaml:"go_tech_book_framework_metrics_statsd_host"`
	GoTechBookFrameworkMetricsStatsdPrefix             string             `json:"go-tech-book-framework-metrics-statsd-prefix" yaml:"go_tech_book_framework_metrics_statsd_prefix"`
	GoTechBookFrameworkMetricsStatsdRate               float64            `json:"go-tech-book-framework-metrics-statsd-rate" yaml:"go_tech_book_framework_metrics_statsd_rate"`
}
type Summary struct {
	Subsystem  string              `json:"subsystem" yaml:"subsystem"`
	Name       string              `json:"name" yaml:"name"`
	Help       string              `json:"help" yaml:"help"`
	Objectives map[float64]float64 `json:"objectives" yaml:"objectives"`
	Labels     []string            `json:"labels" yaml:"labels"`
}

type Histogram struct {
	Subsystem string    `json:"subsystem" yaml:"subsystem"`
	Name      string    `json:"name" yaml:"name"`
	Help      string    `json:"help" yaml:"help"`
	Buckets   []float64 `json:"buckets" yaml:"buckets"`
	Labels    []string  `json:"labels" yaml:"labels"`
}

type Gauge struct {
	Subsystem string   `json:"subsystem" yaml:"subsystem"`
	Name      string   `json:"name" yaml:"name"`
	Help      string   `json:"help" yaml:"help"`
	Labels    []string `json:"labels" yaml:"labels"`
}

type Counter struct {
	Subsystem string   `json:"subsystem" yaml:"subsystem"`
	Name      string   `json:"name" yaml:"name"`
	Help      string   `json:"help" yaml:"help"`
	Labels    []string `json:"labels" yaml:"labels"`
}

type CustomMetricsSpec struct {
	Summaries  []*Summary   `json:"summaries" yaml:"summaries"`
	Histograms []*Histogram `json:"histograms" yaml:"histograms"`
	Gauges     []*Gauge     `json:"gauges" yaml:"gauges"`
	Counters   []*Counter   `json:"counters" yaml:"counters"`
}

func DefaultMetrics() *Metrics {
	return &Metrics{
		GoTechBookFrameworkMetricsConstTags:                map[string]string{},
		GoTechBookFrameworkMetricsCustom:                   &CustomMetricsSpec{Summaries: []*Summary{}, Gauges: []*Gauge{}, Counters: []*Counter{}},
		GoTechBookFrameworkMetricsPrometheusAdditionalTags: map[string]string{},
		GoTechBookFrameworkMetricsPrometheusEnabled:        true,
		GoTechBookFrameworkMetricsPrometheusPort:           9090,
		GoTechBookFrameworkMetricsStatsdEnabled:            true,
		GoTechBookFrameworkMetricsStatsdHost:               "localhost:9125",
		GoTechBookFrameworkMetricsStatsdPrefix:             PREFIX,
		GoTechBookFrameworkMetricsStatsdRate:               1,
	}
}
