module github.com/EwanValentine/shippy/srv/consignment

go 1.13

require (
	github.com/EwanValentine/shippy v0.0.0-20181206231135-cf1b235f5c09 // indirect
	github.com/EwanValentine/shippy/srv/user/proto/user v0.0.0-00010101000000-000000000000
	github.com/EwanValentine/shippy/srv/vessel/proto/vessel v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.3
	github.com/micro/go-micro/v2 v2.0.0
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22
)

// TODO temporary fix package issues
replace (
	// TODO: Prevent usage github.com/EwanValentine/shippy v0.0.0-00010101000000-000000000000
	github.com/EwanValentine/shippy => ../../../shippy
	github.com/EwanValentine/shippy/srv/user/proto/user => ../../../shippy/srv/user/proto/user
	github.com/EwanValentine/shippy/srv/vessel/proto/vessel => ../../../shippy/srv/vessel/proto/vessel
	// TODO: Fix https://github.com/etcd-io/etcd/issues/11563
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
