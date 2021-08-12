// Copyright 2011 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
)

func main() {
	cfg := load()
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
