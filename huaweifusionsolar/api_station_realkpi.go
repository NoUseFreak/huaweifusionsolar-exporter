package huaweifusionsolar

import "encoding/json"

type APIStationRealKpiResponse struct {
	ApiResponse
	Data []struct {
		DataItemMap APIStationRealKpiDataItemMap `json:"dataItemMap"`
		StationCode string                       `json:"stationCode"`
	}
}

type APIStationRealKpiDataItemMap struct {
	DayPower        json.Number `json:"day_power"`
	MonthPower      json.Number `json:"month_power"`
	TotalPower      json.Number `json:"total_power"`
	DayIncome       json.Number `json:"day_income"`
	TotalIncome     json.Number `json:"total_income"`
	RealHealthState json.Number `json:"real_health_state"`
}

type stationRealKpiRequest struct {
	StationCodes string `json:"stationCodes"`
}

func (h *HuaweiInverter) GetStationRealKpi(stationCodes string) (*APIStationRealKpiResponse, error) {
	resp := APIStationRealKpiResponse{}
	if err := h.queryURL("/getStationRealKpi", stationRealKpiRequest{
		StationCodes: stationCodes,
	}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
