package models

type TuyaTokenResponseResult struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireTime   int    `json:"expire_time"`
	UID          string `json:"uid"`
}

type TuyaTokenResponse struct {
	Result    TuyaTokenResponseResult `json:"result"`
	Success   bool                    `json:"success"`
	Timestamp int64                   `json:"t"`
	Tid       string                  `json:"tid"`
}
