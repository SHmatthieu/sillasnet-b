package database

import (
	"database/sql"
	"fmt"
)

// Supplier struct that represent User in db, if id=0 then the Supplier isnt in the database
type Supplier struct {
	Id   int64
	Name string
}

func (s *Supplier) Equal(supplier *Supplier) bool {
	return s.Id == supplier.Id && s.Name == supplier.Name
}

func (s *Supplier) String() string {
	return fmt.Sprintf("Id : %d, Name : %s", s.Id, s.Name)
}

func SupplierFromRow(row *sql.Row) (*Supplier, error) {
	supplier := Supplier{}
	err := row.Scan(&supplier.Id, &supplier.Name)
	return &supplier, err
}

// try to add the supplier in db, return an error if it fails,
// if it worked, the supplier id is updated to match the id in db
// if the supplier is already in the db the the id is also updated
func (s *Supplier) AddInDB() error {
	result, err := db.Exec("INSERT INTO Supplier (name) VALUES (?)", s.Name)
	if err != nil {
		return fmt.Errorf("addSupplier: %v", err)
	}
	id, _ := result.LastInsertId()
	s.Id = id
	return nil
}

func (s *Supplier) DeleteFromDB() error {
	if s.Id == 0 {
		return fmt.Errorf("Supplier not in db")
	}

	_, err := db.Exec("DELETE FROM UserSupplier WHERE supplierid = ?", s.Id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM Supplier WHERE id = ?", s.Id)
	return err
}

func GetSupplierById(id int64) (*Supplier, error) {
	row := db.QueryRow("SELECT * FROM Supplier WHERE id = ?", id)
	return SupplierFromRow(row)
}

func GetSupplierByName(name string) (*Supplier, error) {
	row := db.QueryRow("SELECT * FROM Supplier WHERE name = ?", name)
	return SupplierFromRow(row)
}
