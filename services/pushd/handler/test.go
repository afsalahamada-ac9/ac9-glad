/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"ac9/glad/services/pushd/presenter"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func registerPushNotify() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// errorMessage := "Error retrieving metadata"
		// // Prepare a client
		// c := &fasthttp.Client{}

		// fullPath := "http://" +
		// 	util.GetStrEnvOrConfig("COURSED_ADDR", config.COURSED_ADDR) +
		// 	"/v1/products"
		// l.Log.Infof("fullPath=%v", fullPath)
		// statusCode, body, err := c.Get(nil, fullPath)

		// if err != nil {
		// 	l.Log.Errorf("%v. err=%v", errorMessage, err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	_, _ = w.Write([]byte("Get() coursed products error"))
		// 	return
		// }

		// if statusCode != fasthttp.StatusOK {
		// 	w.WriteHeader(statusCode)
		// 	return
		// }
		pushNotification := presenter.PushNotification{
			PushToken:   "asdbf8ay2",
			RevokeID:    "je182e2",
			AppVersion:  "2024.12.1",
			DeviceInfo:  map[string]interface{}{},
			PlatformInfo: map[string]interface{}{},
		}
		

		if err := json.NewEncoder(w).Encode(pushNotification); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode metadata"))
		}

		w.WriteHeader(http.StatusOK)
	})
}

// MakeTestHandlers make url handlers
func MakeTestHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/push-notify/register", n.With(
		negroni.Wrap(registerPushNotify()),
	)).Methods("GET").Name("register-push-notification")
}
