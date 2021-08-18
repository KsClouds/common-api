package main

import (
	"common/crawler/dujitang"
	"common/crawler/weather"
	"flag"
	"log"
	"time"

	metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
)

var (
	zkAddr   = flag.String("zkAddr", "111.230.25.75:2181", "zookeeper address")
	basePath = flag.String("base", "/common_api", "prefix path")
)

// go run -tags etcd server.go
func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)
	s.Register(new(weather.Weather), "")
	s.Register(new(dujitang.Dujitang), "")
	s.Serve("tcp", ":9001")
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@:9001",
		ZooKeeperServers: []string{*zkAddr},
		BasePath:         *basePath,
		Metrics:          metrics.NewRegistry(),
		UpdateInterval:   time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}
	s.Plugins.Add(r)
}
