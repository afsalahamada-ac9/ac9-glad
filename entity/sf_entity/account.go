package entity

import "strings"

type Account struct {
	Operation string        `json:"operation" bson:"operation"`
	Value     Account_value `json:"value" bson:"value"`
	Object    string        `json:"object" bson:"object"`
}

type Account_value struct {
	Ext_Id     string `json:"Id" gorm:"column:ext_id"`
	Tenant_Id  int    `json:"Tenant_id" gorm:"column:tenant_id"`
	Cognito_Id string `json:"Cognito_User_Id__c" gorm:"column:cognito_id"`
	Name       string `json:"Name" gorm:"column:username"`
	First_Name string `json:"FirstName" gorm:"column:first_name"`
	Last_Name  string `json:"LastName" gorm:"column:last_name"`
	Phone      string `json:"Phone" gorm:"column:phone"`
	Email      string `json:"PersonEmail" gorm:"column:email"`
	Type       string `json:"Account_Type__c" gorm:"column:type"`
	Updated_at string `json:"LastModifiedDate" gorm:"column:updated_at"`
	// User_status string `json:"User_status__c" gorm:"column:user_status"`
	Created_at string `json:"CreatedDate" gorm:"column:created_at"`
}

func (*Account_value) TableName() string {
	return "account"
}

// todo: check timezone, which of the two dbs(rds and sf) will be saved in the rds?
// note: we've to store both the times in separate fields, calculate their difference, and then we'll decide on it. UTC time zone standard.

func (*Account) NewAccount(Id string, Tenant_id int, Cognito_id string, Name string, Fname string, Lname string, Phone string, Email string, Type string, Updated_at string, Created_at string) *Account_value {
	return &Account_value{Ext_Id: Id, Tenant_Id: Tenant_id, Cognito_Id: Cognito_id, Name: Name, First_Name: Fname, Last_Name: Lname, Phone: Phone, Email: Email, Type: strings.ToLower(Type), Updated_at: Updated_at, Created_at: Created_at}
}
