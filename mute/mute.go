package main

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	input := flag.String("n", "operator.dtstack.com_moles_crd.yaml", "The name of the file to be modified")
	output := flag.String("o", "operator.dtstack.com_moles_crd.yaml", "The name of the modified file")
	flag.Parse()

	obj := make(map[interface{}]interface{})
	yamlFile, _ := ioutil.ReadFile(*input)
	_ = yaml.Unmarshal(yamlFile, &obj)
	in := RemoveDescription(obj)
	out, _ := yaml.Marshal(in)

	f, _ := os.Create(*output)
	defer f.Close()
	_, _ = f.Write(out)
}

func RemoveDescription(obj interface{}) interface{} {
	switch valueType := obj.(type) {
	case map[interface{}]interface{}:
		result := make(map[interface{}]interface{})
		for k, v := range valueType {
			if k == "description" {
				continue
			}
			result[k] = RemoveDescription(v)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(valueType))
		for i, v := range valueType {
			result[i] = RemoveDescription(v)
		}
		return result
	default:
		return valueType
	}
}
