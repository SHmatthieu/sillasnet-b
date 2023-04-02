package database

import (
	"database/sql"
	"fmt"
)

// Hardware struct that represent User in db, if id=0 then the hardware isnt in the database
type Hardware struct {
	Id      int64
	Name    string
	Version string
}

func (h *Hardware) Equal(hardware *Hardware) bool {
	return h.Id == hardware.Id && h.Name == hardware.Name && h.Version == hardware.Version
}

func (h *Hardware) String() string {
	return fmt.Sprintf("Id : %d, Name : %s, Version : %s", h.Id, h.Name, h.Version)
}

func HardwareFromRow(row *sql.Row) (*Hardware, error) {
	hardware := Hardware{}
	err := row.Scan(&hardware.Id, &hardware.Name, &hardware.Version)
	return &hardware, err
}

// try to add the hardware in db, return an error if it fails,
// if it worked, the hardware id is updated to match the id in db
// if the hardware is already in the db the the id is also updated
func (s *Hardware) AddInDB() error {
	hardware, err := GetHardwareByNameAndVersion(s.Name, s.Version)
	if err == nil {
		s.Id = hardware.Id
		return nil
	}

	result, err := db.Exec("INSERT INTO Hardware (name, version) VALUES (?, ?)", s.Name, s.Version)
	if err != nil {
		return fmt.Errorf("addHardware: %v", err)
	}
	id, _ := result.LastInsertId()
	s.Id = id
	return nil
}

func (h *Hardware) DeleteFromDB() error {
	if h.Id == 0 {
		return fmt.Errorf("Hardware not in db")
	}
	_, err := db.Exec("DELETE FROM UserHardware WHERE hardwareid = ?", h.Id)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM Hardware WHERE id = ?", h.Id)
	return err
}

func GetHardwareById(id int64) (*Hardware, error) {
	row := db.QueryRow("SELECT * FROM Hardware WHERE id = ?", id)
	return HardwareFromRow(row)
}

func GetHardwareByNameAndVersion(name string, version string) (*Hardware, error) {
	row := db.QueryRow("SELECT * FROM Hardware WHERE name = ? and version = ?", name, version)
	return HardwareFromRow(row)
}
