package database

import (
	"testing"
)

func TestDBConnection(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}
	db.Close()
}

func TestUser(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}

	user := &User{Name: "Test1", Token: "token", Hashpassword: "hash"}
	err = user.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newuser, err := GetUserById(user.Id)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !user.Equal(newuser) {
		t.Fatalf("%s", "Error User")
	}

	err = user.DeleteFromDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	db.Close()

}

func TestSoftware(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}

	soft := &Software{Name: "SoftTest1", Version: "3.0"}
	err = soft.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newsoft, err := GetSoftwareById(soft.Id)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !soft.Equal(newsoft) {
		t.Fatalf("%s", "Error User")
	}

	err = soft.DeleteFromDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	db.Close()
}

func TestHardware(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}

	hard := &Hardware{Name: "HardTest1", Version: "3.0"}
	err = hard.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newhard, err := GetHardwareById(hard.Id)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !hard.Equal(newhard) {
		t.Fatalf("%s", "Error User")
	}

	err = hard.DeleteFromDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	db.Close()
}

func TestSupplier(t *testing.T) {
	db, err := DBConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}

	supplier := &Supplier{Name: "SupplierTest1"}
	err = supplier.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	newSupplier, err := GetSupplierById(supplier.Id)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if !supplier.Equal(newSupplier) {
		t.Fatalf("%s", "Error User")
	}

	err = supplier.DeleteFromDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	db.Close()
}

func TestRelation(t *testing.T) {
	_, err := DBConnection()
	if err != nil {
		t.Fatalf("%v", err)
	}

	user := &User{Name: "Test2", Token: "token", Hashpassword: "hash"}
	err = user.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	soft1 := &Software{Name: "SoftTest2", Version: "3.0"}
	err = soft1.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	soft2 := &Software{Name: "SoftTest3", Version: "3.0"}
	err = soft2.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, err = user.AddSoftware(soft1)
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = user.AddSoftware(soft2)
	if err != nil {
		t.Fatalf("%v", err)
	}

	softwares, err := user.GetSoftwares()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(softwares) != 2 {
		t.Fatalf("Error softwares len size")
	}
	if !softwares[0].Equal(soft1) || !softwares[1].Equal((soft2)) {
		t.Fatalf("Error softwares")

	}

	hard1 := &Hardware{Name: "HardwareTest2", Version: "3.0"}
	err = hard1.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = user.AddHardware(hard1)
	if err != nil {
		t.Fatalf("%v", err)
	}

	hardwares, err := user.GetHarwares()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(hardwares) != 1 {
		t.Fatalf("Error hardwares len size")
	}
	if !hardwares[0].Equal(hard1) {
		t.Fatalf("Error hardwares")
	}

	supplier1 := &Supplier{Name: "SupplierTest2"}
	err = supplier1.AddInDB()
	if err != nil {
		t.Fatalf("%v", err)
	}
	_, err = user.AddSupplier(supplier1)
	if err != nil {
		t.Fatalf("%v", err)
	}

	suppliers, err := user.GetSuppliers()
	if err != nil {
		t.Fatalf("%v", err)
	}
	if len(suppliers) != 1 {
		t.Fatalf("Error suppliers len size")
	}
	if !suppliers[0].Equal(supplier1) {
		t.Fatalf("Error suppliers")
	}

	user.DeleteFromDB()
	soft1.DeleteFromDB()
	soft2.DeleteFromDB()
	hard1.DeleteFromDB()
	supplier1.DeleteFromDB()
	db.Close()

}
