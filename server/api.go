package server

import (
	"be/serv/database"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetRequestData struct {
	Name  string
	Token string
}

type PostSoftwareRequest struct {
	Name      string
	Token     string
	Softwares []Software
}

type Software struct {
	Name    string
	Version string
}

type PostHardwareRequest struct {
	Name      string
	Token     string
	Hardwares []Hardware
}

type Hardware struct {
	Name    string
	Version string
}

type PostSupplierRequest struct {
	Name      string
	Token     string
	Suppliers []Supplier
}

type Supplier struct {
	Name string
}

func GetUserFromRequest(row []byte) (*database.User, error) {

	var req GetRequestData
	err := json.Unmarshal(row, &req)
	if err != nil {
		return nil, errors.New("error request")

	}

	user, err := database.GetUserByName(req.Name)
	if err != nil {
		return nil, errors.New("error request")

	}

	if user.Token != req.Token {
		return nil, errors.New("error request")

	}

	return user, nil
}

func GetSoftware(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	softwares, err := user.GetSoftwares()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error software list",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "ok",
		"software": softwares,
	})
}

func GetSoftwareId(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error id",
		})
		return
	}
	software, err := user.GetSoftwareById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "software doesnt exist",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "ok",
		"software": software,
	})
}

func PostSoftware(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	var req PostSoftwareRequest
	err = json.Unmarshal(row, &req)
	if err != nil {
		fmt.Printf("%v\n", err)
		c.JSON(400, gin.H{
			"message": "error software format",
		})
		return
	}

	for i, software := range req.Softwares {
		dbSoft := &database.Software{Name: software.Name, Version: software.Version}
		err := dbSoft.AddInDB()
		if err != nil {
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("error software %d", i),
			})
			return
		}
		user.AddSoftware(dbSoft)
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetHardware(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	hardwares, err := user.GetHarwares()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error hardware list",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "ok",
		"hardware": hardwares,
	})
}

func GetHardwareId(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error id",
		})
		return
	}
	hardware, err := user.GetHardwareById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "hardware doesnt exist",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "ok",
		"hardware": hardware,
	})
}

func PostHardware(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	var req PostHardwareRequest
	err = json.Unmarshal(row, &req)
	if err != nil {
		fmt.Printf("%v\n", err)
		c.JSON(400, gin.H{
			"message": "error hardware format",
		})
		return
	}

	for i, hardware := range req.Hardwares {
		dbHard := &database.Hardware{Name: hardware.Name, Version: hardware.Version}
		err := dbHard.AddInDB()
		if err != nil {
			fmt.Printf("%v\n", err)
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("error hardware %d", i),
			})
			return
		}
		user.AddHardware(dbHard)
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetSupplier(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	suppliers, err := user.GetSuppliers()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error suppliers list",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":   "ok",
		"suppliers": suppliers,
	})
}

func GetSupplierById(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error id",
		})
		return
	}
	supplier, err := user.GetSupplierById(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "supplier doesnt exist",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "ok",
		"supplier": supplier,
	})
}

func PostSupplier(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}
	user, err := GetUserFromRequest(row)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	var req PostSupplierRequest
	err = json.Unmarshal(row, &req)
	if err != nil {
		fmt.Printf("%v\n", err)
		c.JSON(400, gin.H{
			"message": "error supplier format",
		})
		return
	}

	for i, supplier := range req.Suppliers {
		dbSupplier := &database.Supplier{Name: supplier.Name}
		err := dbSupplier.AddInDB()
		if err != nil {
			fmt.Printf("%v\n", err)
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("error supplier %d", i),
			})
			return
		}
		user.AddSupplier(dbSupplier)
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}
