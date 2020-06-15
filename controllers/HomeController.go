package controllers

import (
	"fmt"
	"github.com/my/repo/majix"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request, app *majix.Application) {
	s := app.Util.Session(w, r)
	s.Set("test_session", time.Now().String())
	fmt.Fprint(w, s.Get("test_session"))
}
