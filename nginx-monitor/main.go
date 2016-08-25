package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dean123363/service-monitor/nginx-monitor/cron"
	"github.com/dean123363/service-monitor/nginx-monitor/funcs"
	"github.com/dean123363/service-monitor/nginx-monitor/g"
	"github.com/dean123363/service-monitor/nginx-monitor/http"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	g.InitRootDir()
	g.InitRpcClients()

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	funcs.BuildMappers()

	cron.Collect()

	go http.Start()

	select {}

}
