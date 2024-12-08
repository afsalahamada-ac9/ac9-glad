/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package common

import (
	"ac9/glad/entity"
	"net/http"
	"strconv"
)

const (
	HttpHeaderTenantID   = "X-GLAD-TenantID"
	HttpHeaderAccountID  = "X-GLAD-AccountID"
	HttpHeaderTotalCount = "X-Total-Count"

	DBFormatDateTimeMS   = "2006-01-02 15:04:05.000"
	DBFormatDate         = "2006-01-02"
	DBFormatTimeHH_MM    = "15:04:00"
	DBFormatTimeHH_MM_SS = "15:04:05"

	DefaultHttpPageNumber = 1
	DefaultHttpPageLimit  = 20

	HttpParamLimit = "limit"
	HttpParamPage  = "page"
	HttpParamQuery = "q"
	HttpParamType  = "type"

	MaxHttpPaginationLimit = 50
)

func HttpGetPathParams(
	w http.ResponseWriter,
	r *http.Request,
) (
	search string,
	page int,
	limit int,
	err error,
) {
	search = r.URL.Query().Get(HttpParamQuery)
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
		err = entity.ErrInvalidEntity // TODO
		return
	}

	return
}
