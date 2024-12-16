/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"strings"
	"time"
)

// Account type
type AccountType string

const (
	AccountTeacher          AccountType = "teacher"
	AccountAssistantTeacher AccountType = "assistant-teacher"
	AccountOrganizer        AccountType = "organizer"
	AccountMember           AccountType = "member"
	AccountUser             AccountType = "user"
	AccountCoOrdinator      AccountType = "coordinator"
	AccountVolunteer        AccountType = "volunteer"
	AccountStudent          AccountType = "student"
	AccountOther            AccountType = "other"
	// Add new types here
)

// Account status
type AccountStatus string

const (
	AccountActive   AccountStatus = "active"
	AccountInactive AccountStatus = "inactive"
	AccountDisabled AccountStatus = "disabled"
	// Add new types here
)

// Account data
type Account struct {
	ID        id.ID
	TenantID  id.ID
	ExtID     string
	CognitoID string

	Username     string
	FirstName    string
	LastName     string
	Phone        string
	Email        string
	Type         AccountType
	Status       AccountStatus
	FullPhotoURL string

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAccount creates a new account
func NewAccount(tenantID id.ID,
	cognitoID string,
	username string,
	first_name string,
	last_name string,
	phone string,
	email string,
	at AccountType,
	as AccountStatus,
) (*Account, error) {
	t := &Account{
		ID:        id.New(),
		TenantID:  tenantID,
		CognitoID: cognitoID,
		Username:  username,
		FirstName: first_name,
		LastName:  last_name,
		Phone:     phone,
		Email:     email,
		Type:      at,
		Status:    as,
		CreatedAt: time.Now(),
	}
	err := t.Validate()
	if err != nil {
		return nil, glad.ErrInvalidEntity
	}
	return t, nil
}

// Transform fixes the data issues, if any
func (a *Account) Transform() {
	l.Log.Debugf("Before transform, account=%#v", a)
	a.Type = AccountType(strings.ToLower(string(a.Type)))

	// TODO: Need to figure out what these (other, co-ord) types are
	switch string(a.Type) {
	case "co-ord":
		a.Type = AccountCoOrdinator
	case "assistant teacher":
		a.Type = AccountAssistantTeacher
	default:
		// TODO: Add a metric to keep track of this
		a.Type = AccountUser
	}

	a.Status = AccountStatus(strings.ToLower(string(a.Status)))
	l.Log.Debugf("After transform, account=%#v", a)
}

// Validate validate account
func (a *Account) Validate() error {
	if a.TenantID == id.IDInvalid {
		l.Log.Warnf("Invalid tenant id=%v, product extID=%v", a.TenantID, a.ExtID)
		return glad.ErrInvalidEntity
	}
	if a.Username == "" {
		l.Log.Warnf("Account extID=%v, has empty username", a.ExtID)
		return glad.ErrInvalidEntity
	}

	return nil
}
