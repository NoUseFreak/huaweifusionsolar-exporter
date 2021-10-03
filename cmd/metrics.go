package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	fusionsolar "github.com/NoUseFreak/huaweifusionsolar/huaweifusionsolar"
)

var (
	labels      = []string{"station"}
	dayPower    = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "huawei_solarfusion_day_power"}, labels)
	monthPower  = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "huawei_solarfusion_month_power"}, labels)
	totalPower  = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "huawei_solarfusion_total_power"}, labels)
	dayIncome   = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "huawei_solarfusion_day_income"}, labels)
	totalIncome = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "huawei_solarfusion_total_income"}, labels)
)

func init() {
	rootCmd.AddCommand(metricCmd)
}

var metricCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Expose metrics on the API",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Exposing metrics http://0.0.0.0:2112/metrics")

		recordMetrics()

		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	},
}

func recordMetrics() {
	invertor, _ := fusionsolar.New(viper.GetString("username"), viper.GetString("password"), viper.GetString("api_endpoint"))
	invertor.Login()
	stationList, _ := invertor.GetStationList()
	stationCode := stationList.Data[0].StationCode

	go func() {
		ticker := time.NewTicker(60 * time.Second)
		for {
			select {
			case <-ticker.C:
				d, err := invertor.GetStationRealKpi(stationList.Data[0].StationCode)
				if err != nil {
					fmt.Println("request failed, retrying - " + err.Error())
					invertor.Login()
				} else {
					if val, err := d.Data[0].DataItemMap.DayPower.Float64(); err == nil {
						dayPower.With(prometheus.Labels{"station": stationCode}).Set(val)
					}
					if val, err := d.Data[0].DataItemMap.MonthPower.Float64(); err == nil {
						monthPower.With(prometheus.Labels{"station": stationCode}).Set(val)
					}
					if val, err := d.Data[0].DataItemMap.TotalPower.Float64(); err == nil {
						totalPower.With(prometheus.Labels{"station": stationCode}).Set(val)
					}
					if val, err := d.Data[0].DataItemMap.DayIncome.Float64(); err == nil {
						dayIncome.With(prometheus.Labels{"station": stationCode}).Set(val)
					}
					if val, err := d.Data[0].DataItemMap.TotalIncome.Float64(); err == nil {
						totalIncome.With(prometheus.Labels{"station": stationCode}).Set(val)
					}
				}
			}
		}
	}()
}
