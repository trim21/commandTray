package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/utahta/go-cronowriter"
	"gopkg.in/yaml.v3"
)

var tz = time.FixedZone("Asia/Shanghai", 3600*8)

type Command struct {
	Name    string   `yaml:"name"`
	Cwd     string   `yaml:"cwd,omitempty"`
	Program string   `yaml:"program"`
	Args    []string `yaml:"args,omitempty"`
	Env     []string `yaml:"env,omitempty"`
}

func (c Command) stdoutWriter() io.Writer {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return cronowriter.MustNew(filepath.Join(homedir, "log/"+c.Name+"/out.%Y-%m-%d.log"), cronowriter.WithLocation(tz))
}

func (c Command) stderrWriter() io.Writer {
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return cronowriter.MustNew(filepath.Join(homedir, "log/"+c.Name+"/err.%Y-%m-%d.log"), cronowriter.WithLocation(tz))
}

func (c Command) job() *exec.Cmd {
	cmd := exec.Command(c.Program, c.Args...)
	if c.Cwd != "" {
		cmd.Dir = NormalizePath(c.Cwd)
	}
	cmd.Env = append(os.Environ(), c.Env...)
	return cmd
}

func (c Command) execute() {
	cmd := c.job()
	cmd.Stdout = c.stdoutWriter()
	buf := bytes.NewBuffer(nil)
	cmd.Stderr = io.MultiWriter(c.stderrWriter(), buf)
	err := cmd.Run()
	if err != nil {
		fmt.Println("err in executing " + c.Name)
		fmt.Println(buf.String())
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

func (c C) MarshalYAML() (interface{}, error) {
	return struct {
		Spec    string `yaml:"spec"`
		Name    string `yaml:"name"`
		Program string `yaml:"program"`
	}{
		Name:    c.Name,
		Spec:    c.Spec,
		Program: strings.Join(append([]string{c.Program}, c.Args...), " "),
	}, nil
}

func (c Command) MarshalYAML() (interface{}, error) {
	return struct {
		Name    string `yaml:"name"`
		Program string `yaml:"program"`
	}{
		Name:    c.Name,
		Program: strings.Join(append([]string{c.Program}, c.Args...), " "),
	}, nil
}

type T struct {
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
