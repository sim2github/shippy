package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	u "github.com/EwanValentine/shippy/srv/user/proto/user"
	user "github.com/EwanValentine/shippy/srv/user/proto/user"
	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
)

func runClient(s micro.Service) {
	var (
		uc    = u.NewUserService("go.micro.srv.user", s.Client())
		user1 = &u.User{
			Name:     "Ewan Valentine",
			Email:    fmt.Sprintf("ewan.valentine%d@gmail.com", time.Now().UnixNano()),
			Password: "test123",
			Company:  "BBC",
		}
	)

	// Create user
	rsp, err := uc.Create(context.Background(), user1)
	if err != nil {
		log.Fatalf("Could not create user: %v", err)
	}

	log.Printf("Created: %s", rsp.User.Id)

	// Print all users
	rsp, err = uc.GetAll(context.Background(), &u.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range rsp.Users {
		log.Printf("%+v\n", v)
	}

	// Auth with previous created
	auth, err := uc.Auth(context.TODO(), user1)
	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", user1.Email, err)
	}

	log.Printf("Your access token is: %s \n", auth.Token)
}

func main() {

	// Creates a database connection and handles
	// closing it again before exit.
	db, err := CreateConnection()
	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// Automatically migrates the user struct
	// into database columns/types etc. This will
	// check for changes and migrate them each time
	// this service is restarted.
	db.AutoMigrate(&user.User{})

	repo := &UserRepository{db}

	tokenService := &TokenService{repo}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),

		// Setup some flags. Specify --run_client to run the client
		// Add runtime flags
		// We could do this below too
		micro.Flags(
			&cli.BoolFlag{
				Name:  "run_client",
				Usage: "Launch the client",
			},
		),
	)

	// Init will parse the command line flags.
	srv.Init(
		// Add runtime action
		// We could actually do this above
		micro.Action(func(c *cli.Context) error {
			if c.Bool("run_client") {
				runClient(srv)
				os.Exit(0)
			}
			return nil
		}),
	)

	// Register handler
	user.RegisterUserServiceHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
