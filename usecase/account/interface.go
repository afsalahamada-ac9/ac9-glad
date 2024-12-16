/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package account

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/id"
)

// Reader interface
type Reader interface {
	GetByName(tenantID id.ID, username string) (*entity.Account, error)
	Get(accountID id.ID) (*entity.Account, error)
	List(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	Search(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	GetCount(tenantID id.ID) (int, error)
	GetByEmail(tenantID id.ID, email string) (*entity.Account, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Account) error
	Update(e *entity.Account) error
	Delete(accountID id.ID) error
	DeleteByName(tenantID id.ID, username string) error
	Upsert(e *entity.Account) (id.ID, error)
}

// Repository interface
type Repository interface {
	Reader
	Writer
}

// UseCase interface
type UseCase interface {
	CreateAccount(
		tenantID id.ID,
		cognitoID string,
		username string,
		first_name string,
		last_name string,
		phone string,
		email string,
		at entity.AccountType,
		as entity.AccountStatus,
	) error
	GetAccount(tenantID id.ID, accountID id.ID) (*entity.Account, error)
	GetAccountByName(tenantID id.ID, username string) (*entity.Account, error)
	ListAccounts(tenantID id.ID, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	UpdateAccount(e *entity.Account) error
	DeleteAccount(tenantID id.ID, accountID id.ID) error
	DeleteAccountByName(tenantID id.ID, username string) error
	GetCount(tenantId id.ID) int
	SearchAccounts(tenantID id.ID, query string, page, limit int, at entity.AccountType) ([]*entity.Account, error)
	GetAccountByEmail(tenantID id.ID, email string) (*entity.Account, error)
	UpsertAccount(e *entity.Account) (id.ID, error)
}
