package model

//func init() {
//	err := ConfigureMysqlDatabase("172.16.8.165", 3306, "root", "dtstack", "dtagent")
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}

//func TestCreateServiceIpNode(t *testing.T) {
//	node := ServiceIpNode{
//		ClusterId:   1,
//		ProductName: "product test",
//		ServiceName: "service 1",
//		Ip:          "127.0.0.1",
//		NodeId:      3,
//	}
//
//	if err := node.Create(); err != nil {
//		t.Error(err)
//	}
//}
//
//func TestGetServiceIpNodes(t *testing.T) {
//	nodes, err := GetServiceNodes(1, "product test", "service 1")
//	if err != nil {
//		t.Error(err)
//	}
//	bytes, _ := json.Marshal(nodes)
//	fmt.Printf("nodes: %s\n", bytes)
//}
//
//func TestGetServiceIpNode(t *testing.T) {
//	node, err := GetServiceIpNode(1, "product test", "service 1", "127.0.0.2")
//	if err != nil {
//		t.Error(err)
//	}
//	fmt.Printf("%v", node)
//}
