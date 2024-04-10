package flag

import (
	"billing_finops/pkg/conf"
	"flag"
)

var (
	cfg  = flag.String("f", "config.yaml", "config file path")
	ADDR = flag.String("addr", ":8080", "server address")
)

func init() {
	flag.Parse()
	conf.Init(*cfg)
}
