package database

import (
	"database/sql"
	"fmt"
)

// Software struct that represent User in db, if id=0 then the software isnt in the database
type Software struct {
	Id      int64
	Name    string
	Version string
}

func (s *Software) Equal(soft *Software) bool {
	return s.Id == soft.Id && s.Name == soft.Name && s.Version == soft.Version
}

func (s Software) String() string {
	return fmt.Sprintf("Id : %d, Name : %s, Version : %s", s.Id, s.Name, s.Version)
}

func SoftwareFromRow(row *sql.Row) (*Software, error) {
	software := Software{}
	err := row.Scan(&software.Id, &software.Name, &software.Version)
	return &software, err
}

// try to add the software in db, return an error if it fails,
// if it worked, the software id is updated to match the id in db
func (s *Software) AddInDB() error {
	software, err := GetSoftwareByNameAndVersion(s.Name, s.Version)
	if err == nil {
		s.Id = software.Id
		return nil
	}
	result, err := db.Exec("INSERT INTO Software (name, version) VALUES (?, ?)", s.Name, s.Version)
	if err != nil {
		return fmt.Errorf("addSoftware: %v", err)
	}
	id, _ := result.LastInsertId()
	s.Id = id
	return nil
}

func (s *Software) DeleteFromDB() error {
	if s.Id == 0 {
		return fmt.Errorf("software not in db")
	}
	_, err := db.Exec("DELETE FROM UserSoftware WHERE softwareid = ?", s.Id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM Software WHERE id = ?", s.Id)
	return err
}

func GetSoftwareById(id int64) (*Software, error) {
	row := db.QueryRow("SELECT * FROM Software WHERE id = ?", id)
	return SoftwareFromRow(row)
}

func GetSoftwareByNameAndVersion(name string, version string) (*Software, error) {
	row := db.QueryRow("SELECT * FROM Software WHERE name = ? and version = ?", name, version)
	return SoftwareFromRow(row)
}
