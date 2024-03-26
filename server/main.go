package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Dertyxx/grpc-test-1/proto" // Replace with your package name
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedPersonServiceServer
	people map[string]*pb.Person
}

func (s *server) CreatePerson(ctx context.Context, req *pb.CreatePersonRequest) (*pb.CreatePersonResponse, error) {
	person := req.GetPerson()
	if person.GetId() == "" {
		return nil, status.Error(codes.InvalidArgument, "ID is required")
	}

	if _, exists := s.people[person.GetId()]; exists {
		return nil, status.Error(codes.AlreadyExists, "Person with ID already exists")
	}

	s.people[person.GetId()] = person
	return &pb.CreatePersonResponse{Person: person}, nil
}

func (s *server) ReadPerson(ctx context.Context, req *pb.ReadPersonRequest) (*pb.ReadPersonResponse, error) {
	id := req.GetId()
	person, ok := s.people[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "Person not found")
	}

	return &pb.ReadPersonResponse{Person: person}, nil
}

func (s *server) UpdatePerson(ctx context.Context, req *pb.UpdatePersonRequest) (*pb.UpdatePersonResponse, error) {
	person := req.GetPerson()
	id := person.GetId()

	_, ok := s.people[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "Person not found")
	}

	s.people[id] = person
	return &pb.UpdatePersonResponse{Person: person}, nil
}

func (s *server) GetAllPersons(ctx context.Context, req *pb.GetAllPersonsRequest) (*pb.GetAllPersonsResponse, error) {
	people := make([]*pb.Person, 0, len(s.people))
	for _, p := range s.people {
		people = append(people, p)
	}
	return &pb.GetAllPersonsResponse{Persons: people}, nil
}

func (s *server) DeletePerson(ctx context.Context, req *pb.DeletePersonRequest) (*pb.DeletePersonResponse, error) {
	id := req.GetId()
	_, ok := s.people[id]
	if !ok {
		return nil, status.Error(codes.NotFound, "Person not found")
	}

	delete(s.people, id)
	return &pb.DeletePersonResponse{Success: true}, nil
}

func main() {
	fmt.Println("gRPC server running...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPersonServiceServer(s, &server{people: make(map[string]*pb.Person)})

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
