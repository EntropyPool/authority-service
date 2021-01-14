package main

import (
	"fmt"
	"github.com/EntropyPool/entropy-logger"
	"github.com/NpoolRD/http-daemon"
	"github.com/NpoolRD/service-register"
	"net/http"
)

func serveReportApi(w http.ResponseWriter, req *http.Request) (interface{}, error, int) {
	return nil, nil, 0
}

func serveLogin(w http.ResponseWriter, req *http.Request) (interface{}, error, int) {
	return nil, nil, 0
}

func serveQuery(w http.ResponseWriter, req *http.Request) (interface{}, error, int) {
	return nil, nil, 0
}

func main() {
	myVal := fmt.Sprintf("%v:%v", MyConfig.IPAddr, MyConfig.Port)
	err := srvreg.Register(MyConfig.Domain, myVal)
	if nil != err {
		elog.Fatalf(elog.Fields{}, "cannot register service register")
	}

	httpdaemon.RegisterRouter(httpdaemon.HttpRouter{
		Location: "/auth/report_api",
		Handler:  serveReportApi,
	})
	httpdaemon.RegisterRouter(httpdaemon.HttpRouter{
		Location: "/auth/login",
		Handler:  serveLogin,
	})
	httpdaemon.RegisterRouter(httpdaemon.HttpRouter{
		Location: "/auth/query",
		Handler:  serveQuery,
	})

	httpdaemon.Run(MyConfig.Port)
	ch := make(chan int, 0)
	<-ch
}
