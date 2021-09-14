package filewatcher

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type config struct {
	DiaryDir string `yaml:"diary-dir"`
}

func readConf(filename string) (*config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	conf := &config{}
	err = yaml.Unmarshal(buf, conf)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}

	return conf, nil
}

func Watcher() {
	// This is the correct file structure: ./config/dev/config.yml
	config, err := readConf("./config/dev/config.yml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", config)
}
