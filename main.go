package main

import (
	"miaomiao/miaomiao"
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("err recovered: %+v", err)
			log.Errorf("%s", debug.Stack())
		}
	}()
	miaomiao.Start()
}
