package Controllers

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCustomerInput struct {
	FirstName        string `JSON: "firstName"  binding:"required"`
	LastName         string `JSON: "lastName"  binding:"required"`
	LegalEntityStage string `JSON: "legalEntityStage"  binding:"required"`
	LegalEntityType  string `JSON: "legalEntityType"  binding:"required"`
	CompanyName      string `JSON: "companyName"  binding:"required"`
}

type UpdateCustomerInput struct {
	FirstName        string `JSON: "firstName"  binding:"required"`
	LastName         string `JSON: "lastName"  binding:"required"`
	LegalEntityStage string `JSON: "legalEntityStage"  binding:"required"`
	LegalEntityType  string `JSON: "legalEntityType"  binding:"required"`
	CompanyName      string `JSON: "companyName"  binding:"required"`
}

type SearchCustomerInput struct {
	FirstName   string `JSON: "firstName"  binding:"required"`
	LastName    string `JSON: "lastName"  binding:"required"`
	CompanyName string `JSON: "companyName"  binding:"required"`
}

func FetchAllCustomers(c *gin.Context) {
	var customers []models.Customer
	models.DB.Find(&customers)

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func SearchCustomers(c *gin.Context) {
	// Validate input
	var input SearchCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var customers []models.Customer
	if err := models.DB.Where("first_name = ? OR last_name = ? OR company_name = ?", input.FirstName, input.LastName, input.CompanyName).Find(&customers).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customers})
}

func GetById(c *gin.Context) {

	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).Find(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func CreateCustomer(c *gin.Context) {
	// Validate input
	var input CreateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	customer := models.Customer{
		FirstName:               input.FirstName,
		LastName:                input.LastName,
		LegalEntityStage:        input.LegalEntityStage,
		LegalEntityType:         input.LegalEntityType,
		CompanyName:             input.CompanyName,
		BankruptcyIndicatorFlag: false}

	models.DB.Create(&customer)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func UpdateCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateCustomerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&customer).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func DeleteCustomer(c *gin.Context) {
	// Get model if exist
	var customer models.Customer
	if err := models.DB.Where("id = ?", c.Param("id")).First(&customer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&customer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
