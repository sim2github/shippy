module github.com/EwanValentine/shippy/srv/user

go 1.13

require (
	github.com/EwanValentine/shippy v0.0.0-00010101000000-000000000000
	github.com/EwanValentine/shippy/srv/user/proto/user v0.0.0-00010101000000-000000000000
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/cli/v2 v2.1.1
	github.com/micro/go-micro/v2 v2.0.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20200128174031-69ecbb4d6d5d
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
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
