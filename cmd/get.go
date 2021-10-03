package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	fusionsolar "github.com/NoUseFreak/huaweifusionsolar/huaweifusionsolar"
	"github.com/araddon/dateparse"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	getCmd.PersistentFlags().StringP("time", "t", "", "Time to get data for")
	getCmd.PersistentFlags().StringP("name", "n", "", "Name")

	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get TYPE",
	Short: "Get metrics on the CLI",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		invertor, _ := fusionsolar.New(viper.GetString("username"), viper.GetString("password"), viper.GetString("api_endpoint"))
		if err := invertor.Login(); err != nil {
			logrus.Fatal(err)
		}

		stationList, err := invertor.GetStationList()
		if err != nil {
			logrus.Fatal(err)
		}

		dt := time.Now()
		if timeString, _ := cmd.Flags().GetString("time"); timeString != "" {
			if dt, err = dateparse.ParseLocal(timeString); err != nil {
				logrus.Fatal("Invalid datetime format")
			}
		}

		switch args[0] {
		case "DevList":
			d, _ := invertor.GetDevList(stationList.Data[0].StationCode)
			printJSON(d)
			break
		case "KpiStation":
			name, _ := cmd.Flags().GetString("name")
			d, _ := invertor.GetKpiStation(name, stationList.Data[0].StationCode, dt)
			printJSON(d)
			break
		case "KpiStationHour":
			d, _ := invertor.GetKpiStationHour(stationList.Data[0].StationCode, dt)
			printJSON(d)
			break
		case "KpiStationDay":
			d, _ := invertor.GetKpiStationDay(stationList.Data[0].StationCode, dt)
			printJSON(d)
			break
		case "KpiStationMonth":
			d, _ := invertor.GetKpiStationMonth(stationList.Data[0].StationCode, dt)
			printJSON(d)
			break
		case "KpiStationYear":
			d, _ := invertor.GetKpiStationYear(stationList.Data[0].StationCode, dt)
			printJSON(d)
			break
		case "StationList":
			d, _ := invertor.GetStationList()
			printJSON(d)
			break
		case "StationRealKpi":
			d, _ := invertor.GetStationRealKpi(stationList.Data[0].StationCode)
			printJSON(d)
			break
		default:
			logrus.Fatal("Unkown Type")
		}
	},
}

func printJSON(data interface{}) {
	js, _ := json.Marshal(data)
	fmt.Printf("%+v\n", string(js))
}
