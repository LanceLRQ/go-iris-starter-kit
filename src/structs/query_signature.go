package structs


type QuerySignatureParams struct {
	// 签名
	Signature string `json:"signature"`
	// 时间戳
	TimeStamp string `json:"timestamp"`
	// 随机字符串
	Nonce string `json:"nonce"`
}
