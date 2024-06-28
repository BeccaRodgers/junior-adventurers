package controllers

import (
	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Homepage(t *testing.T) {
	approveGetPage(t, "/")
}

func TestHandler_GuildPage(t *testing.T) {
	approveGetPage(t, "/guilds/1")
}

func approveGetPage(t *testing.T, path string) {
	t.Helper()

	handler := Handler()
	request := httptest.NewRequest(http.MethodGet, path, nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	approvals.Verify(t, recorder.Body)
}
