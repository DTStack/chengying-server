package deploy

import (
	"dtstack.com/dtstack/easymatrix/matrix/base"
	"dtstack.com/dtstack/easymatrix/matrix/k8s/docker"
	"dtstack.com/dtstack/easymatrix/matrix/log"
	"dtstack.com/dtstack/easymatrix/matrix/model"
	"dtstack.com/dtstack/easymatrix/schema"
	"os"
	"path/filepath"
	"strings"
)

const (
	IMAGE_DIR    = "images"
	IMAGE_SUFFIX = ".tar"
)

func PushImages(store model.ImageStore, sc *schema.SchemaConfig, deployUUID string) error {
	log.Infof("starting push images to registry %v", store.Address)
	log.OutputInfof(deployUUID, "starting push images to registry %v", store.Address)
	err := docker.Login(store.Username, store.Address, store.Password)
	if err != nil {
		log.Errorf("docker login err: %v", err.Error())
		log.OutputInfof(deployUUID, "docker login error: %v", err.Error())
		return err
	}
	log.Infof("docker login success: %v", store.Address)
	log.OutputInfof(deployUUID, "docker login success: %v", store.Address)
	regURL := strings.Split(store.Address, "/")[0]
	hub, err := docker.NewRegClient(regURL, store.Username, store.Password)
	if err != nil {
		log.Errorf("creat registry client err: %v", err.Error())
		log.OutputInfof(deployUUID, "creat registry client err: %v", err.Error())
	}
	for name := range sc.Service {
		baseDir := filepath.Join(base.WebRoot, sc.ProductName, sc.ProductVersion, name, IMAGE_DIR)
		if sc.Service[name].Instance == nil || sc.Service[name].Instance.Image == "" {
			log.Infof("")
			continue
		}
		if sc.Service[name].BaseProduct != "" || sc.Service[name].BaseService != "" {
			continue
		}
		sourceImg := sc.Service[name].Instance.Image
		log.Infof("push images with images dir: %v", baseDir)
		log.OutputInfof(deployUUID, "push images with images dir: %v", baseDir)
		err = filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
			log.Infof("find file: %v", path)
			log.OutputInfof(deployUUID, "find file: %v", path)
			if err != nil {
				return err
			}
			if baseDir == path {
				return nil
			}
			if info.IsDir() {
				return nil
			}
			if !strings.HasSuffix(path, IMAGE_SUFFIX) {
				log.Infof("not regular image file: %v", path)
				log.OutputInfof(deployUUID, "not regular image file: %v", path)
				return nil
			}
			err = docker.Load(path, deployUUID)
			if err != nil {
				return err
			}
			log.Infof("load image %v success", path)
			log.OutputInfof(deployUUID, "load image %v success", path)
			newImg := store.Address + "/" + sourceImg
			err = docker.Tag(newImg, sourceImg)
			if err != nil {
				//the image can be changed in the front, should not affect the other image's push
				log.OutputInfof(deployUUID, "use the declare image %s in the front", sourceImg)
				return nil
			}
			log.Infof("tag image %v success", newImg)
			log.OutputInfof(deployUUID, "tag image %v success", newImg)
			var exists bool
			imgName := strings.SplitN(strings.Split(newImg, ":")[0], "/", 2)[1]
			imgTag := strings.Split(newImg, ":")[1]
			if hub != nil {
				tags, err := hub.Tags(imgName)
				if err != nil {
					log.Infof("search image tag err: %v", err)
				}
				for _, tag := range tags {
					if tag == imgTag {
						exists = true
						break
					}
				}
				if !exists {
					log.OutputInfof(deployUUID, "the image %v does not exist in the docker repository,start pushing the image ...", newImg)
					err = docker.Push(newImg, deployUUID)
					if err != nil {
						return err
					}
					log.Infof("push image %v success", path)
					log.OutputInfof(deployUUID, "push image %v success", path)
				} else {
					log.OutputInfof(deployUUID, "the image %v exists in the docker repository,skip pushing the image!", newImg)
				}
			} else {
				err = docker.Push(newImg, deployUUID)
				if err != nil {
					return err
				}
				log.Infof("push image %v success", path)
				log.OutputInfof(deployUUID, "push image %v success", path)
			}
			sc.Service[name].Instance.Image = newImg
			return nil
		})
		if err != nil {
			log.Errorf("push images err: %v", err.Error())
			log.OutputInfof(deployUUID, "push images error: %v", err.Error())
			return err
		}
	}
	return nil
}
