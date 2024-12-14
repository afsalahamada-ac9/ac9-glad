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
	l.Log.Infof("path=%v", path)

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

	// Print response body
	println(string(resp.Body()))

	if resp.StatusCode() != fasthttp.StatusOK {
		l.Log.Warnf("Unable to import product. resp=%v, status=%v",
			string(resp.Body()), resp.StatusCode())
		return nil, err
	}

	var gResponse []*glad.ProductResponse
	err = json.Unmarshal(resp.Body(), &gResponse)
	if err != nil {
		l.Log.Warnf("Unable to decode import response. err=%v", err)
		return nil, err
	}

	return gResponse, err
}
