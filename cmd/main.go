package main

import (
	"billing_finops/pkg/flag"
	_ "billing_finops/pkg/flag"
	"billing_finops/pkg/sdk"
	"github.com/prometheus/client_golang/prometheus"

	"billing_finops/utils"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	fmt.Println("___.   .__.__  .__  .__                    _____.__                            \n\\_ |__ |__|  | |  | |__| ____    ____    _/ ____\\__| ____   ____ ______  ______\n | __ \\|  |  | |  | |  |/    \\  / ___\\   \\   __\\|  |/    \\ /  _ \\\\____ \\/  ___/\n | \\_\\ \\  |  |_|  |_|  |   |  \\/ /_/  >   |  |  |  |   |  (  <_> )  |_> >___ \\ \n |___  /__|____/____/__|___|  /\\___  /____|__|  |__|___|  /\\____/|   __/____  >\n     \\/                     \\//_____/_____/             \\/       |__|       \\/")

	var CurrentBalance = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "current_balance",
		Help: "current balance of the account",
	})
	prometheus.MustRegister(CurrentBalance)
	balance, err := sdk.AliyunSDK.GetBalance()
	if err != nil {
		balance = 0
	}
	CurrentBalance.Set(balance)
	http.Handle("/metrics", promhttp.Handler())
	utils.Log.Infof("server start at %s", *flag.ADDR)
	err = http.ListenAndServe(*flag.ADDR, nil)
	if err != nil {
		utils.Log.Panic("can not start server:", err)
	}

}
