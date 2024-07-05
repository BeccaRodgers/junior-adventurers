package controller

import (
	"bytes"
	"context"
	approvals "github.com/approvals/go-approval-tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yosssi/gohtml"
	"junior-adventurers/fixtures"
	"junior-adventurers/memdb"
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

func TestHandler_NewMemberPage(t *testing.T) {
	approveGetPage(t, "/members")
}

func approveGetPage(t *testing.T, path string) {
	t.Helper()

	guilds := memdb.NewGuildRepository()
	require.NoError(t, fixtures.InsertGuilds(context.TODO(), guilds))
	members := memdb.NewMemberRepository()
	require.NoError(t, fixtures.InsertMembers(context.TODO(), members))

	handler := Handler(guilds, members)
	request := httptest.NewRequest(http.MethodGet, path, nil)
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	bodyBytes := gohtml.FormatBytes(recorder.Body.Bytes())
	var formattedOutput bytes.Buffer
	formattedOutput.Write(bodyBytes)
	approvals.Verify(t, &formattedOutput)
}
