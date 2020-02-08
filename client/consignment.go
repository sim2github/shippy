package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	c "github.com/EwanValentine/shippy/srv/consignment/proto/consignment"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
)

var (
	endpoint = flag.String("endpoint", "localhost:50501", "go.micro.srv.consignment address")
	file     = flag.String("file", "consignment.json", "consignment file")
	token    = flag.String("token", "", "access token")
)

func parseFile(file string) (*c.Consignment, error) {
	var consignment *c.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	flag.Parse()

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": *token,
	})

	consignment, err := parseFile(*file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Contact the server and print out its response.
	cc := c.NewShippingService("go.micro.srv.consignment", client.DefaultClient)
	rsp, err := cc.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %t", rsp.Created)
	fmt.Printf("rs = %+v\n", rsp)

	// rsp, err = client.GetConsignments(ctx, &c.GetRequest{})
	// if err != nil {
	// 	log.Fatalf("Could not list consignments: %v", err)
	// }
	// for _, v := range rsp.Consignments {
	// 	log.Println(v)
	// }
}
