package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type HcApiHandlerSuite struct {
	suite.Suite

	ctx    context.Context
	router *mux.Router

	req *http.Request
	res *httptest.ResponseRecorder
}

func (s *HcApiHandlerSuite) SetupTest() {
	s.ctx = context.Background()
	s.res = httptest.NewRecorder()
	s.router = mux.NewRouter()
	NewHcApi(s.router)
}

func (s *HcApiHandlerSuite) TestPingShoudReturnPong() {
	// arrange
	s.req = httptest.NewRequest(http.MethodGet, apiPing, nil)

	// act
	s.router.ServeHTTP(s.res, s.req)

	// assert
	result := string(s.res.Body.String())
	s.Assert().Equal(http.StatusOK, s.res.Code)
	s.Assert().Equal("Pong", result)
}

func TestHcApiHandlerSuite(t *testing.T) {
	suite.Run(t, new(HcApiHandlerSuite))
}
