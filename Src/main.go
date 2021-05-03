package main

import (
	"main/Controllers"
	"main/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.Setup()

	r.GET("/Customers", Controllers.FetchAllCustomers)
	r.GET("/GetById/:id", Controllers.GetById)
	r.POST("/CreateCustomer", Controllers.CreateCustomer)
	r.POST("/SearchCustomer", Controllers.SearchCustomers)
	r.PATCH("/UpdateById/:id", Controllers.UpdateCustomer)
	r.DELETE("/Delete/:id", Controllers.DeleteCustomer)

	r.Run()
}
