package view

type NamespacedProductsRsp struct {
	Size 		int 		`json:"size"`
	Products 	Products	`json:"products"`
}

type Products map[ParentProductName]ParentProduct

type ParentProduct map[ProudctName]Product

type Product map[ServiceName]*Service

type Service struct {
	Version				string	`json:"version"`
	Group				string	`json:"group"`
	HealthState			string	`json:"health_state"`
	HealthStateCount	int		`json:"health_state_count"`
	ServiceStatus		string	`json:"service_status"`
	ServiecStatusCount	int		`json:"service_status_count"`
	IsJob 				bool
	WorkloadType 		string
}

type ServiceName string
type ProudctName string
type ParentProductName string

func (sn ServiceName) String() string{
	return string(sn)
}

func (pn ProudctName)String() string{
	return string(pn)
}

func (ppn ParentProductName)String() string{
	return string(ppn)
}
