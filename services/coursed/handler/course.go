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
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	l "ac9/glad/pkg/logger"
	"ac9/glad/usecase/account"
	"ac9/glad/usecase/center"
	"ac9/glad/usecase/course"
	"ac9/glad/usecase/product"

	"ac9/glad/services/coursed/presenter"

	"ac9/glad/entity"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func listCourses(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error reading courses"
		var data []*entity.Course
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
			data, err = service.ListCourses(tenantID, page, limit)
		default:
			// TODO: search need to be reworked; need to add a count
			// for search; also need to see how the caller generates
			// the search query request
			data, err = service.SearchCourses(tenantID, search, page, limit)
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
		var courses []*presenter.Course
		for _, d := range data {
			pc := &presenter.Course{}
			pc.FromEntityCourse(d)

			pc.Address = &presenter.Address{}
			pc.Address.CopyFrom(d.Address)

			courses = append(courses, pc)
		}
		if err := json.NewEncoder(w).Encode(courses); err != nil {
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
		response := &presenter.CourseResponse{
			ID:         courseID,
			DateTimeID: courseTimingsID,
		}

		w.Header().Set(common.HttpHeaderTenantID, tenant)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			l.Log.Errorf(err.Error())
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
		id, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		courseFull, err := service.GetCourse(id)
		if err != nil && err != glad.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
			return
		}

		if courseFull == nil {
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Empty data returned"))
			return
		}

		response := &presenter.Course{}
		response.FromEntityCourseFull(courseFull)

		w.Header().Set(common.HttpHeaderTenantID, courseFull.Course.TenantID.String())
		if err := json.NewEncoder(w).Encode(response); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Unable to encode course"))
		}
	})
}

func getCourseByAccount(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		accountID, err := id.FromString(vars["accountID"])
		if err != nil {
			l.Log.Warnf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		getCourseByAccountImpl(w, r, accountID, service)
	})
}

func getCourseMe(service course.UseCase, accountService account.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID, err := common.HttpGetTenantID(w, r)
		if err != nil {
			l.Log.Debugf("Tenant id is missing")
			return
		}

		email := r.Header.Get(common.HttpHeaderAccountEmail)
		if email == "" {
			l.Log.Warnf("Email id is missing")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to map to an account"))
			return
		}

		account, err := accountService.GetAccountByEmail(tenantID, email)
		switch err {
		case nil:
			break
		case glad.ErrNotFound:
			l.Log.Warnf("Account (%v) mapping doesn't exist", email)
			w.WriteHeader(http.StatusNotFound)
			_, _ = w.Write([]byte("Account mapping doesn't exist"))
			return
		default:
			l.Log.Warnf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		getCourseByAccountImpl(w, r, account.ID, service)
	})
}

func getCourseByAccountImpl(w http.ResponseWriter,
	r *http.Request,
	accountID id.ID,
	service course.UseCase) {
	errorMessage := "Error reading course"

	tenantID, err := common.HttpGetTenantID(w, r)
	if err != nil {
		l.Log.Warnf("Tenant id is missing")
		return
	}

	page, limit, err := common.HttpGetPageParams(w, r)
	if err != nil {
		l.Log.Warnf("%v", err)
		return
	}

	count, cfList, err := service.GetCourseByAccount(tenantID, accountID, page, limit)
	if err != nil && err != glad.ErrNotFound {
		l.Log.Warnf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(errorMessage + ":" + err.Error()))
		return
	}

	if count == 0 || len(cfList) == 0 {
		l.Log.Warnf("count=%v, len(cfList)=%v, err=%v", count, len(cfList), err)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Empty data returned"))
		return
	}

	var courseList []*presenter.Course
	for _, courseFull := range cfList {
		c := &presenter.Course{}
		c.FromEntityCourseFull(courseFull)
		courseList = append(courseList, c)
	}

	w.Header().Set(common.HttpHeaderTotalCount, strconv.Itoa(count))
	w.Header().Set(common.HttpHeaderTenantID, tenantID.String())
	if err := json.NewEncoder(w).Encode(courseList); err != nil {
		l.Log.Warnf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Unable to encode course"))
	}
}

func deleteCourse(service course.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error removing course"
		vars := mux.Vars(r)
		id, err := id.FromString(vars["id"])
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
		case glad.ErrNotFound:
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
		courseID, err := id.FromString(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		var input presenter.CourseReq
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
		course.ID = courseID

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

		response := &presenter.Course{
			CourseResponse: presenter.CourseResponse{
				ID: course.ID,
			},
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

func importCourse(service course.UseCase,
	accountService account.UseCase,
	productService product.UseCase,
	centerService center.UseCase,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error importing courses"

		var gCourses []glad.Course
		tenant := r.Header.Get(common.HttpHeaderTenantID)
		tenantID, err := id.FromString(tenant)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Missing tenant ID"))
			return
		}

		err = json.NewDecoder(r.Body).Decode(&gCourses)
		if err != nil {
			l.Log.Warnf("Unable to decode object. err = %v", err)
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Unable to decode the data. " + err.Error()))
			return
		}

		var response []*presenter.ImportCourseResponse
		for _, gCourse := range gCourses {
			course := &entity.Course{}
			presenter.GladCourseToEntity(gCourse, course)
			course.TenantID = tenantID

			productID, err :=
				productService.GetIDByExtID(tenantID, gCourse.ProductExtID)
			if err != nil {
				l.Log.Warnf("Unable to get product id extID=%v, err=%v", gCourse.ProductExtID, err)
				response = append(response, &presenter.ImportCourseResponse{
					ExtID:   gCourse.ExtID,
					IsError: err != nil,
				})
				continue
			}

			centerID, err :=
				centerService.GetIDByExtID(tenantID, gCourse.CenterExtID)
			if err != nil {
				l.Log.Warnf("Unable to get center id extID=%v, err=%v", gCourse.CenterExtID, err)
				response = append(response, &presenter.ImportCourseResponse{
					ExtID:   gCourse.ExtID,
					IsError: err != nil,
				})
				continue
			}

			course.ProductID = productID
			course.CenterID = centerID

			// TODO: optimize DB operations by doing multiple inserts simultaneously
			courseID, err := service.UpsertCourse(course)
			if err != nil {
				l.Log.Warnf("Unable to upsert course extID=%v, err=%v", course.ExtID, err)
			}

			response = append(response, &presenter.ImportCourseResponse{
				ID:      courseID,
				ExtID:   *course.ExtID,
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

// MakeCourseHandlers make url handlers
func MakeCourseHandlers(r *mux.Router,
	n negroni.Negroni,
	service course.UseCase,
	accountService account.UseCase,
	productService product.UseCase,
	centerService center.UseCase,
) {
	r.Handle("/v1/courses", n.With(
		negroni.Wrap(listCourses(service)),
	)).Methods("GET", "OPTIONS").Name("listCourses")

	r.Handle("/v1/courses", n.With(
		negroni.Wrap(createCourse(service)),
	)).Methods("POST", "OPTIONS").Name("createCourse")

	r.Handle("/v1/courses/import", n.With(
		negroni.Wrap(importCourse(service, accountService, productService, centerService)),
	)).Methods("POST", "OPTIONS").Name("importCourse")

	// get courses by account-id
	r.Handle("/v1/courses/account/{accountID}", n.With(
		negroni.Wrap(getCourseByAccount(service)),
	)).Methods("GET", "OPTIONS").Name("getCourseByAccount")

	r.Handle("/v1/courses/me", n.With(
		negroni.Wrap(getCourseMe(service, accountService)),
	)).Methods("GET", "OPTIONS").Name("getCourseMe")

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
