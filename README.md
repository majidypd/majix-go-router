# majix/router




```go
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


```

### Controller


```go
package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprint(w, "v1",params["method"])
}

func Index2(w http.ResponseWriter, r *http.Request, params map[string]string) {
	fmt.Fprint(w, "v2")
}

```

### Middleware


```go
TODO
```

```go

```


```go

```
## License


