/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package middleware

import (
	"ac9/glad/pkg/logger"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/urfave/negroni"
)

type APILogging struct {
	Log logger.Logger
}

func (l *APILogging) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Log request body
	bodyBytes, _ := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Reassign the body to preserve it
	compact := &bytes.Buffer{}
	json.Compact(compact, []byte(bodyBytes))
	l.Log.Debugf("I %v", compact.String())

	// Capture response body using a ResponseWriter wrapper
	ww := negroni.NewResponseWriter(rw)
	rec := httptest.NewRecorder()

	next(rec, r)

	// Log response body
	l.Log.Debugf("O %v", rec.Body.String())
	for k, v := range rec.Header() {
		ww.Header()[k] = v
	}
	ww.WriteHeader(rec.Code)
	rec.Body.WriteTo(ww)
}
