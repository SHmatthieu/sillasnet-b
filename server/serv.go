// this package is the REST api for the project

package server

import "github.com/gin-gonic/gin"

func StartServer() {
	router := gin.Default()

	router.POST("/api/connection", PostConnection)

	router.GET("/api/software", GetSoftware)
	router.GET("/api/software/:id", GetSoftwareId)
	router.POST("/api/software", PostSoftware)

	router.GET("/api/hardware", GetHardware)
	router.GET("/api/hardware/:id", GetHardwareId)
	router.POST("/api/hardware", PostHardware)

	router.GET("/api/supplier", GetSupplier)
	router.GET("/api/supplier/:id", GetSupplierById)
	router.POST("/api/supplier", PostSupplier)

	router.Run()
}
