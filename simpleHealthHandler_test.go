package simple

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
)

func TestSimpleHandlerOK(t *testing.T) {
	req, _ := http.NewRequest("GET", "/health", nil)

	rr := httptest.NewRecorder()

	testee := SimpleHealthHandler{}

	testee.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}
