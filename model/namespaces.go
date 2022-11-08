package model

type Namespace struct {
	NamespaceID   string `json:"namespace"` // generated ID or custom string
	Namespace     string `json:"namespaceShowName"`
	NamespaceDesc string `json:"namespaceDesc"`
	Quota         int    `json:"quota"`
	ConfigCount   int    `json:"configCount"`
	Type          int    `json:"type"`
}
