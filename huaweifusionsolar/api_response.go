package huaweifusionsolar

type ApiResponse struct {
	BuildCode string      `json:"buildCode"`
	FailCode  int64       `json:"failCode"`
	Message   string      `json:"message"`
	Params    interface{} `json:"params"`
	Success   bool        `json:"success"`
}
