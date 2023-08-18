package commands_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/quantcdn/backend-init/internal/commands"
	"github.com/stretchr/testify/assert"
)

func TestVerifyRun_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	v := commands.Verify{
		Url:     server.URL,
		Delay:   5,
		Retries: 3,
	}
	err := v.Run()
	assert.Nil(t, err)
}

func TestVerifyRun_Fail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer server.Close()

	v := commands.Verify{
		Url:     server.URL,
		Delay:   5,
		Retries: 1,
	}
	err := v.Run()
	assert.Error(t, err)
}
