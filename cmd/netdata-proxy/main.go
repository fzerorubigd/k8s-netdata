package main

import (
	"context"
	"os"
	"os/signal"

	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/fzerorubigd/k8s-netdata-proxy/peers"
	"github.com/fzerorubigd/k8s-netdata-proxy/sets"
)

func main() {
	initialize("NDP")
	ctx, cnl := context.WithCancel(context.Background())
	defer cnl()

	change := make(chan sets.String)
	go peers.Find(ctx, namespace.String(), domain.String(), service.String(), change)
	go routes(ctx, change)
	sig := make(chan os.Signal, 6)
	signal.Notify(sig, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGHUP)
	<-sig
	logrus.Debugf("Closing, buy...")
}
