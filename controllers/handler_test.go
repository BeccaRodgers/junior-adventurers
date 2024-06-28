package controllers

import (
	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_Homepage(t *testing.T) {
	handler := Handler()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	approvals.Verify(t, recorder.Body)
}

func TestHandler_GuildPage(t *testing.T) {
	handler := Handler()
	request := httptest.NewRequest(http.MethodGet, "/guilds/1", nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	approvals.Verify(t, recorder.Body)
}
