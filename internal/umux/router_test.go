package umux

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

const panicUrl = "/panic"

type PanicMiddlewareHandlerSuite struct {
	suite.Suite

	ctx    context.Context
	router *mux.Router

	req *http.Request
	res *httptest.ResponseRecorder
}

func (s *PanicMiddlewareHandlerSuite) SetupTest() {
	s.ctx = context.Background()
	s.res = httptest.NewRecorder()
	s.router = NewMuxRouter()

	s.router.HandleFunc(panicUrl, PanicHandler()).Methods(http.MethodGet)
}

func PanicHandler() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		panic("Deu ruim")
	}
}

func (s *PanicMiddlewareHandlerSuite) TestShouldReturn500IfPanic() {
	// arrange
	s.req = httptest.NewRequest(http.MethodGet, panicUrl, nil)

	// act
	s.router.ServeHTTP(s.res, s.req)

	// assert
	s.Assert().Equal(http.StatusInternalServerError, s.res.Code)
}

func TestPanicMiddlewareHandlerSuite(t *testing.T) {
	suite.Run(t, new(PanicMiddlewareHandlerSuite))
}
