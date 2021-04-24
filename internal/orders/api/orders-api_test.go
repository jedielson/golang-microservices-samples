package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type OrdersApiHandlerSuite struct {
	suite.Suite

	ctx    context.Context
	router *mux.Router

	req *http.Request
	res *httptest.ResponseRecorder

	urls OrdersApiUrls
}

func (s *OrdersApiHandlerSuite) SetupTest() {
	s.ctx = context.Background()
	s.res = httptest.NewRecorder()
	s.router = mux.NewRouter()
	NewOrdersApi(s.router)
	s.urls = OrdersApiUrls{}.Get()
}

func (s *OrdersApiHandlerSuite) TestPlaceShouldReturn200() {
	// arrange
	s.req = httptest.NewRequest(http.MethodPost, s.urls.Place, nil)

	// act
	s.router.ServeHTTP(s.res, s.req)

	// assert
	s.Assert().Equal(http.StatusOK, s.res.Code)
}

func TestOrdersApiHandlerSuite(t *testing.T) {
	suite.Run(t, new(OrdersApiHandlerSuite))
}
