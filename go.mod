module github.com/coreswitch/zebra

go 1.12

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

replace go.etcd.io/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.5

require (
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/coreswitch/cmd v0.0.0-20171009065211-39afb7adac20
	github.com/coreswitch/component v1.0.0
	github.com/coreswitch/dependency v1.0.0 // indirect
	github.com/coreswitch/log v0.0.0-20180520054427-319a7dcf0937
	github.com/coreswitch/netutil v0.0.0-20180206171925-505fe400e6fb
	github.com/coreswitch/openconfigd v0.8.1
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/howeyc/fsnotify v0.9.0
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/mitchellh/mapstructure v1.3.3
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/prometheus/client_golang v1.8.0 // indirect
	github.com/sirupsen/logrus v1.7.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/vishvananda/netlink v1.1.0
	github.com/vishvananda/netns v0.0.0-20200728191858-db3c7e526aae
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
	golang.org/x/sys v0.0.0-20201101102859-da207088b7d1
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e // indirect
	google.golang.org/grpc v1.33.1
	google.golang.org/grpc/examples v0.0.0-20201030225255-4e179b8d3ec4 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)
