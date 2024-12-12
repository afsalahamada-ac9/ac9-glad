/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package common

import (
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"net/http"
	"strconv"
)

const (
	HttpHeaderTenantID      = "X-GLAD-TenantID"
	HttpHeaderAccountID     = "X-GLAD-AccountID"
	HttpHeaderAccountEmail  = "X-GLAD-AccountEmail"
	HttpHeaderTotalCount    = "X-Total-Count"
	HttpHeaderAuthorization = "Authorization"
	HttpHeaderBearer        = "Bearer "

	DefaultHttpPageNumber = 1
	DefaultHttpPageLimit  = 20

	HttpParamLimit = "limit"
	HttpParamPage  = "page"
	HttpParamQuery = "q"
	HttpParamType  = "type"

	MaxHttpPaginationLimit = 50
)

// HttpGetPageParams retrieves the page params such as page and limit from the path
func HttpGetPageParams(
	w http.ResponseWriter,
	r *http.Request,
) (
	page int,
	limit int,
	err error,
) {
	page, _ = strconv.Atoi(r.URL.Query().Get(HttpParamPage))
	limit, _ = strconv.Atoi(r.URL.Query().Get(HttpParamLimit))

	// Guard rails to limit the items queried from DB and sent
	if page == 0 && limit == 0 {
		page = DefaultHttpPageNumber
		limit = min(DefaultHttpPageLimit, MaxHttpPaginationLimit)
	}

	if limit > MaxHttpPaginationLimit {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Page size requested is more than allowed limit"))
		err = glad.ErrInvalidValue
		return
	}

	return
}

// HttpGetTenantID retrieves tenant id from the header
func HttpGetTenantID(
	w http.ResponseWriter, r *http.Request,
) (tenantID id.ID, err error) {
	tenant := r.Header.Get(HttpHeaderTenantID)
	tenantID, err = id.FromString(tenant)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing tenant ID"))
		return
	}
	return
}

// HttpGetAccountID retrieves account id from the header
func HttpGetAccountID(
	w http.ResponseWriter, r *http.Request,
) (accountID id.ID, err error) {
	account := r.Header.Get(HttpHeaderAccountID)
	accountID, err = id.FromString(account)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Missing account ID"))
		return
	}
	return
}
