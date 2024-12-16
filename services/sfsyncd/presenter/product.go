/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/pkg/id"
	"time"
)

// Note: There is no need to have this wrapper if salesforce can use HTTP verbs
// and paths. Salesforce already uses different paths, so verb can be easily added
// and wrapper can be get rid of.
type ProductWrapper struct {
	Object    string  `json:"object"`
	Operation string  `json:"operation"`
	Value     Product `json:"value"`
}

// Master entity from Salesforce maps to Product entity in GLAD
type Product struct {
	ExtID         string    `json:"Id"`
	ExtName       string    `json:"Name"`
	TenantID      id.ID     `json:"Tenant_id"`
	BaseProductID string    `json:"Product__c,omitempty"`
	CType         string    `json:"CType_Id__c"`
	Title         string    `json:"Title__c"`
	Format        string    `json:"Online_Or_In_Person__c"`
	MaxAttendees  int32     `json:"Max_Attendees__c,omitempty"`
	Visibility    string    `json:"Listing_Visibity__c,omitempty"`
	DurationDays  int32     `json:"Event_Duration__c,omitempty"`
	IsAutoApprove bool      `json:"Auto_Approve_Event__c"`
	UpdatedAt     time.Time `json:"LastModifiedDate"`
	CreatedAt     time.Time `json:"CreatedDate"`
}

type ProductResponse struct {
	ID      int64  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}
