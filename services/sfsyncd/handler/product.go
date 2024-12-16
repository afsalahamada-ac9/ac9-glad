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
	"ac9/glad/services/sfsyncd/presenter"
	"ac9/glad/services/sfsyncd/usecase/sf_import"
	"encoding/json"
	"log"
	"net/http"

	l "ac9/glad/pkg/logger"

	"github.com/gorilla/mux"
	"github.com/ulule/deepcopier"
	"github.com/urfave/negroni"
)

func importProducts(service sf_import.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error importing products"

		var sfProducts []presenter.Product
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&sfProducts)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		var gProducts []*glad.Product
		for _, sfProduct := range sfProducts {
			product := &glad.Product{}
			deepcopier.Copy(sfProduct).To(product)
			gProducts = append(gProducts, product)
		}

		gResponses, err := service.ImportProduct(tenantID, gProducts)
		if err != nil {
			l.Log.Warnf("Unable to import products tenantID=%v, err=%v", tenantID, err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to import products. " + err.Error()))
		}

		var sfResponses []*presenter.ProductResponse
		for _, gResponse := range gResponses {
			resp := &presenter.ProductResponse{}
			deepcopier.Copy(gResponse).To(resp)
			sfResponses = append(sfResponses, resp)
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(sfResponses); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func importCenters(service sf_import.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error importing centers"

		var sfCenters []presenter.CenterWrapper
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&sfCenters)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		var gCenters []*glad.Center
		for _, sfCenter := range sfCenters {
			center := &glad.Center{}
			sfCenter.Value.ToGladCenter(center)
			l.Log.Warnf("sfCenter=%#v, center=%#v", sfCenter.Value, center)
			gCenters = append(gCenters, center)
		}

		gResponses, err := service.ImportCenter(tenantID, gCenters)
		if err != nil {
			l.Log.Warnf("Unable to import centers tenantID=%v, err=%v", tenantID, err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to import centers. " + err.Error()))
		}

		var sfResponses []*presenter.CenterResponse
		for _, gResponse := range gResponses {
			resp := &presenter.CenterResponse{}
			deepcopier.Copy(gResponse).To(resp)
			sfResponses = append(sfResponses, resp)
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(sfResponses); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeImportHandlers make import handlers
func MakeProductHandlers(r *mux.Router, n negroni.Negroni, service sf_import.UseCase) {
	r.Handle("/v1/import/salesforce/products", n.With(
		negroni.Wrap(importProducts(service)),
	)).Methods(http.MethodPost, http.MethodOptions).Name("importProducts")

	r.Handle("/v1/import/salesforce/product", n.With(
		negroni.Wrap(importProducts(service)),
	)).Methods(http.MethodPost, http.MethodOptions).Name("importProducts")

	r.Handle("/v1/import/salesforce/centers", n.With(
		negroni.Wrap(importCenters(service)),
	)).Methods(http.MethodPost, http.MethodOptions).Name("importCenters")

}
