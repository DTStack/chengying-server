package view

type NamespacePingReq struct {
	Namespace 	string `json:"namespace"`
	Ip 			string	`json:"ip"`
	Port		string	`json:"port"`
}
