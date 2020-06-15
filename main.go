package main

import (
	"majixProject/controllers"
	"majixProject/majix"
)

func main()  {

	router := majix.NewRouter()
	router.Add("^/list/(?P<method>[0-9.]+)$",controllers.Index)
	router.Add("^/v2$",controllers.Index2)
	router.Start(":8080")
}

