/*
 * Copyright 2024 AboveCloud9.AI Products and Services Private Limited
 * All rights reserved.
 * This code may not be used, copied, modified, or distributed without explicit permission.
 */

package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"ac9/glad/entity"
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/services/coursed/presenter"
	"ac9/glad/usecase/product"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listProducts(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading products"
		var data []*entity.Product
		var err error
		tenant := r.Header.Get(common.HttpHeaderTenantID)

		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		search := r.URL.Query().Get(common.HttpParamQuery)
		page, limit, err := common.HttpGetPageParams(w, r)
		if err != nil {
			return
		}

		switch {
		case search == "":
			data, err = service.ListProducts(tenantID, page, limit)
		default:
			data, err = service.SearchProducts(tenantID, search, page, limit)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(common.HttpHeaderTotalCount, strconv.Itoa(total))

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		var products []*presenter.Product
		for _, d := range data {
			product := &presenter.Product{}
			product.FromEntityProduct(d)

			products = append(products, product)
		}
		if err := json.NewEncoder(w).Encode(products); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode product"))
		}
	})
}

func createProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding product"
		var req presenter.ProductReq

		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		productID, err := service.CreateProduct(
			tenantID,
			req.ExtName,
			req.Title,
			req.CType,
			req.BaseProductExtID,
			req.DurationDays,
			req.Visibility,
			req.MaxAttendees,
			req.Format,
			req.IsAutoApprove,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		response := &presenter.ProductResponse{
			ID: productID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func getProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading product"
		vars := mux.Vars(r)
		id, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetProduct(id)
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

		response := &presenter.Product{}
		response.FromEntityProduct(data)

		w.Header().Set(common.HttpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode product"))
		}
	})
}

func deleteProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing product"
		vars := mux.Vars(r)
		id, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteProduct(id)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case glad.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Product doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating product"

		vars := mux.Vars(r)
		productID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input entity.Product
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
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

		input.ID = productID
		input.TenantID = tenantID
		err = service.UpdateProduct(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
	})
}

func importProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error importing products"

		var iProducts []glad.Product
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&iProducts)
		if err != nil {
			l.Log.Warnf("Unable to decode object. err=%v", err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		var response []*presenter.ProductImportResponse
		for _, input := range iProducts {
			product := &entity.Product{}
			presenter.GladProductToEntity(input, product)
			product.TenantID = tenantID

			// TODO: optimize DB operations by doing multiple inserts simultaneously
			productID, err := service.UpsertProduct(product)
			if err != nil {
				l.Log.Warnf("Unable to upsert product extID=%v, err=%v", product.ExtID, err)
			}

			response = append(response, &presenter.ProductImportResponse{
				ID:      productID,
				ExtID:   product.ExtID,
				IsError: err != nil,
			})
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

// MakeProductHandlers make url handlers
func MakeProductHandlers(r *mux.Router, n negroni.Negroni, service product.UseCase) {
	r.Handle("/v1/products", n.With(
		negroni.Wrap(listProducts(service)),
	)).Methods("GET", "OPTIONS").Name("listProducts")

	r.Handle("/v1/products", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS").Name("createProduct")

	r.Handle("/v1/products/import", n.With(
		negroni.Wrap(importProduct(service)),
	)).Methods("POST", "OPTIONS").Name("importProduct")

	r.Handle("/v1/products/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS").Name("getProduct")

	r.Handle("/v1/products/{id}", n.With(
		negroni.Wrap(deleteProduct(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteProduct")

	r.Handle("/v1/products/{id}", n.With(
		negroni.Wrap(updateProduct(service)),
	)).Methods("PUT", "OPTIONS").Name("updateProduct")
}
