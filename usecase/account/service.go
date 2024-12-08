/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package account

import (
	"time"

	"ac9/glad/entity"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
)

// Service account usecase
type Service struct {
	repo Repository
}

// NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// CreateAccount creates an account
func (s *Service) CreateAccount(
	tenantID id.ID,
	extID string,
	cognitoID string,
	username string,
	first_name string,
	last_name string,
	phone string,
	email string,
	at entity.AccountType,
	as entity.AccountStatus,
) error {
	account, err := entity.NewAccount(tenantID,
		extID,
		cognitoID,
		username,
		first_name,
		last_name,
		phone,
		email,
		at,
		as,
	)
	if err != nil {
		return err
	}
	return s.repo.Create(account)
}

// GetAccount retrieves an account
func (s *Service) GetAccount(id id.ID) (*entity.Account, error) {
	account, err := s.repo.Get(id)
	if account == nil {
		return nil, glad.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}

// GetAccountByName retrieves an account using username
func (s *Service) GetAccountByName(tenantID id.ID, username string) (*entity.Account, error) {
	account, err := s.repo.GetByName(tenantID, username)
	if account == nil {
		return nil, glad.ErrNotFound
	}
	if err != nil {
		return nil, err
	}

	return account, nil
}

// ListAccounts list accounts
func (s *Service) ListAccounts(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	accounts, err := s.repo.List(tenantID, page, limit, at)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, glad.ErrNotFound
	}
	return accounts, nil
}

// UpdateAccount Update a account
func (s *Service) UpdateAccount(t *entity.Account) error {
	err := t.Validate()
	if err != nil {
		return err
	}
	t.UpdatedAt = time.Now()
	return s.repo.Update(t)
}

// DeleteAccount Deletes an account
func (s *Service) DeleteAccount(id id.ID) error {
	account, err := s.GetAccount(id)
	if account == nil {
		return glad.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// DeleteAccount Deletes an account using username
func (s *Service) DeleteAccountByName(tenantID id.ID, username string) error {
	account, err := s.GetAccountByName(tenantID, username)
	if account == nil {
		return glad.ErrNotFound
	}
	if err != nil {
		return err
	}

	return s.repo.DeleteByName(tenantID, username)
}

// GetCount gets total account count
func (s *Service) GetCount(tenantID id.ID) int {
	count, err := s.repo.GetCount(tenantID)
	if err != nil {
		return 0
	}

	return count
}

// SearchAccounts search accounts
func (s *Service) SearchAccounts(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error) {
	accounts, err := s.repo.Search(tenantID, query, page, limit, at)
	if err != nil {
		return nil, err
	}
	if len(accounts) == 0 {
		return nil, glad.ErrNotFound
	}
	return accounts, nil
}
