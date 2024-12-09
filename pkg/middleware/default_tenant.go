/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package middleware

import (
	"ac9/glad/config"
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/util"
	"net/http"
)

// AddDefaultTenant Adds default tenant identifier
func AddDefaultTenant(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get(common.HttpHeaderTenantID) == "" {
		r.Header.Set(common.HttpHeaderTenantID, util.GetStrEnvOrConfig("DEFAULT_TENANT", config.DEFAULT_TENANT))
	}
	next(w, r)
}

// TODO: Until authentication piece is in place
// AddDefaultAccount Adds default account identifier
func AddDefaultAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get(common.HttpHeaderAccountID) == "" {
		r.Header.Set(common.HttpHeaderAccountID, util.GetStrEnvOrConfig("DEFAULT_ACCOUNT", config.DEFAULT_ACCOUNT))
	}
	next(w, r)
}
