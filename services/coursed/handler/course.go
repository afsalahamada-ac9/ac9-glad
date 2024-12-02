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

	"ac9/glad/pkg/common"
	"ac9/glad/usecase/course"

	"ac9/glad/services/coursed/presenter"

	"ac9/glad/entity"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
)

func listCourses(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading courses"
		var data []*entity.Course
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
			data, err = service.ListCourses(tenantID, page, limit)
		default:
			// TODO: search need to be reworked; need to add a count
			// for search; also need to see how the caller generates
			// the search query request
			data, err = service.SearchCourses(tenantID, search, page, limit)
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
		var toJ []*presenter.Course
		for _, d := range data {
			pc := &presenter.Course{
				ID:           d.ID,
				Name:         &d.Name,
				Mode:         &d.Mode,
				CenterID:     &d.CenterID,
				Notes:        &d.Notes,
				Timezone:     &d.Timezone,
				Status:       &d.Status,
				MaxAttendees: &d.MaxAttendees,
				NumAttendees: &d.NumAttendees,
			}
			pc.Address = &presenter.Address{}
			pc.Address.CopyFrom(d.Address)

			toJ = append(toJ, pc)
		}
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.Header().Set(common.HttpHeaderTenantID, tenant)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode course"))
		}
	})
}

func createCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error adding course"
		var input presenter.CourseReq

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

		course, err := input.ToCourse(tenantID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to copy to course entity"))
			return
		}

		cos, _ := input.ToCourseOrganizer()
		cts, _ := input.ToCourseTeacher()
		ccs, _ := input.ToCourseContact()
		cns, _ := input.ToCourseNotify()

		// TODO: validation checks to be performed

		courseTimings, err := input.ToCourseTiming()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to copy to course timing entity"))
			return
		}

		courseID, courseTimingsID, err := service.CreateCourse(
			course,
			cos,
			cts,
			ccs,
			cns,
			courseTimings,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}
		toJ := &presenter.CourseResponse{
			ID:         courseID,
			DateTimeID: courseTimingsID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func getCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading course"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		data, err := service.GetCourse(id)
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

		toJ := &presenter.Course{
			ID:           data.ID,
			Name:         &data.Name,
			Mode:         &data.Mode,
			CenterID:     &data.CenterID,
			Notes:        &data.Notes,
			Timezone:     &data.Timezone,
			Status:       &data.Status,
			MaxAttendees: &data.MaxAttendees,
			NumAttendees: &data.NumAttendees,
		}

		toJ.Address = &presenter.Address{}
		toJ.Address.CopyFrom(data.Address)

		w.Header().Set(common.HttpHeaderTenantID, data.TenantID.String())
		if err := json.NewEncoder(w).Encode(toJ); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode course"))
		}
	})
}

func deleteCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing course"
		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
		err = service.DeleteCourse(id)
		switch err {
		case nil:
			w.WriteHeader(http.StatusOK)
			return
		case entity.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Course doesn't exist"))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage))
			return
		}
	})
}

func updateCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error updating course"

		vars := mux.Vars(r)
		id, err := entity.StringToID(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input presenter.CourseReq
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

		course, err := input.ToCourse(tenantID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to copy to course entity"))
			return
		}
		// ID sent in the body will be ignored for update. Only the ID sent in path
		// will be taken into consideration
		// Bugs in client could cause more damage. We can add a check to see whether the
		// ID sent in the body and the path are same.
		course.ID = id

		cos, _ := input.ToCourseOrganizer()
		cts, _ := input.ToCourseTeacher()
		ccs, _ := input.ToCourseContact()
		cns, _ := input.ToCourseNotify()

		// TODO: validation checks to be performed

		// TODO: Course timings should contain ID
		// Note: Once course is created, then additional day cannot be added via API
		courseTimings, err := input.ToCourseTiming()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to copy to course timing entity"))
			return
		}

		err = service.UpdateCourse(
			course,
			cos,
			cts,
			ccs,
			cns,
			courseTimings,
		)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		toJ := &presenter.Course{
			ID: course.ID,
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

// MakeCourseHandlers make url handlers
func MakeCourseHandlers(r *mux.Router, n negroni.Negroni, service course.UseCase) {
	r.Handle("/v1/courses", n.With(
		negroni.Wrap(listCourses(service)),
	)).Methods("GET", "OPTIONS").Name("listCourses")

	r.Handle("/v1/courses", n.With(
		negroni.Wrap(createCourse(service)),
	)).Methods("POST", "OPTIONS").Name("createCourse")

	r.Handle("/v1/courses/{id}", n.With(
		negroni.Wrap(getCourse(service)),
	)).Methods("GET", "OPTIONS").Name("getCourse")

	r.Handle("/v1/courses/{id}", n.With(
		negroni.Wrap(deleteCourse(service)),
	)).Methods("DELETE", "OPTIONS").Name("deleteCourse")

	r.Handle("/v1/courses/{id}", n.With(
		negroni.Wrap(updateCourse(service)),
	)).Methods("PUT", "OPTIONS").Name("updateCourse")
}
