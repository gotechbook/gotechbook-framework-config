package config

import "time"

type Conn struct {
	GoTechBookFrameworkConnRateLimitingLimit        int           `json:"go-tech-book-framework-conn-rateLimiting-limit" yaml:"go_tech_book_framework_conn_rate_limiting_limit"`
	GoTechBookFrameworkConnRateLimitingInterval     time.Duration `json:"go-tech-book-framework-conn-rateLimiting-interval" yaml:"go_tech_book_framework_conn_rate_limiting_interval"`
	GoTechBookFrameworkConnRateLimitingForceDisable bool          `json:"go-tech-book-framework-conn-rateLimiting-forceDisable" yaml:"go_tech_book_framework_conn_rate_limiting_force_disable"`
}

func DefaultConn() *Conn {
	return &Conn{
		GoTechBookFrameworkConnRateLimitingLimit:        20,
		GoTechBookFrameworkConnRateLimitingInterval:     time.Duration(time.Second),
		GoTechBookFrameworkConnRateLimitingForceDisable: false,
	}
}
