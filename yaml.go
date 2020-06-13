package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Setting is the structure to configure watch repositories
type Setting struct {
	Github struct {
		URL          string `yaml:"url"`
		UserName     string `yaml:"userName"`
		Token        string `yaml:"token"`
		Repositories []struct {
			Owner    string `yaml:"owner"`
			RepoName string `yaml:"repoName"`
		} `yaml:"repositories"`
	} `yaml:"github"`
}

func parseSetting(settingPath string) Setting {
	var setting Setting
	os.Stat(settingPath)
	fileInfo, err := os.Stat(settingPath)
	if os.IsNotExist(err) || fileInfo.IsDir() {
		return Setting{}
	}
	data, err := ioutil.ReadFile(settingPath)
	if err != nil {
		log.Printf("read setting file(%s) error(%s)", settingPath, err)
		return Setting{}
	}
	err = yaml.Unmarshal(data, &setting)
	if err != nil {
		log.Printf("unmarshal setting file(%s) failed(%s)", settingPath, err)
		return Setting{}
	}
	return setting
}
