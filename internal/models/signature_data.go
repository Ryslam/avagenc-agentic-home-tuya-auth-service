package models

type SignatureData struct {
	Sign       string `json:"sign"`
	Timestamp  string `json:"t"`
	Nonce      string `json:"nonce"`
	SignMethod string `json:"sign_method"`
}
