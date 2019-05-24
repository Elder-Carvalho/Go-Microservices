package main

import (
	"context"
	"fmt"
	pb "github.com/Elder-Carvalho/shippy-service-consignment/consignment-service/proto/consignment"
	micro "github.com/micro/go-micro"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Mock de um repositorio
type Repository struct {
	consignments []*pb.Consignment
}

// Cria uma nova consignação
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := s.repo.Create(req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &Repository{}

	// Create a new service instance
	srv := micro.NewService(micro.Name("shipping.service")) // O nome do serviço deve ser igual ao nome do pacote de protobuf

	// Inicia o servidor
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{repo})

	// Roda o servidor
	if err := srv.Run(); err != nil {
		fmt.Println("failed to serve: %v", err)
	}
}
