package hello_test

import (
	"github.com/modanisa/bootcamp-project-go/internal/hello"
	"github.com/modanisa/bootcamp-project-go/pkg/test"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	app := test.NewFiberApp()

	handler := hello.NewHandler()
	handler.RegisterRoutes(app)

	req := httptest.NewRequest(http.MethodGet, "/hello", http.NoBody)
	res, err := app.Test(req)
	assert.Nil(t, err)
	defer res.Body.Close()

	resBytes, _ := io.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "Hello, World 👋!", string(resBytes))
}
