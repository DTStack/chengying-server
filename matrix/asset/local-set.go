package asset

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"dtstack.com/dtstack/easymatrix/matrix/log"
)

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Errorf("[SetAssetWithLocalFile] getCurrentDirectory err: %v", err)
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func ResettemplatesWithLocalFile() error {
	file := ""
	for _, typ := range []string{"script_wrapper_install.sh", "environment_init.sh", "install_agentx.sh", "docker_environment_init.sh"} {
		file = typ
		if _, err := os.Stat(file); os.IsNotExist(err) {
			log.Infof("[InitInstallScriptShWithLocalFile] %v, warning: %v", typ, err)
			continue
		}
		fi, err := os.Open(file)
		if err != nil {
			log.Errorf("[SetAssetWithLocalFile] %v, err: %v", typ, err)
			continue
		}
		content, err := ioutil.ReadAll(fi)
		fi.Close()
		if err != nil {
			log.Errorf("[SetAssetWithLocalFile] %v, err: %v", typ, err)
			continue
		}
		switch typ {
		case "script_wrapper_install.sh":
			_templatesInstallScriptWrapperSh = content
		case "environment_init.sh":
			_templatesEnvironmentInitSh = content
		case "install_agentx.sh":
			_templatesInstallAgentxSh = content
		case "docker_environment_init.sh":
			_templatesDockerEnvironmentInitSh = content
		}
	}
	return nil
}

func ResetInstallAgentXShWithLocalFile() {
	file := ""
	typ := "install_agentx.sh"
	if len(getCurrentDirectory()) > 0 {
		file = getCurrentDirectory() + "/" + typ
	} else {
		file = typ
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Infof("[ResetInstallAgentXShWithLocalFile] %v, err: %v", typ, err)
		return
	}
	fi, err := os.Open(file)
	if err != nil {
		log.Infof("[ResetInstallAgentXShWithLocalFile] %v, err: %v", typ, err)
		return
	}
	content, err := ioutil.ReadAll(fi)
	fi.Close()
	if err != nil {
		log.Infof("[ResetInstallAgentXShWithLocalFile] %v, err: %v", typ, err)
		return
	}
	_templatesInstallAgentxSh = content
}

func ResetPatchUpdateScriptWithLocalFile() {
	file := ""
	typ := "patchupdate.sh"
	if len(getCurrentDirectory()) > 0 {
		file = getCurrentDirectory() + "/" + typ
	} else {
		file = typ
	}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		log.Infof("[ResetPatchUpdateScriptWithLocalFile] %v, err: %v", typ, err)
		return
	}
	fi, err := os.Open(file)
	if err != nil {
		log.Infof("[ResetPatchUpdateScriptWithLocalFile] %v, err: %v", typ, err)
		return
	}
	patchupdate_content, err := ioutil.ReadAll(fi)
	fi.Close()
	if err != nil {
		log.Infof("[ResetPatchUpdateScriptWithLocalFile] %v, err: %v", typ, err)
		return
	}
	_templatesPatchupdateSh = patchupdate_content
}

//it will not return a error,when a error occurs,just use the default template
func ResetSelfBuildTemplateWithLocalFile() {
	for _, file := range []string{
		"selfbuild-component.yml",
		"selfbuild-component-v1beta1.yml",
	} {
		content := readLocalFile(file)
		if content == nil {
			continue
		}
		switch file {
		case "selfbuild-component.yml":
			_templatesSelfbuildComponentYml = content
		case "selfbuild-component-v1beta1.yml":
			_templatesSelfbuildComponentV1beta1Yml = content
		}
	}
}

func ResetImportClusterNsTemplateWithLocalFile() {
	file := "import-ns-resouce.yml"
	content := readLocalFile(file)
	if content != nil {
		_templatesImportNsResouceYml = content
	}
}

func ResetImportClusterTemplateWithLocalFile() {
	for _, file := range []string{
		"import-cluster-resource.yml",
		"import-cluster-resource-v1beta1.yml",
	} {
		content := readLocalFile(file)
		if content == nil {
			continue
		}
		switch file {
		case "import-cluster-resource.yml":
			_templatesImportClusterResourceYml = content
		case "selfbuild-component-v1beta1.yml":
			_templatesImportClusterResourceV1beta1Yml = content
		}
	}
}

func readLocalFile(filename string) []byte {
	_, err := os.Stat(filename)
	if err != nil {
		return nil
	}
	fi, err := os.Open(filename)
	defer fi.Close()
	if err != nil {
		log.Infof("[local-set]: open file %s, err: %v", filename, err)
		return nil
	}
	content, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Infof("[local-set] read file %s, err: %v", filename, err)
		return nil
	}
	return content
}
