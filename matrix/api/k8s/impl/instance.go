package impl

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/view"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource"
	"github.com/kataras/iris/context"
)
func InstanceReplica(ctx context.Context) apibase.Result{
	req := &view.InstanceReplicaReq{}
	ctx.ReadJSON(req)
	productName := ctx.Params().Get("product_name")
	serviceName := ctx.Params().Get("servce_name")
	req.ProductName =productName
	req.ServiceName = serviceName
	clusterid := ctx.GetCookie(view.ClusterId)
	return resource.InstanceReplica(ctx,clusterid,req)
}
