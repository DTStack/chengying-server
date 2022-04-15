package role

import (
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"encoding/json"
	"k8s.io/apimachinery/pkg/runtime/schema"
	rbacv1 "k8s.io/api/rbac/v1"
)

var GVK = schema.GroupVersionKind{
	Group:   "rbac.authorization.k8s.io",
	Version: "v1",
	Kind:    "Role",
}

func ToObject(bts []byte)(*rbacv1.Role,error){
	r,err := base.Schema.New(GVK)
	if err != nil{
		log.Errorf("[role]: new object error: %v",err)
		return nil,err
	}
	err = json.Unmarshal(bts,r)
	if err!= nil{
		log.Errorf("[role]: json %s unmarshal error: %v",string(bts),err)
		return nil,err
	}
	role := r.(*rbacv1.Role)
	return role,nil
}
