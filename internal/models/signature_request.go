package models

type SignatureRequest struct {
	AccessID     string `json:"access_id"`
	AccessSecret string `json:"access_secret"`
	Method       string `json:"method"`
	URLPath      string `json:"url_path"`
	Body         string `json:"body"`
	AccessToken  string `json:"access_token"`
}
