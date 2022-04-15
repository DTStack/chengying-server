package apibase

// 1: Administrator, 2: Cluster Operator, 4: Cluster Reader

// permission7 = Administrator & Cluster Operator & Cluster Reader
// permission3 = Administrator & Cluster Operator
// permission1 = Administrator

import (
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

const (
	Administrator   = 1
	ClusterOperator = 2
	ClusterReader   = 4
)

func CheckPermission3(ctx context.Context) {
	rulePermission := Administrator | ClusterOperator
	userPermission, err := GetTokenUserPermission(ctx)
	if err != nil {
		log.Errorf(err.Error())
	}
	if rulePermission&userPermission == 0 {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.Next()
}

func CheckPermission1(ctx context.Context) {
	rulePermission := Administrator
	userPermission, err := GetTokenUserPermission(ctx)
	if err != nil {
		log.Errorf(err.Error())
	}
	if rulePermission&userPermission == 0 {
		ctx.StatusCode(iris.StatusForbidden)
		return
	}
	ctx.Next()
}
