package sf_import

import (
	"ac9/glad/pkg/common"
	"ac9/glad/pkg/glad"
	"ac9/glad/pkg/id"
	"encoding/json"
	"net/http"

	l "ac9/glad/pkg/logger"

	"github.com/valyala/fasthttp"
)

// Service import usecase
type Service struct {
	c        *fasthttp.Client
	basePath string
}

// NewService creates a new service
func NewService(basePath string) *Service {
	return &Service{
		c:        &fasthttp.Client{},
		basePath: basePath,
	}
}

// ImportProduct imports products
func (s *Service) ImportProduct(tenantID id.ID,
	p []*glad.Product,
) ([]*glad.ProductResponse, error) {
	path := s.basePath + "/v1/products/import"
	l.Log.Debugf("path=%v", path)

	jsonProducts, err := json.Marshal(p)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(path)
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.Header.Set(common.HttpHeaderTenantID, tenantID.String())
	req.SetBody([]byte(jsonProducts))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = s.c.Do(req, resp)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		l.Log.Warnf("Unable to import product. resp=%v, status=%v",
			string(resp.Body()), resp.StatusCode())
		return nil, err
	}

	var gResponse []*glad.ProductResponse
	err = json.Unmarshal(resp.Body(), &gResponse)
	if err != nil {
		l.Log.Warnf("Unable to decode import product response. err=%v", err)
		return nil, err
	}

	return gResponse, err
}

// ImportCenter imports centers
func (s *Service) ImportCenter(tenantID id.ID,
	p []*glad.Center,
) ([]*glad.CenterResponse, error) {
	path := s.basePath + "/v1/centers/import"
	l.Log.Debugf("path=%v", path)

	jsonCenters, err := json.Marshal(p)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(path)
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.Header.Set(common.HttpHeaderTenantID, tenantID.String())
	req.SetBody([]byte(jsonCenters))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = s.c.Do(req, resp)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		l.Log.Warnf("Unable to import center. resp=%v, status=%v",
			string(resp.Body()), resp.StatusCode())
		return nil, err
	}

	var gResponse []*glad.CenterResponse
	err = json.Unmarshal(resp.Body(), &gResponse)
	if err != nil {
		l.Log.Warnf("Unable to decode import center response. err=%v", err)
		return nil, err
	}

	return gResponse, err
}

// ImportAccount imports accounts
func (s *Service) ImportAccount(tenantID id.ID,
	p []*glad.Account,
) ([]*glad.AccountResponse, error) {
	path := s.basePath + "/v1/accounts/import"
	l.Log.Debugf("path=%v", path)

	jsonAccounts, err := json.Marshal(p)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(path)
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.Header.Set(common.HttpHeaderTenantID, tenantID.String())
	req.SetBody([]byte(jsonAccounts))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = s.c.Do(req, resp)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		l.Log.Warnf("Unable to import account. resp=%v, status=%v",
			string(resp.Body()), resp.StatusCode())
		return nil, err
	}

	var gResponse []*glad.AccountResponse
	err = json.Unmarshal(resp.Body(), &gResponse)
	if err != nil {
		l.Log.Warnf("Unable to decode import account response. err=%v", err)
		return nil, err
	}

	return gResponse, err
}

// ImportCourse imports courses
func (s *Service) ImportCourse(tenantID id.ID,
	p []*glad.Course,
) ([]*glad.CourseResponse, error) {
	path := s.basePath + "/v1/courses/import"
	l.Log.Debugf("path=%v", path)

	jsonCourses, err := json.Marshal(p)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(path)
	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.Header.Set(common.HttpHeaderTenantID, tenantID.String())
	req.SetBody([]byte(jsonCourses))

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = s.c.Do(req, resp)
	if err != nil {
		l.Log.Errorf("err=%v", err)
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		l.Log.Warnf("Unable to import course. resp=%v, status=%v",
			string(resp.Body()), resp.StatusCode())
		return nil, err
	}

	var gResponse []*glad.CourseResponse
	err = json.Unmarshal(resp.Body(), &gResponse)
	if err != nil {
		l.Log.Warnf("Unable to decode import course response. err=%v", err)
		return nil, err
	}

	return gResponse, err
}
