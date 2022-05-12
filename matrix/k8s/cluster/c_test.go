// Licensed to Apache Software Foundation(ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation(ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cluster

import (
	"dtstack.com/dtstack/easymatrix/matrix/k8s/constant"
	modelkube "dtstack.com/dtstack/easymatrix/matrix/model/kube"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	info := &GeneratorInfo{
		Type:        constant.TYPE_SELF_BUILD,
		ClusterInfo: &modelkube.ClusterInfo{
			Id:            135,
			Name:          "em_mao_self",
		},
	}
	bts,_:=GetTemplateFile(info,false)
	fmt.Println(string(bts))
}

func Test2(t *testing.T) {

}
