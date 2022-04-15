package view

type NamespaceGetRsp struct {
	Id			int 	`json:"id"`
	Type 		string	`json:"type"`
	Namespace 	string	`json:"namespace"`
	Registry	int		`json:"registry"`
	Yaml		string	`json:"yaml"`
	FileName    string  `json:"file_name"`
}
