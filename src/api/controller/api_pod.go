package controller

import (
	"context"

	"net/http"

	"github.com/jake-willog/go-k8s-api.git/src/api/model"
	"github.com/jake-willog/go-k8s-api.git/src/config"
	"github.com/jake-willog/go-k8s-api.git/src/logger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetPodInfo(w http.ResponseWriter, r *http.Request) {
	reqId := getRequestId(w, r)
	logger.Debugf(reqId, "user/v1/pod/info GET started")

	client, err := config.LoadKubeConfig()
	if err != nil {
		logger.Errorf(reqId, "Failed to get kubeconfig %s", err)
	}

	pods, _ := client.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})

	podItems := make([]model.Pod, 0)
	for _, item := range pods.Items {
		podItems = append(podItems, model.Pod{
			Name: item.GetName(),
			Node: item.Spec.NodeName,
		})
	}

	resp := newOkResponse(w, reqId, "Ok")
	resp.Pods = &podItems
	writeResponse(reqId, w, resp)
}
