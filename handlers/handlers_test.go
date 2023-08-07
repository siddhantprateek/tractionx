package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	handlers "tractionx/handlers"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBaseRoute(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	err := handlers.BaseRoute(ctx)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expectedJSON := `{"message":"Server is Healthy."}`
		assert.JSONEq(t, expectedJSON, rec.Body.String())
	}
}
