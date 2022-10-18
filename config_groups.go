package config

import "time"

type Groups struct {
	GoTechBookFrameworkGroupsEtcdDialTimeout        time.Duration `json:"go-tech-book-framework-groups-etcd-dialTimeout" yaml:"go_tech_book_framework_groups_etcd_dial_timeout"`
	GoTechBookFrameworkGroupsEtcdEndpoints          []string      `json:"go-tech-book-framework-groups-etcd-endpoints" yaml:"go_tech_book_framework_groups_etcd_endpoints"`
	GoTechBookFrameworkGroupsEtcdPrefix             string        `json:"go-tech-book-framework-groups-etcd-prefix" yaml:"go_tech_book_framework_groups_etcd_prefix"`
	GoTechBookFrameworkGroupsEtcdTransactionTimeout time.Duration `json:"go-tech-book-framework-groups-etcd-transactionTimeout" yaml:"go_tech_book_framework_groups_etcd_transaction_timeout"`
	GoTechBookFrameworkGroupsMemoryTickDuration     time.Duration `json:"go-tech-book-framework-groups-memory-tickDuration" yaml:"go_tech_book_framework_groups_memory_tick_duration"`
}

func DefaultGroups() *Groups {
	return &Groups{
		GoTechBookFrameworkGroupsEtcdDialTimeout:        time.Duration(5 * time.Second),
		GoTechBookFrameworkGroupsEtcdEndpoints:          []string{"localhost:2379"},
		GoTechBookFrameworkGroupsEtcdPrefix:             PREFIX,
		GoTechBookFrameworkGroupsEtcdTransactionTimeout: time.Duration(5 * time.Second),
		GoTechBookFrameworkGroupsMemoryTickDuration:     time.Duration(30 * time.Second),
	}
}
