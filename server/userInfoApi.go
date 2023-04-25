package server

import (
	"be/serv/database"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ERROR_REQUEST_MSG string = "error request"
var ERROR_FORMAT_MSG string = "error request format"
var ERROR_ID_MSG string = "error id"

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
	ID      uint
	Name    string
	Version string
}

type PostHardwareRequest struct {
	Name      string
	Token     string
	Hardwares []Hardware
}

type Hardware struct {
	ID      uint
	Name    string
	Version string
}

type PostSupplierRequest struct {
	Name      string
	Token     string
	Suppliers []Supplier
}

type PostReportRequest struct {
	Text string
}

type Supplier struct {
	ID   uint
	Name string
}

func GetUserFromRequest(c *gin.Context) (*database.User, error) {

	name := c.Request.Header.Get("name")
	token := c.Request.Header.Get("token")

	user, err := database.GetUserByName(name)
	if err != nil {
		return nil, errors.New(ERROR_REQUEST_MSG)

	}

	if user.Token != token {
		return nil, errors.New(ERROR_REQUEST_MSG)

	}

	return user, nil
}

func GetSoftware(c *gin.Context) {

	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	softwares := make([]Software, len(user.Softwares))
	for i := range user.Softwares {
		softwares[i].Name = user.Softwares[i].Name
		softwares[i].Version = user.Softwares[i].Version
		softwares[i].ID = user.Softwares[i].ID

	}
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
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_ID_MSG,
		})
		return
	}
	for _, software := range user.Softwares {
		if software.ID == uint(id) {
			c.JSON(200, gin.H{
				"message":  "ok",
				"software": Software{Name: software.Name, Version: software.Version, ID: software.ID},
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"message": "software doesnt exist",
	})
}

func PostSoftware(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
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
		dbSoft := database.GetSoftware(software.Name, software.Version)
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
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	hardwares := make([]Hardware, len(user.Hardwares))
	for i := range user.Hardwares {
		hardwares[i].Name = user.Hardwares[i].Name
		hardwares[i].Version = user.Hardwares[i].Version
		hardwares[i].ID = user.Hardwares[i].ID

	}
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
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_ID_MSG,
		})
		return
	}
	for _, hardware := range user.Hardwares {
		if hardware.ID == uint(id) {
			c.JSON(200, gin.H{
				"message":  "ok",
				"hardware": Hardware{Name: hardware.Name, Version: hardware.Version, ID: hardware.ID},
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"message": "hardware doesnt exist",
	})
}

func PostHardware(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
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
		dbHard := database.GetHardware(hardware.Name, hardware.Version)
		if err != nil {
			fmt.Printf("%v\n", err)
			c.JSON(400, gin.H{
				"message": fmt.Sprintf("error hardware %d", i),
			})
			return
		}
		user.Addhardware(dbHard)
	}

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetSupplier(c *gin.Context) {
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	suppliers := make([]Supplier, len(user.Suppliers))
	for i := range user.Suppliers {
		suppliers[i].Name = user.Suppliers[i].Name
		suppliers[i].ID = user.Suppliers[i].ID

	}
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
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_ID_MSG,
		})
		return
	}
	for _, supplier := range user.Suppliers {
		if supplier.ID == uint(id) {
			c.JSON(200, gin.H{
				"message":  "ok",
				"supplier": Software{Name: supplier.Name, ID: supplier.ID},
			})
			return
		}
	}
	c.JSON(400, gin.H{
		"message": "supplier doesnt exist",
	})
}

func PostSupplier(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
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
		dbSupplier := database.GetSupplier(supplier.Name)
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

func PostReport(c *gin.Context) {
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}
	user, err := GetUserFromRequest(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": ERROR_FORMAT_MSG,
		})
		return
	}

	var req PostReportRequest
	err = json.Unmarshal(row, &req)
	if err != nil {
		fmt.Printf("%v\n", err)
		c.JSON(400, gin.H{
			"message": "error supplier format",
		})
		return
	}

	report := database.GetReport(req.Text)
	user.AddReport(report)

	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func GetTips(c *gin.Context) {
	dbtips := database.GetAllTips()
	tips := make([]string, 0)
	for _, tip := range dbtips {
		if len(tip.Text) > 0 {
			fmt.Println(len(tip.Text))
			tips = append(tips, tip.Text)
		}
	}

	c.JSON(200, gin.H{
		"Tips": tips,
	})
}
