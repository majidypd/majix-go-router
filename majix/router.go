package majix

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
)

type RouterHandler func(http.ResponseWriter, *http.Request, map[string]string)

type RouteMapper struct {
	mapper map[*regexp.Regexp]*Route
}

type Route struct {
	handler interface{}
	pattern *regexp.Regexp
	kind    reflect.Type
}

func (r *RouteMapper) Add(pattern string, handler interface{}) {

	reg, _ := regexp.Compile(pattern)
	route := &Route{
		handler: handler,
		pattern: reg,
		kind:    reflect.TypeOf(handler),
	}
	r.mapper[reg] = route
}

func (r *RouteMapper) Start(address string) {
	http.ListenAndServe(address, r)
}

func NewRouter() *RouteMapper {
	return &RouteMapper{
		mapper: make(map[*regexp.Regexp]*Route),
	}
}

func (r *RouteMapper) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	for pattern, f := range r.mapper {

		if matches := pattern.FindStringSubmatch(req.URL.Path); matches != nil {
			urlParams := make(map[string]string)

			for i := 1; i < len(matches); i++ {
				urlParams[pattern.SubexpNames()[i]] = matches[i]
			}

			switch f.handler.(type) {
			case func(http.ResponseWriter, *http.Request, map[string]string):
				x := f.handler.(func(http.ResponseWriter, *http.Request, map[string]string))
				x(w, req, urlParams)
			case func(http.ResponseWriter, *http.Request):
				x := f.handler.(func(http.ResponseWriter, *http.Request))
				x(w, req)
			}

		}
	}
	fmt.Println(req.URL.Path)
}

type Request struct {
}
