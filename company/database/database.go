package database

import (
	"database/sql"
	"fmt"

	"example.com/grpc-sql/company/companypb"
	"example.com/grpc-sql/company/views"
)

// SelectRow queries for the album with the specified ID.
func SelectRow(id int64, db *sql.DB) (*companypb.Company, error) {
	// A Company to hold data from the returned row.
	com := views.Comp{}
	row := db.QueryRow("SELECT * FROM company WHERE id = ?", id)

	err := row.Scan(&com.ID, &com.Name, &com.Creator)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("CompanyByID %d: no such Company", id)
		}
		return nil, fmt.Errorf("CompanyByID %d: %v", id, err)
	}

	res := com.ToCompany()
	return res, nil
}

//InsertRow inserts a company into the table
func InsertRow(com views.Comp, db *sql.DB) (int64, error) {

	//Inserting a new row into the table
	result, err := db.Exec("INSERT INTO company (name, creator) VALUES (?, ?)", com.Name, com.Creator)
	if err != nil {
		return 0, fmt.Errorf("AddCompany: %v", err)
	}
	// Get the new company generated ID for the client.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("AddCompany: %v", err)
	}
	// Return the new company's ID.
	return id, nil

}

//UpdateRow updates a row in the table
func UpdateRow(update views.Comp, db *sql.DB) (*companypb.UpdateResponse, error) {

	//getting the actual row
	actual_row, err1 := SelectRow(update.ID, db)
	if err1 != nil {
		return nil, fmt.Errorf("Error while getting the row: %v", err1)
	}

	//updating the row values
	_, err := db.Exec("UPDATE company SET name=? , creator=? WHERE id=?", update.Name, update.Creator, update.ID)
	if err != nil {
		return nil, fmt.Errorf("UpdateCompany: %v", err)
	}

	res := &companypb.UpdateResponse{
		Company: actual_row,
	}
	return res, nil
}

// DeleteRow is used to delete a row from table
func DeleteRow(req *companypb.DeleteRequest, db *sql.DB) (*companypb.DeleteResponse, error) {
	//getting the actual row
	actual_row, err1 := SelectRow(req.GetId(), db)
	if err1 != nil {
		return nil, fmt.Errorf("Error while getting the row: %v", err1)
	}

	//deleting the row
	_, err := db.Exec("DELETE FROM company WHERE id=?", req.Id)
	if err != nil {
		return nil, fmt.Errorf("DeleteRow: %v", err)
	}

	res := &companypb.DeleteResponse{
		Company: actual_row,
	}
	return res, nil
}
