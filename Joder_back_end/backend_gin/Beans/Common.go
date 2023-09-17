package beans

type ResponseWrapper struct {
	ResultCode    string         `json:"resultCode"`
	ResultMessage string         `json:"resultMessage"`
	Model         map[string]any `json:"model"`
}
