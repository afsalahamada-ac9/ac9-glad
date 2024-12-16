package presenter

import (
	"ac9/glad/pkg/glad"
	"time"

	"github.com/ulule/deepcopier"
)

type CourseWrapper struct {
	Object    string `json:"object"`
	Operation string `json:"operation"`
	Value     Course `json:"value"`
}

type Course struct {
	ExtID        string `json:"Id"`
	CenterExtID  string `json:"Location__c"`
	ProductExtID string `json:"Workshop_Type__c"`
	// TODO: Maybe, it's not name, but something else such as Title?
	Name         string    `json:"Name"`
	Notes        string    `json:"Notes__c"`
	Timezone     string    `json:"Timezone__c"`
	Address1     string    `json:"Street_Address_1__c"`
	Address2     string    `json:"Street_Address_2__c"`
	City         string    `json:"City__c"`
	State        string    `json:"State__c"`
	PostalCode   string    `json:"Zip_Postal_Code__c"`
	Country      string    `json:"Country__c"`
	Status       string    `json:"Status__c"`
	Mode         string    `json:"Online_Or_In_Person__c"`
	MaxAttendees int       `json:"Max_attendees__c"`
	NumAttendees int       `json:"Number_Of_Students__c"`
	URL          string    `json:"Journey_App_Url_Details__c"`
	CheckoutURL  string    `json:"Journey_App_Url__c"`
	CreatedAt    time.Time `json:"CreatedDate"`
	UpdatedAt    time.Time `json:"LastModifiedDate"`
}

type CourseResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}

// ToGladCourse populates glad center using presenter center
func (c Course) ToGladCourse(gc *glad.Course) {
	deepcopier.Copy(c).To(gc)

	gc.Address.Street1 = c.Address1
	gc.Address.Street2 = c.Address2
	gc.Address.City = c.City
	gc.Address.State = c.State
	gc.Address.Zip = c.PostalCode
	gc.Address.Country = c.Country
}
