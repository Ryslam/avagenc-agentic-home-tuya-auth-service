package services

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/avagenc/agentic-tuya-sign-service/internal/models"
)

func GetSign(accessID string, accessSecret string, req models.SignRequest, accessToken string) (*models.SignResponse, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	hash := sha256.New()
	hash.Write([]byte(req.Body))
	contentSha256 := hex.EncodeToString(hash.Sum(nil))

	stringToSign := req.Method + "\n" + contentSha256 + "\n\n" + req.URLPath

	nonceBytes := make([]byte, 16)
	if _, err := rand.Read(nonceBytes); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}
	nonce := hex.EncodeToString(nonceBytes)

	tuyaStr := accessID + accessToken + timestamp + nonce + stringToSign

	mac := hmac.New(sha256.New, []byte(accessSecret))
	mac.Write([]byte(tuyaStr))
	sign := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))

	return &models.SignResponse{
		Sign:       sign,
		Timestamp:  timestamp,
		Nonce:      nonce,
		SignMethod: "HMAC-SHA256",
		AccessToken: accessToken,
	},
 nil
}
