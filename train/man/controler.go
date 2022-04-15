package man

import (
	"dtstack.com/dtstack/easymatrix/matrix/agent"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"dtstack.com/dtstack/easymatrix/matrix/service"
	"dtstack.com/dtstack/easymatrix/train/define"
	"dtstack.com/dtstack/easymatrix/train/util"
	"fmt"
	"os/exec"
)

var (
	TrainMan = &trainman{}
)

type trainman struct {
	dockerCompose  string
	orderConfigMap map[string]map[string]map[string]int
	actionOrderMap map[string]map[string]map[int][]model.DeployInstanceInfo
}

type Controller interface {
	Start() error
	Stop() error
	Down() error
	Up() error
	Exec(command string) error
}

func InitDockerCompose(compose string, operators []define.Operator) {
	TrainMan.dockerCompose = compose
	TrainMan.orderConfigMap = make(map[string]map[string]map[string]int)
	TrainMan.actionOrderMap = make(map[string]map[string]map[int][]model.DeployInstanceInfo)
	for _, o := range operators {
		if _, ok := TrainMan.orderConfigMap[o.Product]; !ok {
			TrainMan.orderConfigMap[o.Product] = make(map[string]map[string]int)
		}
		if _, ok := TrainMan.orderConfigMap[o.Action]; !ok {
			TrainMan.orderConfigMap[o.Product][o.Action] = make(map[string]int)
		}
		for index, name := range o.Order {
			TrainMan.orderConfigMap[o.Product][o.Action][name] = index
		}
	}
}

func InitAgentClient(host string, compose string) error {
	agent.InitAgentClient(host)
	return nil
}

func (this *trainman) Start() error {
	this.Operator("start")
	return nil
}

func (this *trainman) Stop() error {
	this.Operator("stop")
	return nil
}

func (this *trainman) Operator(action string) error {
	list, err := model.DeployInstanceList.GetInstanceList()
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	var unOrderInstanceList []model.DeployInstanceInfo
	for _, instance := range list {
		product, err := model.DeployProductList.GetProductInfoById(instance.Pid)
		if err != nil {
			fmt.Errorf("get product info err: %v", err.Error())
			continue
		}
		if _, ok := this.actionOrderMap[product.ProductName]; !ok {
			this.actionOrderMap[product.ProductName] = make(map[string]map[int][]model.DeployInstanceInfo)
		}
		if _, ok := this.actionOrderMap[product.ProductName][action]; !ok {
			this.actionOrderMap[product.ProductName][action] = make(map[int][]model.DeployInstanceInfo)
		}
		if index, ok := this.orderConfigMap[product.ProductName][action][instance.ServiceName]; ok {
			this.actionOrderMap[product.ProductName][action][index] = append(this.actionOrderMap[product.ProductName][action][index], instance)
		} else {
			unOrderInstanceList = append(unOrderInstanceList, instance)
		}
	}
	for product, actionOrder := range this.actionOrderMap {
		fmt.Printf("%s product %v services in order ...", action, product)
		for index, insts := range actionOrder[action] {
			fmt.Printf("%s %v:", action, index)
			for _, inst := range insts {
				fmt.Printf("\t%s %s of %s with status: %s \n", action, inst.ServiceName, inst.Ip, inst.Status)
				servicer, err := service.NewServicer(inst.Pid, 0, inst.ServiceName)
				if err != nil {
					fmt.Printf("\t\t%v failed %v\n", action, err)
					continue
				}
				err = servicer.Start()
				if err != nil {
					fmt.Printf("\t\t%v failed %v\n", action, err)
					continue
				}
				fmt.Println("\t\t%v success", action)
			}
		}
	}

	fmt.Printf("%v unOrdered services in order ...", action)

	for _, instance := range unOrderInstanceList {
		fmt.Printf("\t%s %s of %s with status: %s \n", action, instance.ServiceName, instance.Ip, instance.Status)
		servicer, err := service.NewServicer(instance.Pid, 0, instance.ServiceName)
		if err != nil {
			fmt.Printf("\t\t%v failed %v\n", action, err)
			continue
		}
		err = servicer.Start()
		if err != nil {
			fmt.Printf("\t\t%s failed %v\n", action, err)
			continue
		}
		fmt.Printf("\t\t%s success", action)
	}
	return nil
}

func (this *trainman) Down() error {
	args := append([]string{"-f"}, this.dockerCompose, "down")
	cmd := exec.Command("docker-compose", args...)
	stdBuf := &util.PrefixSuffixSaver{N: 512 << 10}
	cmd.Stdout = stdBuf
	cmd.Stderr = stdBuf
	fmt.Printf("run %v", cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\tdocker-compose down err: %v, %v", err.Error(), string(stdBuf.Bytes()))
	}
	fmt.Printf("\tdocker-compose down: %v", string(stdBuf.Bytes()))
	return nil
}

func (this *trainman) Up() error {
	args := append([]string{"-f"}, this.dockerCompose, "up", "-d")
	cmd := exec.Command("docker-compose", args...)
	stdBuf := &util.PrefixSuffixSaver{N: 512 << 10}
	cmd.Stdout = stdBuf
	cmd.Stderr = stdBuf
	fmt.Printf("run %v", cmd.Args)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("\tdocker-compose up err: %v, %v", err.Error(), string(stdBuf.Bytes()))
		return err
	}
	fmt.Printf("\tdocker-compose up: %v", string(stdBuf.Bytes()))
	return nil
}

func (this *trainman) Exec(command string) error {
	fmt.Printf("Train Man get command: %s\n", command)
	switch command {
	case "start":
		this.Start()
	case "stop":
		this.Stop()
	case "em-down":
		this.Down()
	case "em-up":
		this.Up()
	default:
		fmt.Println("wrong command, do nothing! Surpport start|stop|em-down|em-up")
	}
	return nil
}
