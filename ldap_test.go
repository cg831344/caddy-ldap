package caddy_ldap

import (
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRequest(t *testing.T, handler CaddyLdapHandler, user string, path string, method string, code int) {
	r, _ := http.NewRequest(method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	if w.Code != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, w.Code, code)
	}
}

func TestNoneAuth(t *testing.T) {
	r, _ := http.NewRequest("get", "/base", nil)
	//r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	handler := CaddyLdapHandler{
		Next: httpserver.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
			return http.StatusOK, nil
		}),
	}
	handler.ServeHTTP(w, r)
	code:=401
	if w.Code != code {
		t.Errorf("%d, supposed to be %d",w.Code, code)
	}
}


func TestBasic(t *testing.T) {
	handler := CaddyLdapHandler{
		Next: httpserver.HandlerFunc(func(w http.ResponseWriter, r *http.Request) (int, error) {
			return http.StatusOK, nil
		}),
	}
	testRequest(t, handler, "alice", "/dataset1/resource1", "GET", 200)
}

