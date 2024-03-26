package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"https://github.com/Dertyxx/grpc-test-1/proto"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	c := person.NewPersonServiceClient(conn)

	// Test CreatePerson RPC
	respCreate, err := c.CreatePerson(context.Background(), &person.CreatePersonRequest{
		Person: &person.Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		},
	})
	if err != nil {
		log.Fatalf("CreatePerson failed: %v", err)
	}
	fmt.Println("CreatePerson Response:", respCreate.GetPerson())

	// Test GetPerson RPC
	respGet, err := c.GetPerson(context.Background(), &person.GetPersonRequest{
		FirstName: "John",
	})
	if err != nil {
		log.Fatalf("GetPerson failed: %v", err)
	}
	fmt.Println("GetPerson Response:", respGet.GetPerson())

	// Test UpdatePerson RPC
	respUpdate, err := c.UpdatePerson(context.Background(), &person.UpdatePersonRequest{
		FirstName: "John",
		UpdatedPerson: &person.Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       35, // Updated age
		},
	})
	if err != nil {
		log.Fatalf("UpdatePerson failed: %v", err)
	}
	fmt.Println("UpdatePerson Response:", respUpdate.GetUpdatedPerson())

	// Test GetAllPersons RPC
	respGetAll, err := c.GetAllPersons(context.Background(), &person.GetAllPersonsRequest{})
	if err != nil {
		log.Fatalf("GetAllPersons failed: %v", err)
	}
	fmt.Println("GetAllPersons Response:", respGetAll.GetPersons())

	// Test DeletePerson RPC
	respDelete, err := c.DeletePerson(context.Background(), &person.DeletePersonRequest{
		FirstName: "John",
	})
	if err != nil {
		log.Fatalf("DeletePerson failed: %v", err)
	}
	fmt.Println("DeletePerson Response:", respDelete.GetSuccess())
}
