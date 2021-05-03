package models

import (
	"time"
)

type Customer struct {
	Id                      uint      `JSON: "id" gorm:"primary_key"`
	FirstName               string    `JSON: "firstName"`
	LastName                string    `JSON: "lastName"`
	LegalEntityStage        string    `JSON: "legalEntityStage"`
	LegalEntityType         string    `JSON: "legalEntityType"`
	CompanyName             string    `JSON: "companyName"`
	BankruptcyIndicatorFlag bool      `JSON:"bankruptcyIndicatorFlag"`
	CreatedDate             time.Time `JSON:"createdDate"`
	DateOfBirth             time.Time `JSON:"dateOfBirth"`
}
