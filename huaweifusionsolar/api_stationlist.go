package huaweifusionsolar

type APIStationListResponse struct {
	ApiResponse
	Data []struct {
		AidType        int64   `json:"aidType"`
		BuildState     string  `json:"buildState"`
		Capacity       float64 `json:"capacity"`
		CombineType    string  `json:"combineType"`
		LinkmanPho     string  `json:"linkmanPho"`
		StationAddr    string  `json:"stationAddr"`
		StationCode    string  `json:"stationCode"`
		StationLinkman string  `json:"stationLinkman"`
		StationName    string  `json:"stationName"`
	}
}

type stationListRequest struct{}

func (h *HuaweiInverter) GetStationList() (*APIStationListResponse, error) {
	resp := APIStationListResponse{}
	if err := h.queryURL("/getStationList", stationListRequest{}, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
