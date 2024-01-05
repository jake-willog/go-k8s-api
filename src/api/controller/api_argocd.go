package controller

import (
	"net/http"

	"github.com/jake-willog/go-k8s-api.git/src/api/model"
	"github.com/jake-willog/go-k8s-api.git/src/config"
	"github.com/jake-willog/go-k8s-api.git/src/logger"
)

func GetArgoCluster(w http.ResponseWriter, r *http.Request) {
	reqId := getRequestId(w, r)
	logger.Debugf(reqId, "v1/argocd/cluster/info GET started")

	client, err := config.LoadArgoConfig()
	if err != nil {
		logger.Errorf(reqId, "Failed to get ArgoConfig %s", err)
	}
	clusters, err := client.GetClusters()

	clusterItems := make([]model.ArgoCluster, 0)
	for _, cluster := range clusters {
		// fmt.Println(cluster.Server, cluster.ConnectionState.Status, cluster.Info.ServerVersion)
		clusterItems = append(clusterItems, model.ArgoCluster{
			Name:             cluster.Name,
			ConnectionStatus: cluster.ConnectionState.Status,
			ServerVersion:    cluster.ServerVersion,
		})
	}

	resp := newOkResponse(w, reqId, "Ok")
	resp.ArgoClusters = &clusterItems
	writeResponse(reqId, w, resp)
}

// func main() {
// 	connection := argocd.Connection{
// 		Address: "localhost:8080",
// 		Token:   "my-foo-account-token",
// 	}

// 	client, err := argocd.NewClient(&connection)
// 	if err != nil {
// 		panic(err)
// 	}

// 	createProject, err := client.CreateProject("foo")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(createProject.UID)

// 	err = client.AddDestination(createProject.Name, "server", "namespace", "name")
// 	if err != nil {
// 		panic(err)
// 	}

// 	getProject, err := client.GetProject("foo")
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(getProject.Namespace)

// 	err = client.DeleteProject(getProject.Name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	clusters, err := client.GetClusters()
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, cluster := range clusters {
// 		fmt.Println(cluster.Name)
// 	}
// }
