package controller

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Control Plane API\n\n"))
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetPodInfo",
		strings.ToUpper("Get"),
		"/v1/pod/info",
		GetPodInfo,
	},
	Route{
		"CreateDeployment",
		strings.ToUpper("Post"),
		"/v1/pod/create",
		CreateDeployment,
	},
	// Route{
	// 	"GetArgoProject",
	// 	strings.ToUpper("Get"),
	// 	"/v1/argo/project/info",
	// 	GetArgoProject,
	// },
	Route{
		"CreateArgoProject",
		strings.ToUpper("Post"),
		"/v1/argo/project/create",
		CreateArgoProject,
	},
	Route{
		"GetArgoCluster",
		strings.ToUpper("Get"),
		"/v1/argo/cluster/info",
		GetArgoCluster,
	},
	Route{
		"GetArgoApplication",
		strings.ToUpper("Get"),
		"/v1/argo/application/info",
		GetArgoApplication,
	},
	Route{
		"CreateArgoApplication",
		strings.ToUpper("Post"),
		"/v1/argo/application/create",
		CreateArgoApplication,
	},
}
