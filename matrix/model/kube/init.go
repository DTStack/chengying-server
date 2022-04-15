package kube

var localtbs = []KubeTable{
	DeployClusterList,
	DeployClusterK8sAvailable,
	DeployNamespaceList,
	DeployClusterImageStore,
	DeployNamespaceClient,
	ImportInitMoudle,
	DeployClusterProductRel,
	DeployNamespaceEvent,
	WorkloadDefinition,
	WorkloadPart,
	WorkloadStep,
}
//prepare sql
func Build() error{
	for _,tb := range localtbs{
		if err := tb.Prepare();err!= nil{
			return err
		}
	}
	return nil
}

