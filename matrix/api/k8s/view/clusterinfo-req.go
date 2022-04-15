package view

import "fmt"

type ClusterInfoReq struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Mode       int    `json:"mode"`
	Version    string `json:"version"`
	Desc       string `json:"desc"`
	Tags       string `json:"tags"`
	Configs    string `json:"configs"`
	Yaml       string `json:"yaml"`
	Status     int    `json:"status"`
	ErrorMsg   string `json:"errorMsg"`
	CreateUser string `json:"create_user"`
	NetworkPlugin NetWorkPlugin `json:"network_plugin"`
}


type NetWorkPlugin string

func (n NetWorkPlugin) String() string{
	return fmt.Sprintf("{\"network_plugin\":\"%s\"}",string(n))
}
