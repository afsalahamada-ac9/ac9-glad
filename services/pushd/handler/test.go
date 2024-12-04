/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"ac9/glad/entity"
	"ac9/glad/pkg/common"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/pushd/presenter"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func registerPushNotify() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input presenter.PushNotifyInfo

		tenant := r.Header.Get(common.HttpHeaderTenantID)
		_, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		l.Log.Debugf("Push notify register received")

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
	})
}

// MakeTestHandlers make push handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/push-notify/register", n.With(
		negroni.Wrap(registerPushNotify()),
	)).Methods("POST").Name("registerPushNotify")
}
