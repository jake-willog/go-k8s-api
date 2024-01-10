package config

import (
	"github.com/jake-willog/go-k8s-api.git/src/api/argocd"
	"github.com/jake-willog/go-k8s-api.git/src/api/model"
)

func LoadArgoConfig() (*argocd.Client, error) {
	argoconfig := model.Connection{
		Address: "argocd-server.argocd.svc.cluster.local:443",
		Token:   "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcmdvY2QiLCJzdWIiOiJmb286YXBpS2V5IiwibmJmIjoxNzA0MzYxOTk3LCJpYXQiOjE3MDQzNjE5OTcsImp0aSI6Ijg3ZWMzOTdkLTY4MzgtNDllNS05NjNkLTI4ZGJmMDA3ZThiZCJ9.Mv4fN7cGnebFeTd9orHU30gO33znspwBjkNdL0_cfn0",
	}

	client, err := argocd.NewClient(argoconfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}
