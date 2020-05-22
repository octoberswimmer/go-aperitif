package echoapitest

import (
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
)

func Serve(e *echo.Echo, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	e.ServeHTTP(rr, req)
	return rr
}
