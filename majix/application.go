package majix

import (
	"reflect"
)



type ApplicationInterface interface {
	Bind(string, interface{})
	Resolve(name string) interface{}
	resolveType(name reflect.Type) Service
	Start(address string, f func(mapper *RouteManager))
}

type Application struct {
	Util         Util
	Services     map[string]Service
	ServicesType map[reflect.Type]Service
}

type Service struct {
	Type  reflect.Type
	Value interface{}
}

func (a *Application) Bind(name string, i interface{}) {
	service := Service{Type: reflect.TypeOf(i), Value: i}
	a.Services[name] = service
	a.ServicesType[reflect.TypeOf(i)] = service
}

func (a *Application) Resolve(name string) interface{} {
	service, _ := a.Services[name]
	return service.Value
}

func (a *Application) resolveType(name reflect.Type) Service {
	service, _ := a.ServicesType[name]
	return service
}

func (a *Application) Start(address string, f func(mapper *RouteManager)) {
	//-------------------Setup Session------------------------------------
	if sessionManager := a.resolveType(reflect.TypeOf(&SessionManager{})); sessionManager != (Service{}) {
		a.Util.sessionManager = sessionManager.Value.(*SessionManager)
	}
	//--------------------------------------------------------------------

	//-------------------Setup Router And Start App-----------------------
	if routeManager := a.resolveType(reflect.TypeOf(&RouteManager{})); routeManager != (Service{}) {
		router := routeManager.Value.(*RouteManager)
		f(router)
		router.Start(a, address)
	}
	//--------------------------------------------------------------------
}

func NewApp() ApplicationInterface {
	return &Application{
		Services:     map[string]Service{},
		ServicesType: map[reflect.Type]Service{},
	}
}
