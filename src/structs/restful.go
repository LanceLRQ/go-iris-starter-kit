package structs

type RESTfulAPIResult struct {
	Status bool `json:"status"`
	ErrCode int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
