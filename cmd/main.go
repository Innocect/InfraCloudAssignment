package main

import (
	"shortlink/internal/server"

	"github.com/golang/glog"
)

func main() {
	err := server.RunServer()
	if err != nil {
		glog.Fatal("Server Error" + err.Error())
	}
}
