package model

type Pod struct {
	Name      string `json:"name,omitempty"`
	NodeName  string `json:"nodename,omitempty"`
	NameSpace string `json:"namespace,omitempty"`
}
