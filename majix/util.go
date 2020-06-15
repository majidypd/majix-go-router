package majix

import (
	"net/http"
	_ "net/http"
)

type Util struct {
	sessionManager *SessionManager
	route          *Route
}

func (u *Util) Session(w http.ResponseWriter, r *http.Request) SessionInterface {

	return u.sessionManager.Session(w, r)

}

func (u *Util) Param(key string) string {
	return u.route.urlParams[key]
}
