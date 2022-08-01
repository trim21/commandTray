// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
)

func main() {
	cfg := load()
	for index, command := range cfg.Program {
		if command.Name == "" {
			fmt.Printf("program %d need a name \n", index)
			return
		}
		if command.Program == "" {
			fmt.Printf("program %d need a executable \n", index)
			return
		}
	}

	for index, command := range cfg.Cron {
		if command.Name == "" {
			fmt.Printf("cron %d need a name \n", index)
		}
		if command.Program == "" {
			fmt.Printf("cron %d need a executable \n", index)
			return
		}
	}

	for _, command := range cfg.Program {
		go command.execute()
	}
	c := buildCron(cfg)

	c.Run()
	return
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
