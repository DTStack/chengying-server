package union

import modelkube "dtstack.com/dtstack/easymatrix/matrix/model/kube"

var localtbs = []modelkube.KubeTable{
	UnionT4T7,
}
func Build() error{
	for _,tb := range localtbs{
		if err := tb.Prepare();err!= nil{
			return err
		}
	}
	return nil
}
