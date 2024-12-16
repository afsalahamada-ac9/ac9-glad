package presenter

import (
	"ac9/glad/pkg/glad"
	"time"

	"github.com/ulule/deepcopier"
)

type CenterWrapper struct {
	Object    string `json:"object"`
	Operation string `json:"operation"`
	Value     Center `json:"value"`
}

type Center struct {
	ExtID       string    `json:"Id"`
	ExtName     string    `json:"Name"`
	Name        string    `json:"Center_Name__c"`
	Address1    string    `json:"Street_Address_1__c"`
	Address2    string    `json:"Street_Address_2__c"`
	City        string    `json:"City__c"`
	State       string    `json:"State_Province__c"`
	PostalCode  string    `json:"Postal_Or_Zip_Code__c"`
	Country     string    `json:"Country__c"`
	Latitude    float64   `json:"Geolocation__Latitude__s"`
	Longitude   float64   `json:"Geolocation__Longitude__s"`
	MaxCapacity int32     `json:"Max_Capacity__c"`
	Mode        string    `json:"Center_Mode__c"`
	IsNational  bool      `json:"Is_National_Center__c"`
	IsEnabled   bool      `json:"Is_Enable__c"`
	CenterURL   string    `json:"Center_URL__c"`
	UpdatedAt   time.Time `json:"LastModifiedDate"`
	CreatedAt   time.Time `json:"CreatedDate"`
}

type CenterResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}

// ToGladCenter populates glad center using presenter center
func (c Center) ToGladCenter(gc *glad.Center) {
	deepcopier.Copy(c).To(gc)

	gc.Address.Street1 = c.Address1
	gc.Address.Street2 = c.Address2
	gc.Address.City = c.City
	gc.Address.State = c.State
	gc.Address.Zip = c.PostalCode
	gc.Address.Country = c.Country

	gc.GeoLocation.Latitude = c.Latitude
	gc.GeoLocation.Longitude = c.Longitude
}
