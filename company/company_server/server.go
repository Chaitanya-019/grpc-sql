package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"database/sql"

	"github.com/go-sql-driver/mysql"

	"example.com/grpc-sql/company/companypb"
	"example.com/grpc-sql/company/database"
	"example.com/grpc-sql/company/views"
	"google.golang.org/grpc"
)

type server struct{}

var db *sql.DB

func (*server) Get(ctx context.Context, req *companypb.GetRequest) (*companypb.GetResponse, error) {
	fmt.Println("Get function is invoked")
	id := req.GetId()

	row, err := database.SelectRow(id, db)
	if err != nil || row == nil {
		return nil, err
	}

	res := &companypb.GetResponse{
		Company: row,
	}
	return res, nil

}

func (*server) Post(ctx context.Context, req *companypb.PostRequest) (*companypb.PostResponse, error) {
	fmt.Println("Post function is Invoked")
	com := views.Comp{
		Name:    req.GetName(),
		Creator: req.GetPerson(),
	}

	id, err := database.InsertRow(com, db)
	if err != nil || id == 0 {
		return nil, err
	}

	res := &companypb.PostResponse{
		Id: id,
	}
	return res, nil
}

func (*server) Update(ctx context.Context, req *companypb.UpdateRequest) (*companypb.UpdateResponse, error) {
	fmt.Println("Update function is incoked")
	update_req := views.Comp{
		ID:      req.Company.GetId(),
		Name:    req.Company.GetName(),
		Creator: req.Company.GetPerson(),
	}

	res, err := database.UpdateRow(update_req, db)
	if err != nil || res == nil {
		return nil, err
	}

	return res, nil
}

func (*server) Delete(ctx context.Context, req *companypb.DeleteRequest) (*companypb.DeleteResponse, error) {
	fmt.Println("Delete function is Invoked")
	deleted_row, err := database.DeleteRow(req, db)
	if err != nil || deleted_row == nil {
		return nil, err
	}

	return deleted_row, nil
}

func main() {
	fmt.Println("Get Ready! server is starting")
	lis, err1 := net.Listen("tcp", "0.0.0.0:50051")
	if err1 != nil {
		log.Fatalf("Failed to listen: %v", err1)
	}

	// ------------------------------------------------------------------------
	// Database
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),

		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "company",
		Params:               map[string]string{},
		AllowNativePasswords: true,
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to the database!")
	//--------------------------------------------------------------------------

	opts := []grpc.ServerOption{}

	s := grpc.NewServer(opts...)
	companypb.RegisterCompanyServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
