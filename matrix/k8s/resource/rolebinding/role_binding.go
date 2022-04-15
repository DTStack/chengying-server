package rolebinding

import (
	"dtstack.com/dtstack/easymatrix/addons/easykube/pkg/client/base"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"encoding/json"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var GVK = schema.GroupVersionKind{
	Group:   "rbac.authorization.k8s.io",
	Version: "v1",
	Kind:    "RoleBinding",
}

func ToObject(bts []byte)(*rbacv1.RoleBinding,error){
	r,err := base.Schema.New(GVK)
	if err != nil{
		log.Errorf("[role_binding]: new object error: %v",err)
		return nil,err
	}
	err = json.Unmarshal(bts,r)
	if err!= nil{
		log.Errorf("[role_binding]: json %s unmarshal error: %v",string(bts),err)
		return nil,err
	}
	sa := r.(*rbacv1.RoleBinding)
	return sa,nil
}
