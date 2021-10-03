package huaweifusionsolar

type devListRequest struct {
	StationCodes string `json:"stationCodes"`
}

func (h *HuaweiInverter) GetDevList(stationCodes string) (*ApiResponse, error) {
	resp := ApiResponse{}
	if err := h.queryURL("/getKpiStationYear", devListRequest{
		StationCodes: stationCodes,
	}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
