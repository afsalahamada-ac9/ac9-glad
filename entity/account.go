/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package entity

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
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

	Username  string
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Type      AccountType
	Status    AccountStatus

	// meta data
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAccount create a new account
func NewAccount(tenantID id.ID,
	extID string,
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
		ExtID:     extID,
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

// Validate validate account
func (t *Account) Validate() error {
	if t.Username == "" {
		return glad.ErrInvalidEntity
	}

	return nil
}
