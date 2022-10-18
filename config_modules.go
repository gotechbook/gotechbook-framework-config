package config

import "time"

type Modules struct {
	GoTechBookFrameworkModulesBindingStorageEtcdDialTimeout time.Duration `json:"go-tech-book-framework-modules-binding-storage-etcd-dial-timeout" yaml:"go_tech_book_framework_modules_binding_storage_etcd_dial_timeout"`
	GoTechBookFrameworkModulesBindingStorageEtcdEndpoints   []string      `json:"go-tech-book-framework-modules-binding-storage-etcd-endpoints" yaml:"go_tech_book_framework_modules_binding_storage_etcd_endpoints"`
	GoTechBookFrameworkModulesBindingStorageEtcdLeaseTtl    time.Duration `json:"go-tech-book-framework-modules-binding-storage-etcd-lease-ttl" yaml:"go_tech_book_framework_modules_binding_storage_etcd_lease_ttl"`
	GoTechBookFrameworkModulesBindingStorageEtcdPrefix      string        `json:"go-tech-book-framework-modules-binding-storage-etcd-prefix" yaml:"go_tech_book_framework_modules_binding_storage_etcd_prefix"`
}

func DefaultModules() *Modules {
	return &Modules{
		GoTechBookFrameworkModulesBindingStorageEtcdDialTimeout: time.Duration(5 * time.Second),
		GoTechBookFrameworkModulesBindingStorageEtcdEndpoints:   []string{"localhost:2379"},
		GoTechBookFrameworkModulesBindingStorageEtcdLeaseTtl:    time.Duration(5 * time.Hour),
		GoTechBookFrameworkModulesBindingStorageEtcdPrefix:      PREFIX,
	}
}
