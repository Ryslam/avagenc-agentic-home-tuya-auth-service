package models

type SignRequest struct {
	Method  string `json:"method"`
	URLPath string `json:"url_path"`
	Body    string `json:"body"`
}

type SignResponse struct {
	Sign        string `json:"sign"`
	Timestamp   string `json:"t"`
	Nonce       string `json:"nonce"`
	SignMethod  string `json:"sign_method"`
	AccessToken string `json:"access_token"`
}
