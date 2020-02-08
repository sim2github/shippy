module github.com/EwanValentine/shippy/srv/vessel

// TODO temporary fix package issues
replace (
	// TODO: Prevent usage github.com/EwanValentine/shippy v0.0.0-00010101000000-000000000000
	github.com/EwanValentine/shippy => ../../../shippy
	github.com/EwanValentine/shippy/srv/user/proto/user => ../../../shippy/srv/user/proto/user
	github.com/EwanValentine/shippy/srv/vessel/proto/vessel => ../../../shippy/srv/vessel/proto/vessel
	// TODO: Fix https://github.com/etcd-io/etcd/issues/11563
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

go 1.13

require (
	github.com/EwanValentine/shippy/srv/vessel/proto/vessel v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.3.2
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-micro/v2 v2.0.0
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)
