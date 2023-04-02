package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func main() {
	fmt.Println("main !")
	dsn := "usersec1:password@tcp(127.0.0.1:3306)/projecttest2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Software{}, &Hardware{}, &Supplier{}, &User{})

	//db.Create(&User{Name: "matthieu2", Token: "toktok", HashPassoword: "hashshsh"})
	var mat User
	db.First(&mat, "Name = ?", "matthieu2")
	//db.Create(&Software{Name: "soft2", Version: "2.0", Users: []*User{&mat}})

	var softwares []*Software
	db.Model(&mat).Association("Softwares").Find(&softwares)

	for _, s := range softwares {
		fmt.Println(s)
	}
}
