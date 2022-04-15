package view

type NamespaceSaveReq struct {
	Id          int    `json:"id" db:"id"`
	Type		string `json:"type" db:"type"`
	Namespace 	string `json:"namespace" db:"namespace"`
	RegistryId	int    `json:"registry_id" db:"registry_id"`
	Yaml		string `json:"yaml"`
	FileName    string `json:"file_name"`
	Ip			string `json:"ip" db:"ip"`
	Port		string `json:"port" db"port"`
	
}
