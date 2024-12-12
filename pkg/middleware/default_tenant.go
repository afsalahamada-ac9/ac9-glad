/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package middleware

import (
	"ac9/glad/config"
	"ac9/glad/pkg/common"
	l "ac9/glad/pkg/logger"
	"ac9/glad/pkg/util"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// AddDefaultTenant Adds default tenant identifier
func AddDefaultTenant(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get(common.HttpHeaderTenantID) == "" {
		r.Header.Set(common.HttpHeaderTenantID, util.GetStrEnvOrConfig("DEFAULT_TENANT", config.DEFAULT_TENANT))
	}
	next(w, r)
}

// TODO: Remove the temporary workaround that is in place until authentication is complete
// AddDefaultAccount Adds default account identifier
func AddDefaultAccount(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.Header.Get(common.HttpHeaderAccountID) == "" {
		r.Header.Set(common.HttpHeaderAccountID, util.GetStrEnvOrConfig("DEFAULT_ACCOUNT", config.DEFAULT_ACCOUNT))

		authHeader := r.Header.Get(common.HttpHeaderAuthorization)
		if authHeader == "" || !strings.HasPrefix(authHeader, common.HttpHeaderBearer) {
			l.Log.Warnf("Missing or invalid token, skipping for now. header=%v", authHeader)
			next(w, r)
		}

		bearerToken := strings.TrimPrefix(authHeader, common.HttpHeaderBearer)
		token, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(bearerToken, jwt.MapClaims{})
		if err != nil {
			l.Log.Warnf("Failed to parse the token. err=%v", err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			l.Log.Warnf("Unable to get claims=%v", claims)
			return
		}

		email := claims["email"].(string)
		if email == "" {
			l.Log.Warnf("Unable to get email from claims=%v", claims)
			return
		}
		l.Log.Infof("Account=%v", email)
		r.Header.Set(common.HttpHeaderAccountEmail, email)
	}
	next(w, r)
}
