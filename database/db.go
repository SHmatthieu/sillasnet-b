package database

import (
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mydb *gorm.DB = nil

type DbElement interface {
	Update()
}

type User struct {
	gorm.Model
	Name          string `gorm:"unique"`
	Token         string
	HashPassoword string
	Softwares     []*Software `gorm:"many2many:user_softwares;"`
	Hardwares     []*Hardware `gorm:"many2many:user_hardwares;"`
	Suppliers     []*Supplier `gorm:"many2many:user_suppliers;"`
}

type Software struct {
	gorm.Model
	Users   []*User `gorm:"many2many:user_softwares;"`
	Name    string  `gorm:"index:idx_name,unique"`
	Version string  `gorm:"index:idx_name,unique"`
}
type Hardware struct {
	gorm.Model
	Users   []*User `gorm:"many2many:user_hardwares;"`
	Name    string  `gorm:"index:idx_name,unique"`
	Version string  `gorm:"index:idx_name,unique"`
}
type Supplier struct {
	gorm.Model
	Users []*User `gorm:"many2many:user_suppliers;"`
	Name  string  `gorm:"index:idx_name,unique"`
}

func (user *User) AddSoftware(software *Software) {

	mydb.Model(user).Update("Softwares", append(user.Softwares, software))

}
func (user *User) Addhardware(hardware *Hardware) {

	mydb.Model(user).Update("Hardwares", append(user.Hardwares, hardware))

}
func (user *User) AddSupplier(supplier *Supplier) {

	mydb.Model(user).Update("Suppliers", append(user.Suppliers, supplier))

}
func InitDB() error {
	dsn := "usersec1:password@tcp(127.0.0.1:3306)/projecttest2?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	mydb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = mydb.AutoMigrate(&Software{}, &Hardware{}, &Supplier{}, &User{})
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(name string, password string) (*User, error) {

	tok := uuid.New().String()
	user := &User{Name: name, Token: tok, HashPassoword: password}
	result := mydb.Create(user)
	return user, result.Error
}
func GetUserByName(name string) (*User, error) {
	user := &User{}
	result := mydb.First(&user, "Name = ?", name)
	if result.Error != nil {
		return nil, result.Error
	}
	mydb.Model(user).Association("Softwares").Find(&user.Softwares)
	mydb.Model(user).Association("Hardwares").Find(&user.Hardwares)
	mydb.Model(user).Association("Suppliers").Find(&user.Suppliers)

	return user, nil
}
func GetUserByID(id int) (*User, error) {
	user := &User{}
	result := mydb.First(&user, "ID = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	mydb.Model(user).Association("Softwares").Find(&user.Softwares)
	mydb.Model(user).Association("Hardwares").Find(&user.Hardwares)
	mydb.Model(user).Association("Suppliers").Find(&user.Suppliers)
	return user, nil
}
func GetUserByToken(token string) (*User, error) {
	user := &User{}
	result := mydb.First(&user, "Token = ?", token)
	if result.Error != nil {
		return nil, result.Error
	}
	mydb.Model(user).Association("Softwares").Find(&user.Softwares)
	mydb.Model(user).Association("Hardwares").Find(&user.Hardwares)
	mydb.Model(user).Association("Suppliers").Find(&user.Suppliers)
	return user, nil
}

// Get a Software object , add it into the db if it doesnt already exist
func GetSoftware(name string, version string) *Software {
	soft := &Software{Name: name, Version: version}

	var existingSoft Software
	if err := mydb.Where(&Software{Name: name, Version: version}).First(&existingSoft).Error; err == nil {
		mydb.Model(&existingSoft).Association("Users").Find(&existingSoft.Users)
		return &existingSoft
	}

	mydb.Create(soft)
	return soft
}

// Get a Hardware object , add it into the db if it doesnt already exist
func GetHardware(name string, version string) *Hardware {
	hardware := &Hardware{Name: name, Version: version}

	var existingHardware Hardware
	if err := mydb.Where(&Hardware{Name: name, Version: version}).First(&existingHardware).Error; err == nil {
		mydb.Model(&existingHardware).Association("Users").Find(&existingHardware.Users)
		return &existingHardware
	}

	mydb.Create(hardware)
	return hardware
}

// Get a Supplier object , add it into the db if it doesnt already exist
func GetSupplier(name string) *Supplier {
	supplier := &Supplier{Name: name}

	var existingSupplier Supplier
	if err := mydb.Where(&Supplier{Name: name}).First(&existingSupplier).Error; err == nil {
		mydb.Model(&existingSupplier).Association("Users").Find(&existingSupplier.Users)
		return &existingSupplier
	}

	mydb.Create(supplier)
	return supplier
}
