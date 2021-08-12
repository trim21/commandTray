package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"time"

	"github.com/utahta/go-cronowriter"
	"gopkg.in/yaml.v3"
)

var tz = time.FixedZone("Asia/Shanghai", 3600*8)

type Command struct {
	Name    string   `yaml:"name"`
	Cwd     string   `yaml:"cwd"`
	Program string   `yaml:"program"`
	Args    []string `yaml:"args"`
}

func (c Command) stdoutWriter() io.Writer {
	f := cronowriter.MustNew("./cfg/log/"+c.Name+"/out.%Y-%m-%d.log", cronowriter.WithLocation(tz))
	return f
}

func (c Command) stderrWriter() io.Writer {
	f := cronowriter.MustNew("./cfg/log/"+c.Name+"/err.%Y-%m-%d.log", cronowriter.WithLocation(tz))
	return f
}

func (c Command) job() *exec.Cmd {
	cmd := exec.Command(c.Program, c.Args...)
	if c.Cwd != "" {
		cmd.Dir = NormalizePath(c.Cwd)
	}
	return cmd
}

func (c Command) execute() {
	cmd := c.job()
	cmd.Stdout = c.stdoutWriter()
	cmd.Stderr = c.stderrWriter()
	err := cmd.Run()
	if err != nil {
		fmt.Println("err in executing " + c.Name)
	}
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
	j, err := yaml.Marshal(t)
	check(err)
	fmt.Println(string(j))
	return t
}
