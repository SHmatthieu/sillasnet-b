package database

import (
	"database/sql"
	"fmt"
)

// User struct that represent User in db, if id=0 then the user isnt in the database
type User struct {
	Id           int64
	Name         string
	Token        string
	Hashpassword string
}

func (u *User) Equal(user *User) bool {
	return u.Id == user.Id && u.Name == user.Name
}

func UserFromRow(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.Id, &user.Name, &user.Token, &user.Hashpassword)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// try to add the user in db, return an error if it fails,
// if it worked, the user id is updated to match the id in db
func (u *User) AddInDB() error {
	result, err := db.Exec("INSERT INTO User (name, token, hashpassword) VALUES (?, ?, ?)", u.Name, u.Token, u.Hashpassword)
	if err != nil {
		return fmt.Errorf("addUser: %v", err)
	}
	id, _ := result.LastInsertId()
	u.Id = id
	return nil
}

func (u User) String() string {
	return fmt.Sprintf("Id : %d, Name : %s, Token : %s, Hashpassword : %s", u.Id, u.Name, u.Token, u.Hashpassword)
}

func (u *User) AddSoftware(software *Software) (int64, error) {
	result, err := db.Exec("INSERT INTO UserSoftware (userid, softwareid) VALUES (?, ?)", u.Id, software.Id)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (u *User) AddHardware(hardware *Hardware) (int64, error) {
	result, err := db.Exec("INSERT INTO UserHardware (userid, hardwareid) VALUES (?, ?)", u.Id, hardware.Id)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (u *User) AddSupplier(supplier *Supplier) (int64, error) {
	result, err := db.Exec("INSERT INTO UserSupplier (userid, supplierid) VALUES (?, ?)", u.Id, supplier.Id)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (u *User) GetSoftwares() ([]Software, error) {
	rows, err := db.Query("SELECT Software.id, Software.name, Software.version FROM (Software INNER JOIN UserSoftware on Software.id=UserSoftware.softwareid) WHERE UserSoftware.userid = ?", u.Id)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	var softwares []Software
	for rows.Next() {
		software := Software{}
		rows.Scan(&software.Id, &software.Name, &software.Version)
		softwares = append(softwares, software)
	}

	return softwares, nil
}

func (u *User) GetHarwares() ([]Hardware, error) {
	rows, err := db.Query("SELECT Hardware.id, Hardware.name, Hardware.version FROM (Hardware INNER JOIN UserHardware on Hardware.id=UserHardware.hardwareid) WHERE UserHardware.userid = ?", u.Id)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	var hardwares []Hardware
	for rows.Next() {
		hardware := Hardware{}
		rows.Scan(&hardware.Id, &hardware.Name, &hardware.Version)
		hardwares = append(hardwares, hardware)
	}

	return hardwares, nil
}

func (u *User) GetSuppliers() ([]Supplier, error) {
	rows, err := db.Query("SELECT Supplier.id, Supplier.name FROM (Supplier INNER JOIN UserSupplier on Supplier.id=UserSupplier.supplierid) WHERE UserSupplier.userid = ?", u.Id)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	var suppliers []Supplier
	for rows.Next() {
		supplier := Supplier{}
		rows.Scan(&supplier.Id, &supplier.Name)
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (u *User) DeleteFromDB() error {
	if u.Id == 0 {
		return fmt.Errorf("user not in db")
	}
	_, err := db.Exec("DELETE FROM UserSoftware WHERE userid = ?", u.Id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM UserHardware WHERE userid = ?", u.Id)
	if err != nil {
		return err
	}
	_, err = db.Exec("DELETE FROM UserSupplier WHERE userid = ?", u.Id)
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM User WHERE id = ?", u.Id)
	return err
}

func GetUserById(id int64) (*User, error) {
	row := db.QueryRow("SELECT * FROM User WHERE id = ?", id)
	return UserFromRow(row)
}

func GetUserByName(name string) (*User, error) {
	row := db.QueryRow("SELECT * FROM User WHERE name = ?", name)
	return UserFromRow(row)
}

func (u *User) GetSoftwareById(id int64) (*Software, error) {
	row := db.QueryRow("SELECT Software.id, Software.name, Software.version FROM (Software INNER JOIN UserSoftware on Software.id=UserSoftware.softwareid) WHERE UserSoftware.userid = ? AND UserSoftware.softwareid = ?", u.Id, id)
	return SoftwareFromRow(row)
}

func (u *User) GetHardwareById(id int64) (*Hardware, error) {
	row := db.QueryRow("SELECT Hardware.id, Hardware.name, Hardware.version FROM (Hardware INNER JOIN UserHardware on Hardware.id=UserHardware.hardwareid) WHERE UserHardware.userid = ? AND UserHardware.hardwareid = ?", u.Id, id)
	return HardwareFromRow(row)
}

func (u *User) GetSupplierById(id int64) (*Supplier, error) {
	row := db.QueryRow("SELECT Supplier.id, Supplier.name FROM (Supplier INNER JOIN UserSupplier on Supplier.id=UserSupplier.supplierid) WHERE UserSupplier.userid = ? AND UserSupplier.supplierid = ?", u.Id, id)
	return SupplierFromRow(row)
}
