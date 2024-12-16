/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/pkg/glad"
	"time"

	"github.com/ulule/deepcopier"
)

type Account struct {
	ExtID        string `json:"Id"`
	CognitoID    string `json:"Cognito_User_Id__c"`
	Username     string `json:"Name"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	Phone        string `json:"Phone"`
	Email        string `json:"PersonEmail"`
	Type         string `json:"Account_Type__c"`
	Status       string `json:"User_status__pc"`
	FullPhotoURL string `json:"FullPhotoURL"`
	// TODO: ctypes support to be added later
	CType          string    `json:"CType_Id__c"`
	AssistantCType string    `json:"CTypeId_Eligibility_As_Assistant_Teacher__c"`
	UpdatedAt      time.Time `json:"LastModifiedDate"`
	CreatedAt      time.Time `json:"CreatedDate"`
}

type AccountResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}

// ToGladAccount populates glad account using presenter account
func (c Account) ToGladAccount(gc *glad.Account) {
	deepcopier.Copy(c).To(gc)
}
