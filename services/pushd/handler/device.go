/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
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
		switch err {
		case nil:
			break
		case glad.ErrAlreadyExists:
			w.WriteHeader(http.StatusConflict)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		default:
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

func getByAccount(service device.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading device"
		vars := mux.Vars(r)
		accountID, err := id.FromString(vars["accountID"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		tenantID, err := common.HttpGetTenantID(w, r)
		if err != nil {
			l.Log.Debugf("Tenant id is missing")
			return
		}

		l.Log.Debugf("Tenant id=%v, Account id=%v", tenantID, accountID)
		data, err := service.GetByAccount(tenantID, accountID)
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Empty data returned"))
			return
		}

		var jsonDevices []*presenter.Device
		for _, d := range data {
			pc := &presenter.Device{
				ID:         d.ID,
				AccountID:  d.AccountID,
				TenantID:   d.TenantID,
				PushToken:  d.PushToken,
				RevokeID:   d.RevokeID,
				AppVersion: d.AppVersion,
			}

			jsonDevices = append(jsonDevices, pc)
		}

		w.Header().Set(common.HttpHeaderTenantID, tenantID.String())
		if err := json.NewEncoder(w).Encode(jsonDevices); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode devices"))
		}
	})
}

func notify(service device.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req presenter.Notify
		errorMessage := "Unable to notify device"

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			l.Log.Warnf("Unable to decode the request. err=%v", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the request. " + err.Error()))
			return
		}

		var jsonResponse []*presenter.NotifyStatus
		statuses, err := service.Notify(req.TenantID,
			req.AccountID,
			req.NotificationMessage.Header,
			req.NotificationMessage.Content,
		)

		for idx, status := range statuses {
			notifyStatus := &presenter.NotifyStatus{AccountID: req.AccountID[idx], Status: status}
			jsonResponse = append(jsonResponse, notifyStatus)
		}

		w.Header().Set(common.HttpHeaderTenantID, r.Header.Get(common.HttpHeaderTenantID))
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
			l.Log.Warnf("%v, err=%v", jsonResponse, err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func delete(service device.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Unable to delete device"
		vars := mux.Vars(r)
		deviceID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		err = service.Delete(deviceID)
		w.Header().Set(common.HttpHeaderTenantID, r.Header.Get(common.HttpHeaderTenantID))
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case glad.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Device doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeDeviceHandlers make push handlers
func MakeDeviceHandlers(r *mux.Router, n negroni.Negroni, service device.UseCase) {
	// Note: We should deprecate this in the upcoming release
	r.Handle("/v1/push-notify/register", n.With(
		negroni.Wrap(register(service)),
	)).Methods("POST", "OPTIONS").Name("register")

	r.Handle("/v1/device/register", n.With(
		negroni.Wrap(register(service)),
	)).Methods("POST", "OPTIONS").Name("register")

	r.Handle("/v1/device/notify", n.With(
		negroni.Wrap(notify(service)),
	)).Methods("POST", "OPTIONS").Name("notify")

	r.Handle("/v1/device/account/{accountID}", n.With(
		negroni.Wrap(getByAccount(service)),
	)).Methods("GET", "OPTIONS").Name("getByAccount")

	r.Handle("/v1/device/{id}", n.With(
		negroni.Wrap(delete(service)),
	)).Methods("DELETE", "OPTIONS").Name("delete")
}
