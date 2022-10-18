package config

import "time"

type App struct {
	GoTechBookFrameworkAppBufferAgentMessages                     int           `json:"go-tech-book-framework-app-buffer-agent-messages" yaml:"go_tech_book_framework_app_buffer_agent_messages"`
	GoTechBookFrameworkAppBufferHandlerLocalProcess               int           `json:"go-tech-book-framework-app-buffer-handler-localProcess" yaml:"go_tech_book_framework_app_buffer_handler_local_process"`
	GoTechBookFrameworkAppBufferHandlerRemoteProcess              int           `json:"go-tech-book-framework-app-buffer-handler-remoteProcess" yaml:"go_tech_book_framework_app_buffer_handler_remote_process"`
	GoTechBookFrameworkAppBufferHandlerDispatch                   int           `json:"go-tech-book-framework-app-buffer-handler-dispatch" yaml:"go_tech_book_framework_app_buffer_handler_dispatch"`
	GoTechBookFrameworkAppSessionUnique                           bool          `json:"go-tech-book-framework-app-session-unique" yaml:"go_tech_book_framework_app_session_unique"`
	GoTechBookFrameworkAppHandlerMessagesCompression              bool          `json:"go-tech-book-framework-app-handler-messages-compression" yaml:"go_tech_book_framework_app_handler_messages_compression"`
	GoTechBookFrameworkAppHeartbeatInterval                       time.Duration `json:"go-tech-book-framework-app-heartbeat-interval" yaml:"go_tech_book_framework_app_heartbeat_interval"`
	GoTechBookFrameworkAppMetricsPeriodicMetricsPeriod            time.Duration `json:"go-tech-book-framework-app-metrics-periodicMetrics-period" yaml:"go_tech_book_framework_app_metrics_periodic_metrics_period"`
	GoTechBookFrameworkAppDefaultPipelinesStructValidationEnabled bool          `json:"go-tech-book-framework-app-default-pipelines-struct-validation-enabled" yaml:"go_tech_book_framework_app_default_pipelines_struct_validation_enabled"`
}

func DefaultApp() *App {
	return &App{
		GoTechBookFrameworkAppBufferAgentMessages:                     100,
		GoTechBookFrameworkAppBufferHandlerLocalProcess:               20,
		GoTechBookFrameworkAppBufferHandlerRemoteProcess:              20,
		GoTechBookFrameworkAppBufferHandlerDispatch:                   25,
		GoTechBookFrameworkAppSessionUnique:                           true,
		GoTechBookFrameworkAppHandlerMessagesCompression:              true,
		GoTechBookFrameworkAppHeartbeatInterval:                       time.Duration(30 * time.Second),
		GoTechBookFrameworkAppMetricsPeriodicMetricsPeriod:            time.Duration(15 * time.Second),
		GoTechBookFrameworkAppDefaultPipelinesStructValidationEnabled: false,
	}
}
