package caddy_ldap

import (
	"fmt"
	"github.com/caddyserver/caddy"
	"github.com/caddyserver/caddy/caddyhttp/httpserver"
	"net/http"
)

type CaddyLdapHandler struct {
	Next httpserver.Handler
}


func init() {
	caddy.RegisterPlugin("ldap", caddy.Plugin{
		ServerType: "http",
		Action:     Setup,
	})
}

func Setup(c *caddy.Controller) error {
	newMiddleWare := func(next httpserver.Handler) httpserver.Handler {
		return &CaddyLdapHandler{
			Next: next,
		}
	}
	// Add middleware
	cfg := httpserver.GetConfig(c)
	cfg.AddMiddleware(newMiddleWare)

	return nil
}

func (a *CaddyLdapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) (int, error) {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
		w.WriteHeader(http.StatusUnauthorized)
		return http.StatusUnauthorized, nil
	}

	fmt.Println(username, password, ok)
	return a.Next.ServeHTTP(w, r)

}

