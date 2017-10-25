package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/fzerorubigd/expand"
	onion "gopkg.in/fzerorubigd/onion.v3"
	"gopkg.in/fzerorubigd/onion.v3/extraenv"
	_ "gopkg.in/fzerorubigd/onion.v3/yamlloader"
)

var (
	o = onion.New()

	debug     = o.RegisterBool("debug", false)
	namespace = o.RegisterString("namespace", "default")
	service   = o.RegisterString("service", "netdatasvc")
	domain    = o.RegisterString("domain", "")
)

//initialize try to initialize config
func initialize(prefix string) {
	// Now load external config to overwrite them all.
	if err := o.AddLayer(onion.NewFileLayer("/etc/netdata-proxy/config.yaml")); err == nil {
		logrus.Infof("loading config from %s", "/etc/netdata-proxy/config.yaml")
	}
	p, err := expand.Path("$HOME/.netdata-proxy/config.yaml")
	if err == nil {
		if err = o.AddLayer(onion.NewFileLayer(p)); err == nil {
			logrus.Infof("loading config from %s", p)
		}
	}

	p, err = expand.Path("$PWD/config.yaml")
	if err == nil {
		if err = o.AddLayer(onion.NewFileLayer(p)); err == nil {
			logrus.Infof("loading config from %s", p)
		}
	}

	o.AddLazyLayer(extraenv.NewExtraEnvLayer(prefix))

	// load all registered variables
	o.Load()
	setConfigParameter()
}

// setConfigParameter try to set the config parameter for the logrus base on config
func setConfigParameter() {
	if debug.Bool() {
		// In development mode I need colors :) candy mode is GREAT!
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true, DisableColors: false})
		logrus.SetLevel(logrus.DebugLevel)

	} else {
		logrus.SetFormatter(&logrus.TextFormatter{ForceColors: false, DisableColors: true})
		logrus.SetLevel(logrus.WarnLevel)
	}
}
