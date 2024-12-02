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

	"sudhagar/glad/coursed/presenter"
	"sudhagar/glad/entity"
	"sudhagar/glad/pkg/common"
	"sudhagar/glad/usecase/product"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func listProducts(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading products"
		var data []*entity.Product
		var err error
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		search := r.URL.Query().Get(httpParamQuery)
		page, _ := strconv.Atoi(r.URL.Query().Get(httpParamPage))
		limit, _ := strconv.Atoi(r.URL.Query().Get(httpParamLimit))

		tenantID, err := entity.StringToID(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to parse tenant id"))
			return
		}

		switch {
		case search == "":
			data, err = service.ListProducts(tenantID, page, limit)
		default:
			data, err = service.SearchProducts(tenantID, search, page, limit)
		}
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		total := service.GetCount(tenantID)
		w.Header().Set(httpHeaderTotalCount, strconv.Itoa(total))

		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte(errorMessage))
			return
		}

		var toJ []*presenter.Product
		for _, d := range data {
			toJ = append(toJ, &presenter.Product{
				ID:               d.ID,
				ExtName:          d.ExtName,
				Title:            d.Title,
				CType:            d.CType,
				BaseProductExtID: d.BaseProductExtID,
				DurationDays:     d.DurationDays,
				Visibility:       d.Visibility,
				MaxAttendees:     d.MaxAttendees,
				Format:           d.Format,
				IsAutoApprove:    d.IsAutoApprove,
			})
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode product"))
		}
	})
}

func createProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding product"
		var input struct {
			ExtID            string                   `json:"extId"`
			ExtName          string                   `json:"extName"`
			Title            string                   `json:"title"`
			CType            string                   `json:"ctype"`
			BaseProductExtID string                   `json:"baseProductExtId"`
			DurationDays     int32                    `json:"durationDays"`
			Visibility       entity.ProductVisibility `json:"visibility"`
			MaxAttendees     int32                    `json:"maxAttendees"`
			Format           entity.ProductFormat     `json:"format"`
			IsAutoApprove    bool                     `json:"isAutoApprove"`
		}

		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
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

		id, err := service.CreateProduct(
			tenantID,
			input.ExtID,
			input.ExtName,
			input.Title,
			input.CType,
			input.BaseProductExtID,
			input.DurationDays,
			input.Visibility,
			input.MaxAttendees,
			input.Format,
			input.IsAutoApprove,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Product{
			ID:               id,
			ExtName:          input.ExtName,
			Title:            input.Title,
			CType:            input.CType,
			BaseProductExtID: input.BaseProductExtID,
			DurationDays:     input.DurationDays,
			Visibility:       input.Visibility,
			MaxAttendees:     input.MaxAttendees,
			Format:           input.Format,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
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
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetProduct(id)
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

		toJ := &presenter.Product{
			ID:               data.ID,
			ExtName:          data.ExtName,
			Title:            data.Title,
			CType:            data.CType,
			BaseProductExtID: data.BaseProductExtID,
			DurationDays:     data.DurationDays,
			Visibility:       data.Visibility,
			MaxAttendees:     data.MaxAttendees,
			Format:           data.Format,
			IsAutoApprove:    data.IsAutoApprove,
		}

		w.Header().Set(common.HttpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode product"))
		}
	})
}

func deleteProduct(service product.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing product"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
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
		case entity.ErrNotFound:
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
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input entity.Product
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := entity.StringToID(tenant)
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

		input.ID = id
		input.TenantID = tenantID
		err = service.UpdateProduct(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Product{
			ID:               input.ID,
			ExtName:          input.ExtName,
			Title:            input.Title,
			CType:            input.CType,
			BaseProductExtID: input.BaseProductExtID,
			DurationDays:     input.DurationDays,
			Visibility:       input.Visibility,
			MaxAttendees:     input.MaxAttendees,
			Format:           input.Format,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
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
