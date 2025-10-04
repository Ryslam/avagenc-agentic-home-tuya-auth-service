package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/avagenc/agentic-tuya-sign-service/internal/models"
)

const tuyaTokenEndpointPath = "/v1.0/token?grant_type=1"

func GetAccessToken(accessID string, baseURL string, signature *models.SignResponse) (string, error) {
	fullURLPath := baseURL + tuyaTokenEndpointPath

	httpReq, err := http.NewRequest("GET", fullURLPath, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request to %s: %w", fullURLPath, err)
	}

	httpReq.Header.Set("client_id", accessID)
	httpReq.Header.Set("sign", signature.Sign)
	httpReq.Header.Set("t", signature.Timestamp)
	httpReq.Header.Set("nonce", signature.Nonce)
	httpReq.Header.Set("sign_method", signature.SignMethod)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", fmt.Errorf("failed to send request to %s: %w", fullURLPath, err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("tuya api call error: status code %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	var tuyaTokenResponse models.TuyaTokenResponse
	if err := json.Unmarshal(bodyBytes, &tuyaTokenResponse); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if !tuyaTokenResponse.Success {
		return "", fmt.Errorf("tuya api call fail: %v", tuyaTokenResponse.Result)
	}

	return tuyaTokenResponse.Result.AccessToken, nil
}
