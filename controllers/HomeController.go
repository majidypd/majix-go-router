package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprint(w, "v1",params["method"])
}

func Index2(w http.ResponseWriter, r *http.Request,) {
	fmt.Fprint(w, "v2")
}
