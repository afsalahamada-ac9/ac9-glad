/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package account

import (
	"sudhagar/glad/entity"
)

// Reader interface
type Reader interface {
	GetByName(tenantID entity.ID, username string) (*entity.Account, error)
	Get(id entity.ID) (*entity.Account, error)
	List(tenantID entity.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	Search(tenantID entity.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	GetCount(tenantId entity.ID) (int, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Account) error
	Update(e *entity.Account) error
	Delete(id entity.ID) error
	DeleteByName(tenantID entity.ID, username string) error
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreateAccount(
		tenantID entity.ID,
		extID string,
		cognitoID string,
		username string,
		first_name string,
		last_name string,
		phone string,
		email string,
		at entity.AccountType,
		as entity.AccountStatus,
	) error
	GetAccount(id entity.ID) (*entity.Account, error)
	GetAccountByName(tenantID entity.ID, username string) (*entity.Account, error)
	ListAccounts(tenantID entity.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	UpdateAccount(e *entity.Account) error
	DeleteAccount(id entity.ID) error
	DeleteAccountByName(tenantID entity.ID, username string) error
	GetCount(tenantId entity.ID) int
	SearchAccounts(tenantID entity.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error)
}
