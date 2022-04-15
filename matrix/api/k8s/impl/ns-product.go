package impl

import (
	apibase "dtstack.com/dtstack/easymatrix/go-common/api-base"
	"dtstack.com/dtstack/easymatrix/matrix/api/k8s/view"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/resource"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"github.com/kataras/iris/context"
)

func GetParentProducts(ctx context.Context) apibase.Result{
	log.Debugf("[ns-product]: %s",ctx.Request().RequestURI)
	ns := ctx.Params().Get("namespace")
	clusterid := ctx.GetCookie(view.ClusterId)
	rsp,err := resource.GetParentProductList(ns,clusterid)
	if err != nil{
		return err
	}
	return rsp
}

func GetProducts(ctx context.Context) apibase.Result{
	log.Debugf("[ns-product]: %s",ctx.Request().RequestURI)
	parentProduct := ctx.Params().Get("parent_product_name")
	namespace := ctx.Params().Get("namespace")
	clusterid := ctx.GetCookie(view.ClusterId)
	rsp,err := resource.GetProductList(namespace,clusterid,parentProduct)
	if err != nil{
		return err
	}
	return rsp
}
//
func GetServiceList(ctx context.Context) apibase.Result{
	log.Debugf("[ns-product]: %s",ctx.Request().RequestURI)
	namespace := ctx.Params().Get("namespace")
	parentProduct := ctx.Params().Get("parent_product_name")
	productName := ctx.Params().Get("product_name")
	clusterid := ctx.GetCookie(view.ClusterId)
	rsp,err := resource.GetServiceList(namespace,clusterid,parentProduct,productName)
	if err != nil{
		return err
	}
	return rsp
}

func GetService(ctx context.Context) apibase.Result{
	log.Debugf("[ns-product]: %s",ctx.Request().RequestURI)
	namespace := ctx.Params().Get("namespace")
	parentProduct := ctx.Params().Get("parent_product_name")
	productName := ctx.Params().Get("product_name")
	servicename := ctx.Params().Get("service_name")
	clusterid := ctx.GetCookie(view.ClusterId)
	rsp,err := resource.GetService(ctx,namespace,clusterid,parentProduct,productName,servicename)
	if err != nil{
		return err
	}
	return rsp
}
