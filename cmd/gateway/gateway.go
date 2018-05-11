package main

import (
	"github.com/oikomi/FishChatServer3/cmd/gateway/app/options"
	"math/rand"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/golang/glog"
)

func main() {

	glog.Info()
	rand.Seed(time.Now().UTC().UnixNano())

	config := options.NewGatewayFlags()
	config.AddFlags(flag.CommandLine)

	flag.Parse()

}
