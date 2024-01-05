package model

type ArgoApplicationCreateReq struct {
	Name        string `json:"name,omitempty"`
	Project     string `json:"project,omitempty"`
	RepoURL     string `json:"repoURL,omitempty"`
	Path        string `json:"path,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	Namespace   string `json:"namespace,omitempty"`
}
