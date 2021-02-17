module github.com/sthaha/go-restful-example

go 1.15

replace (
	go.etcd.io/etcd/api/v3 => go.etcd.io/etcd/api/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/client/v3 => go.etcd.io/etcd/client/v3 v3.0.0-20201103155942-6e800b9b0161
	go.etcd.io/etcd/pkg/v3 => go.etcd.io/etcd/pkg/v3 v3.0.0-20201103155942-6e800b9b0161
)

require (
	github.com/emicklei/go-restful v2.15.0+incompatible
	go.etcd.io/etcd/client/v3 v3.0.0-20210210213918-d33a1c91f0cf
)
