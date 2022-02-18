package main

import (
	"context"
	"fmt"
	"log"

	"example.com/grpc-sql/company/companypb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm Client")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("Could not connect : %v", err)
	}

	defer cc.Close()

	c := companypb.NewCompanyServiceClient(cc)
	fmt.Println("Client created : ", c)

	//calling get function
	getCompany(c)

	//calling post function
	addNewCompany(c)

	//calling update function
	UpdateCompany(c)

	//calling delete function
	deleteCompany(c)
}

//This function is to get the company of specific id
func getCompany(c companypb.CompanyServiceClient) {
	fmt.Println("Starting to do getById function ...")
	req := &companypb.GetRequest{
		Id: 2,
	}
	res, err := c.Get(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response from the server : ", res)
}

//This function is to add new company to the database
func addNewCompany(c companypb.CompanyServiceClient) {
	fmt.Println("Starting to do addNewCompany function  ...")

	req := &companypb.PostRequest{
		Name:   "Amazon",
		Person: "Jeff",
	}

	res, err := c.Post(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New Added row's ID : ", res.Id)
}

//This function is update the row in a table
func UpdateCompany(c companypb.CompanyServiceClient) {
	fmt.Println("Starting to do updateRow function...")

	req := &companypb.UpdateRequest{
		Company: &companypb.Company{
			Id:     2,
			Name:   "Tesla",
			Person: "Elon",
		},
	}

	res, err := c.Update(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated row:")
	fmt.Println("From : ", res)
	fmt.Println("TO : ", req)
}

//This function is to delete a row form table by its ID
func deleteCompany(c companypb.CompanyServiceClient) {
	fmt.Println("starting to do deleteCompany function ...")

	req := &companypb.DeleteRequest{
		Id: 1,
	}
	deleted_row, err := c.Delete(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Company with name: %v \nwith Person: %v is removed \n", deleted_row.Company.GetName(), deleted_row.Company.GetPerson())
}
