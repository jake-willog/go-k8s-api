package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jake-willog/go-k8s-api.git/src/api/model"
	"github.com/jake-willog/go-k8s-api.git/src/config"
	"github.com/jake-willog/go-k8s-api.git/src/logger"
)

func GetArgoApplication(w http.ResponseWriter, r *http.Request) {
	reqId := getRequestId(w, r)
	logger.Debugf(reqId, "v1/argocd/application/info GET started")

	client, err := config.LoadArgoConfig()
	if err != nil {
		logger.Errorf(reqId, "Failed to get ArgoConfig %s", err)
	}

	applications, err := client.GetAppliction()

	applicationItems := make([]model.ArgoProject, 0)
	for _, application := range applications {
		applicationItems = append(applicationItems, model.ArgoProject{
			Name:       application.Name,
			NameSpace:  application.Spec.Destination.Namespace,
			RepoURL:    application.Spec.Source.RepoURL,
			SyncStatus: string(application.Status.Sync.Status),
		})
	}

	resp := newOkResponse(w, reqId, "Ok")
	resp.ArgoProjects = &applicationItems
	writeResponse(reqId, w, resp)
}

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

func CreateArgoProject(w http.ResponseWriter, r *http.Request) {
	reqId := getRequestId(w, r)
	logger.Debugf(reqId, "v1/argocd/project/create POST started")

	client, err := config.LoadArgoConfig()
	if err != nil {
		logger.Errorf(reqId, "Failed to get ArgoConfig %s", err)
		resp := newResponse(w, reqId, 500, "Internal Server Error")
		writeResponse(reqId, w, resp)
	}

	argoProjectCreateReq := &model.ArgoProjectCreateReq{}

	err = json.NewDecoder(r.Body).Decode(argoProjectCreateReq)
	if err != nil {
		logger.Errorf(reqId, "Failed to decode request body %s", err)
		resp := newResponse(w, reqId, 400, "Bad Request")
		writeResponse(reqId, w, resp)
		return
	}

	createProject, err := client.CreateProject(argoProjectCreateReq.Name)
	if err != nil {
		logger.Errorf(reqId, "Failed to create ArgoCD project: %s", err)
		resp := newResponse(w, reqId, 500, "Internal Server Error")
		writeResponse(reqId, w, resp)
		return
	}

	fmt.Println(createProject)

	resp := newOkResponse(w, reqId, "Ok")
	writeResponse(reqId, w, resp)
}
