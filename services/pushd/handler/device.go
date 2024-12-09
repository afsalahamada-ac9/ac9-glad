/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"ac9/glad/pkg/common"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/pushd/presenter"
	"ac9/glad/services/pushd/usecase/device"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func register(service device.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req presenter.DeviceRegisterRequest
		errorMessage := "Unable to register device"

		tenantID, err := common.HttpGetTenantID(w, r)
		if err != nil {
			l.Log.Debugf("Tenant id is missing")
			return
		}

		accountID, err := common.HttpGetAccountID(w, r)
		if err != nil {
			l.Log.Debugf("Account id is missing")
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		device, err := req.ToDevice(tenantID, accountID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to copy to device entity"))
			return
		}

		deviceID, err := service.Create(device)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		jsonResponse := &presenter.DeviceRegisterResponse{
			ID: deviceID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenantID.String())
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
			l.Log.Errorf("%v, err=%v", jsonResponse, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeDeviceHandlers make push handlers
func MakeDeviceHandlers(r *mux.Router, n negroni.Negroni, service device.UseCase) {
	r.Handle("/v1/push-notify/register", n.With(
		negroni.Wrap(register(service)),
	)).Methods("POST").Name("register")
}
