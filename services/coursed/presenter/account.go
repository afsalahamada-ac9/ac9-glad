/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package presenter

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"

	"github.com/ulule/deepcopier"
)

// Account data - TenantID is returned in the HTTP header (may be not, as account is global?)
// X-GLAD-TenantID
type Account struct {
	ID        id.ID              `json:"id"`
	Username  string             `json:"username"`
	FirstName string             `json:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty"`
	Phone     string             `json:"phone,omitempty"`
	Email     string             `json:"email,omitempty"`
	Type      entity.AccountType `json:"type"`
	CognitoID string             `json:"cognitoID,omitempty"`
}

type AccountImportResponse struct {
	ID      id.ID  `json:"id"`
	ExtID   string `json:"extID"`
	IsError bool   `json:"isError"`
}

// FromAccountEntity creates account response from account entity
func (c *Account) FromAccountEntity(e *entity.Account) error {
	deepcopier.Copy(e).To(c)
	return nil
}

// GladAccountToEntity populates account entity from glad account object
func GladAccountToEntity(ga glad.Account, e *entity.Account) error {
	deepcopier.Copy(ga).To(e)

	e.Type = entity.AccountType(ga.Type)
	e.Status = entity.AccountStatus(ga.Status)

	l.Log.Infof("Account=%#v, entity.account=%#v", ga, e)
	return nil
}
