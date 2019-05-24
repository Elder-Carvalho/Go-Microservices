package main

import (
	"context"
	"encoding/json"
	pb "github.com/Elder-Carvalho/shippy-service-consignment/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
	"io/ioutil"
	"log"
	"os"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func main() {
	service := micro.NewService(micro.Name("shipping.service.client"))
	service.Init()

	client := pb.NewShippingService("shipping.service", service.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalln("Failed not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	all, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Failed to get consignments: %v", err)
	}

	for _, consignment := range all.Consignments {
		log.Println(consignment)
	}
}
