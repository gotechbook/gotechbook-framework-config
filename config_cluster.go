package config

import "time"

type Cluster struct {
	GoTechBookFrameworkClusterInfoRegion                    string        `json:"go-tech-book-framework-cluster-info-region" yaml:"go_tech_book_framework_cluster_info_region"`
	GoTechBookFrameworkClusterRpcClientGrpcDialTimeout      time.Duration `json:"go-tech-book-framework-cluster-rpc-client-grpc-dial-timeout" yaml:"go_tech_book_framework_cluster_rpc_client_grpc_dial_timeout"`
	GoTechBookFrameworkClusterRpcClientGrpcRequestTimeout   time.Duration `json:"go-tech-book-framework-cluster-rpc-client-grpc-request-timeout" yaml:"go_tech_book_framework_cluster_rpc_client_grpc_request_timeout"`
	GoTechBookFrameworkClusterRpcClientGrpcLazyConnection   bool          `json:"go-tech-book-framework-cluster-rpc-client-grpc-lazy-connection" yaml:"go_tech_book_framework_cluster_rpc_client_grpc_lazy_connection"`
	GoTechBookFrameworkClusterRpcServerGrpcPort             int           `json:"go-tech-book-framework-cluster-rpc-server-grpc-port" yaml:"go_tech_book_framework_cluster_rpc_server_grpc_port"`
	GoTechBookFrameworkClusterSdEtcdUser                    string        `json:"go-tech-book-framework-cluster-sd-etcd-user" yaml:"go_tech_book_framework_cluster_sd_etcd_user"`
	GoTechBookFrameworkClusterSdEtcdPass                    string        `json:"go-tech-book-framework-cluster-sd-etcd-pass" yaml:"go_tech_book_framework_cluster_sd_etcd_pass"`
	GoTechBookFrameworkClusterSdEtcdDialTimeout             time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-dial-timeout" yaml:"go_tech_book_framework_cluster_sd_etcd_dial_timeout"`
	GoTechBookFrameworkClusterSdEtcdEndpoints               []string      `json:"go-tech-book-framework-cluster-sd-etcd-endpoints" yaml:"go_tech_book_framework_cluster_sd_etcd_endpoints"`
	GoTechBookFrameworkClusterSdEtcdPrefix                  string        `json:"go-tech-book-framework-cluster-sd-etcd-prefix" yaml:"go_tech_book_framework_cluster_sd_etcd_prefix"`
	GoTechBookFrameworkClusterSdEtcdGrantLeaseMaxRetries    int           `json:"go-tech-book-framework-cluster-sd-etcd-grant-lease-max-retries" yaml:"go_tech_book_framework_cluster_sd_etcd_grant_lease_max_retries"`
	GoTechBookFrameworkClusterSdEtcdGrantLeaseRetryInterval time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-grant-lease-retry-interval" yaml:"go_tech_book_framework_cluster_sd_etcd_grant_lease_retry_interval"`
	GoTechBookFrameworkClusterSdEtcdGrantLeaseTimeout       time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-grant-lease-timeout" yaml:"go_tech_book_framework_cluster_sd_etcd_grant_lease_timeout"`
	GoTechBookFrameworkClusterSdEtcdHeartbeatLog            bool          `json:"go-tech-book-framework-cluster-sd-etcd-heartbeat-log" yaml:"go_tech_book_framework_cluster_sd_etcd_heartbeat_log"`
	GoTechBookFrameworkClusterSdEtcdHeartbeatTtl            time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-heartbeat-ttl" yaml:"go_tech_book_framework_cluster_sd_etcd_heartbeat_ttl"`
	GoTechBookFrameworkClusterSdEtcdRevokeTimeout           time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-revoke-timeout" yaml:"go_tech_book_framework_cluster_sd_etcd_revoke_timeout"`
	GoTechBookFrameworkClusterSdEtcdSyncServersInterval     time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-sync-servers-interval" yaml:"go_tech_book_framework_cluster_sd_etcd_sync_servers_interval"`
	GoTechBookFrameworkClusterSdEtcdSyncServersParallelism  int           `json:"go-tech-book-framework-cluster-sd-etcd-sync-servers-parallelism" yaml:"go_tech_book_framework_cluster_sd_etcd_sync_servers_parallelism"`
	GoTechBookFrameworkClusterSdEtcdShutdownDelay           time.Duration `json:"go-tech-book-framework-cluster-sd-etcd-shutdown-delay" yaml:"go_tech_book_framework_cluster_sd_etcd_shutdown_delay"`
	GoTechBookFrameworkClusterSdEtcdServerTypeBlacklist     []string      `json:"go-tech-book-framework-cluster-sd-etcd-server-type-blacklist" yaml:"go_tech_book_framework_cluster_sd_etcd_server_type_blacklist"`
}

func DefaultCluster() *Cluster {
	return &Cluster{
		GoTechBookFrameworkClusterSdEtcdUser:                    "",
		GoTechBookFrameworkClusterSdEtcdPass:                    "",
		GoTechBookFrameworkClusterInfoRegion:                    "",
		GoTechBookFrameworkClusterRpcClientGrpcDialTimeout:      time.Duration(5 * time.Second),
		GoTechBookFrameworkClusterRpcClientGrpcRequestTimeout:   time.Duration(5 * time.Second),
		GoTechBookFrameworkClusterRpcClientGrpcLazyConnection:   false,
		GoTechBookFrameworkClusterRpcServerGrpcPort:             1234,
		GoTechBookFrameworkClusterSdEtcdDialTimeout:             time.Duration(5 * time.Second),
		GoTechBookFrameworkClusterSdEtcdEndpoints:               []string{"localhost:2379"},
		GoTechBookFrameworkClusterSdEtcdPrefix:                  PREFIX,
		GoTechBookFrameworkClusterSdEtcdGrantLeaseMaxRetries:    15,
		GoTechBookFrameworkClusterSdEtcdGrantLeaseRetryInterval: time.Duration(5 * time.Second),
		GoTechBookFrameworkClusterSdEtcdGrantLeaseTimeout:       time.Duration(60 * time.Second),
		GoTechBookFrameworkClusterSdEtcdHeartbeatLog:            false,
		GoTechBookFrameworkClusterSdEtcdHeartbeatTtl:            time.Duration(60 * time.Second),
		GoTechBookFrameworkClusterSdEtcdRevokeTimeout:           time.Duration(5 * time.Second),
		GoTechBookFrameworkClusterSdEtcdSyncServersInterval:     time.Duration(120 * time.Second),
		GoTechBookFrameworkClusterSdEtcdSyncServersParallelism:  10,
		GoTechBookFrameworkClusterSdEtcdShutdownDelay:           time.Duration(300 * time.Millisecond),
		GoTechBookFrameworkClusterSdEtcdServerTypeBlacklist:     nil,
	}
}
