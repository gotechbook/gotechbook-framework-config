package config

type Worker struct {
	GoTechBookFrameworkWorkerConcurrency      int    `json:"go-tech-book-framework-worker-concurrency" yaml:"go_tech_book_framework_worker_concurrency"`
	GoTechBookFrameworkWorkerRedisPool        string `json:"go-tech-book-framework-worker-redis-pool" yaml:"go_tech_book_framework_worker_redis_pool"`
	GoTechBookFrameworkWorkerRedisUrl         string `json:"go-tech-book-framework-worker-redis-url" yaml:"go_tech_book_framework_worker_redis_url"`
	GoTechBookFrameworkWorkerRedisPassword    string `json:"go-tech-book-framework-worker-redis-password" yaml:"go_tech_book_framework_worker_redis_password"`
	GoTechBookFrameworkWorkerNamespace        string `json:"go-tech-book-framework-worker-namespace" yaml:"go_tech_book_framework_worker_namespace"`
	GoTechBookFrameworkWorkerRetryEnabled     bool   `json:"go-tech-book-framework-worker-retry-enabled" yaml:"go_tech_book_framework_worker_retry_enabled"`
	GoTechBookFrameworkWorkerRetryExponential int    `json:"go-tech-book-framework-worker-retry-exponential" yaml:"go_tech_book_framework_worker_retry_exponential"`
	GoTechBookFrameworkWorkerRetryMax         int    `json:"go-tech-book-framework-worker-retry-max" yaml:"go_tech_book_framework_worker_retry_max"`
	GoTechBookFrameworkWorkerRetryMaxDelay    int    `json:"go-tech-book-framework-worker-retry-maxDelay" yaml:"go_tech_book_framework_worker_retry_max_delay"`
	GoTechBookFrameworkWorkerRetryMaxRandom   int    `json:"go-tech-book-framework-worker-retry-maxRandom" yaml:"go_tech_book_framework_worker_retry_max_random"`
	GoTechBookFrameworkWorkerRetryMinDelay    int    `json:"go-tech-book-framework-worker-retry-minDelay" yaml:"go_tech_book_framework_worker_retry_min_delay"`
}

func DefaultWorker() *Worker {
	return &Worker{
		GoTechBookFrameworkWorkerConcurrency:      1,
		GoTechBookFrameworkWorkerRedisPool:        "10",
		GoTechBookFrameworkWorkerRedisUrl:         "localhost:6379",
		GoTechBookFrameworkWorkerRetryEnabled:     true,
		GoTechBookFrameworkWorkerRetryExponential: 5,
		GoTechBookFrameworkWorkerRetryMax:         2,
		GoTechBookFrameworkWorkerRetryMaxDelay:    10,
		GoTechBookFrameworkWorkerRetryMaxRandom:   0,
		GoTechBookFrameworkWorkerRetryMinDelay:    10,
		GoTechBookFrameworkWorkerNamespace:        PREFIX,
	}
}
