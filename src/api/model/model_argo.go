package model

type Connection struct {
	Address string
	Token   string
}

type ArgoCluster struct {
	Name             string
	ConnectionStatus string
	ServerVersion    string
}

type ArgoProject struct {
	Name       string
	Namespace  string
	RepoURL    string
	SyncStatus string
}
