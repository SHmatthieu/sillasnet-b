package server

import (
	"be/serv/database"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Datastructure use to decode JSON in connection request
type ConnectionRequestData struct {
	Name         string
	Hashpassword string
}

func (req ConnectionRequestData) String() string {
	return fmt.Sprintf("Name : %s, Hashpassword : %s", req.Name, req.Hashpassword)
}

// connection POST request handler that return a connection token
// if name and hash password are correct
func PostConnection(c *gin.Context) {
	fmt.Println("POSTED")
	row, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	var req ConnectionRequestData
	err = json.Unmarshal(row, &req)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request format",
		})
		return
	}

	fmt.Println(req)
	user, err := database.GetUserByName(req.Name)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error request data",
		})
		return
	}

	if user.Hashpassword != req.Hashpassword {
		c.JSON(400, gin.H{
			"message": "error request data",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "ok",
		"token":   user.Token,
	})
}
