package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/your_username/your_project_name/person"
)

type server struct{}

var persons []*person.Person

func (s *server) CreatePerson(ctx context.Context, req *person.CreatePersonRequest) (*person.CreatePersonResponse, error) {
	p := req.GetPerson()
	persons = append(persons, p)
	return &person.CreatePersonResponse{Person: p}, nil
}

func (s *server) GetPerson(ctx context.Context, req *person.GetPersonRequest) (*person.GetPersonResponse, error) {
	firstName := req.GetFirstName()
	for _, p := range persons {
		if p.FirstName == firstName {
			return &person.GetPersonResponse{Person: p}, nil
		}
	}
	return nil, fmt.Errorf("person with first name %s not found", firstName)
}

func (s *server) UpdatePerson(ctx context.Context, req *person.UpdatePersonRequest) (*person.UpdatePersonResponse, error) {
	firstName := req.GetFirstName()
	updatedPerson := req.GetUpdatedPerson()

	for i, p := range persons {
		if p.FirstName == firstName {
			persons[i] = updatedPerson
			return &person.UpdatePersonResponse{UpdatedPerson: updatedPerson}, nil
		}
	}
	return nil, fmt.Errorf("person with first name %s not found", firstName)
}

func (s *server) GetAllPersons(ctx context.Context, req *person.GetAllPersonsRequest) (*person.GetAllPersonsResponse, error) {
	return &person.GetAllPersonsResponse{Persons: persons}, nil
}

func (s *server) DeletePerson(ctx context.Context, req *person.DeletePersonRequest) (*person.DeletePersonResponse, error) {
	firstName := req.GetFirstName()
	for i, p := range persons {
		if p.FirstName == firstName {
			persons = append(persons[:i], persons[i+1:]...)
			return &person.DeletePersonResponse{Success: true}, nil
		}
	}
	return nil, fmt.Errorf("person with first name %s not found", firstName)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	person.RegisterPersonServiceServer(s, &server{})
	log.Println("Server started at port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
