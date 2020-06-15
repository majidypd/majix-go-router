package path

import (
	"github.com/my/repo/controllers"
	"github.com/my/repo/majix"
	"net/http"
)

func Web(r *majix.RouteManager) {
	r.Get("test/(?P<id>[0-9]+)$", func(w http.ResponseWriter, req *http.Request, u *majix.Application) {
		controllers.Index(w, req, u)
	}).Middleware(checkLogin)

}

func checkLogin(f majix.Handler) majix.Handler {
	return func(w http.ResponseWriter, r *http.Request, p *majix.Application) {
		f(w, r, p)
	}
}
