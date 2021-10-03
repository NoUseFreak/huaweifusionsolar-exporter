package huaweifusionsolar

import "time"

type APIKpiStationMetricResponse struct {
	ApiResponse
	Data []struct {
		CollectTime int64                          `json:"collectTime"`
		DataItemMap APIKpiStationMetricDataItemMap `json:"dataItemMap"`
		StationCode string                         `json:"stationCode"`
	}
}

type APIKpiStationMetricDataItemMap struct {
	RadiationIntensity interface{} `json:"radiation_intensity"`
	TheoryPower        interface{} `json:"theory_power"`
	InverterPower      float64     `json:"inverter_power"`
	OngridPower        interface{} `json:"ongrid_power"`
	PowerProfit        interface{} `json:"power_profit"`
	InstalledCapacity  float64     `json:"installed_capacity"`
	PerformanceRatio   interface{} `json:"performance_ratio"`
	UsePower           interface{} `json:"use_power"`
	PerpowerRatio      float64     `json:"perpower_ratio"`
	ReductionTotalCo2  float64     `json:"reduction_total_co2"`
	ReductionTotalCoal float64     `json:"reduction_total_coal"`
	ReductionTotalTree interface{} `json:"reduction_total_tree"`
}

type kpiStationRequest struct {
	StationCodes string `json:"stationCodes"`
	CollectTime  int64  `json:"collectTime"`
}

func (h *HuaweiInverter) GetKpiStationHour(stationCodes string, collectTime time.Time) (*APIKpiStationMetricResponse, error) {
	return h.GetKpiStation("Hour", stationCodes, collectTime)
}

func (h *HuaweiInverter) GetKpiStationDay(stationCodes string, collectTime time.Time) (*APIKpiStationMetricResponse, error) {
	return h.GetKpiStation("Day", stationCodes, collectTime)
}

func (h *HuaweiInverter) GetKpiStationMonth(stationCodes string, collectTime time.Time) (*APIKpiStationMetricResponse, error) {
	return h.GetKpiStation("Month", stationCodes, collectTime)
}

func (h *HuaweiInverter) GetKpiStationYear(stationCodes string, collectTime time.Time) (*APIKpiStationMetricResponse, error) {
	return h.GetKpiStation("Year", stationCodes, collectTime)
}

func (h *HuaweiInverter) GetKpiStation(name string, stationCodes string, collectTime time.Time) (*APIKpiStationMetricResponse, error) {
	resp := APIKpiStationMetricResponse{}
	if err := h.queryURL("/getKpiStation"+name, kpiStationRequest{
		StationCodes: stationCodes,
		CollectTime:  collectTime.UnixNano() / int64(time.Millisecond),
	}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
