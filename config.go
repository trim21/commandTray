package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

type Command struct {
	Name    string   `yaml:"name"`
	Cwd     string   `yaml:"cwd"`
	Program string   `yaml:"program"`
	Args    []string `yaml:"args"`
}

func (c Command) job() *exec.Cmd {
	cmd := exec.Command(NormalizePath(c.Program), c.Args...)
	if c.Cwd != "" {
		cmd.Dir = NormalizePath(c.Cwd)
	}
	return cmd
}

func (c Command) execute() {
	cmd := c.job()
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		logrus.Errorln(err)
		fmt.Println(stderr.String())
	}
	fmt.Println(out.String())
}

func (c Command) check() error {
	if c.Program == "" {
		return fmt.Errorf("must set command")
	}
	return nil
}

type C struct {
	Command `yaml:",inline"`
	Spec    string `yaml:"spec"`
}

type T struct {
	A       string
	Cron    []C       `yaml:"cron"`
	Program []Command `yaml:"program"`
}

func load() T {
	var t = T{}
	data, err := ioutil.ReadFile("./config.yaml")
	check(err)
	check(yaml.Unmarshal(data, &t))
	j, err := json.Marshal(t)
	check(err)
	fmt.Println(string(j))
	return t
}
