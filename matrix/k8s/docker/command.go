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

package docker

import (
	"bufio"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func Login(username, address, password string) error {
	log.Debugf("docker login ...")
	login := exec.Command("docker", "login", "-u", username, address, "-p", password)
	login.Stdout = os.Stdout
	login.Stderr = os.Stderr
	if err := login.Run(); err != nil {
		return err
	}
	return nil
}

func OutputStdLog(std io.Reader, deployUUID string) {
	//实时循环读取输出流中的一行内容
	reader := bufio.NewReader(std)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		log.OutputInfof(deployUUID, "%v", line)
	}
}

func Load(file, deployUuid string) error {
	log.Debugf("docker load ...")
	load := exec.Command("docker", "load", "-i", file)
	stdout, err := load.StdoutPipe()
	if err == nil {
		go OutputStdLog(stdout, deployUuid)
	}
	stderr, err := load.StderrPipe()
	if err == nil {
		go OutputStdLog(stderr, deployUuid)
	}
	if err := load.Run(); err != nil {
		return err
	}
	return nil
}

func Tag(new, old string) error {
	tag := exec.Command("docker", "tag", old, new)
	tag.Stdout = os.Stdout
	tag.Stderr = os.Stderr
	if err := tag.Run(); err != nil {
		return err
	}
	return nil
}

func Push(image, deployUuid string) error {
	push := exec.Command("docker", "push", image)
	stdout, err := push.StdoutPipe()
	if err == nil {
		go OutputStdLog(stdout, deployUuid)
	}
	stderr, err := push.StderrPipe()
	if err == nil {
		go OutputStdLog(stderr, deployUuid)
	}
	if err := push.Run(); err != nil {
		return err
	}
	return nil
}

func newTransport(transport http.RoundTripper, registryURL, username, password string) *registry.Registry {
	transport = registry.WrapTransport(transport, registryURL, username, password)
	registry := &registry.Registry{
		URL: registryURL,
		Client: &http.Client{
			Transport: transport,
		},
		Logf: registry.Log,
	}
	return registry
}

func NewRegClient(registryURL, username, password string) (*registry.Registry, error) {
	transport := http.DefaultTransport
	url := fmt.Sprintf("http://%s", registryURL)
	registry := newTransport(transport, url, username, password)
	if err := registry.Ping(); err != nil {
		url = fmt.Sprintf("https://%s", registryURL)
		registry = newTransport(transport, url, username, password)
		if err := registry.Ping(); err != nil {
			return nil, err
		}
	}
	return registry, nil
}
