package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oktavarium/gomart/internal/app/internal/server/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	server, err := NewServer("db_addr", "accrual_addr", "test")
	assert.NoError(t, err)

	ts := httptest.NewServer(server)
	defer ts.Close()

	resp, token, body := testRequest(
		t,
		ts,
		"POST",
		"application/json",
		"",
		"/api/user/register",
		model.User{Login: "user", Password: "pass"},
	)

	require.Equal(t, 200, resp.StatusCode, "register user")
	require.NotEmpty(t, token, "register user")
	require.Empty(t, body, "register user")

	tableTests := []struct {
		name          string
		method        string
		contentType   string
		token         string
		path          string
		body          any
		status        int
		respToken     string
		emptyRespBody bool
	}{
		{"post order",
			"POST",
			"text/plain",
			token,
			"/api/user/orders",
			"4561261212345467",
			202,
			"",
			true,
		},
		{"post order",
			"POST",
			"text/plain",
			token,
			"/api/user/orders",
			"4561261212345467",
			200,
			"",
			true,
		},
		{"get orders",
			"GET",
			"text/plain",
			token,
			"/api/user/orders",
			"",
			200,
			"",
			false,
		},
	}

	for _, test := range tableTests {
		resp, token, body := testRequest(t, ts, test.method, test.contentType, test.token, test.path, test.body)
		require.Equal(t, test.status, resp.StatusCode, test.name)
		require.Equal(t, token, test.respToken, test.name)
		if test.emptyRespBody {
			require.Empty(t, body, test.name)
		} else {
			require.NotEmpty(t, body, test.name)
			fmt.Println(body)
		}
	}
}

func testRequest(
	t *testing.T,
	ts *httptest.Server,
	method string,
	contentType string,
	token string,
	path string,
	body any,
) (*http.Response, string, string) {
	var reader *bytes.Reader
	if contentType == "application/json" {
		raw, err := json.Marshal(body)
		require.NoError(t, err)
		reader = bytes.NewReader(raw)
	} else {
		reader = bytes.NewReader([]byte(body.(string)))
	}

	req, err := http.NewRequest(method, ts.URL+path, reader)
	require.NoError(t, err)

	req.Header.Set("Content-Type", contentType)
	if len(token) != 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	responseToken := resp.Header.Get("Authorization")

	return resp, responseToken, string(respBody)
}
