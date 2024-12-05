/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"net/http"

	"ac9/glad/pkg/common"
	"ac9/glad/services/coursed/presenter"
	"ac9/glad/usecase/account"

	"ac9/glad/entity"

	"github.com/gorilla/mux"
	"github.com/ulule/deepcopier"
	"github.com/urfave/negroni"
)

// TODO: This is temporary until we implement the correct search
func getParticipantByCourse(service account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading participants"
		tenant := r.Header.Get(common.HttpHeaderTenantID)

		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		// hard-coded value
		data, err := service.GetAccountByName(tenantID, "Sri Sri Ravi Shankar")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Empty data returned"))
			return
		}

		account := &presenter.Account{}
		deepcopier.Copy(data).To(account)

		p := &presenter.Participant{
			ID:      account.ID + 3,
			Email:   account.Email,
			Account: account,
		}

		toJ := []*presenter.Participant{p}

		w.Header().Set(common.HttpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode course"))
		}
	})
}

// MakeCourseHandlers make url handlers
func MakeParticipantHandlers(r *mux.Router, n negroni.Negroni, service account.UseCase) {
	// TODO: implement get participants by course-id
	r.Handle("/v1/participants/{courseId}", n.With(
		negroni.Wrap(getParticipantByCourse(service)),
	)).Methods("GET", "OPTIONS").Name("getParticipantByCourse")
}
