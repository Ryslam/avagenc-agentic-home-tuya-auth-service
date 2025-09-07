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

	"github.com/Ryslam/avagenc-agentic-home-tuya-auth-service/internal/models"
)

// GetSignature generates the signature required for Tuya API requests.
// GetSignatue is dynamic, can be used for any tuya endpoint
func GetSignature(accessID, accessSecret, method, urlPath, body, accessToken string) (*models.SignatureData, error) {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)

	hash := sha256.New()
	hash.Write([]byte(body))
	contentSha256 := hex.EncodeToString(hash.Sum(nil))

	stringToSign := method + "\n" + contentSha256 + "\n\n" + urlPath

	nonceBytes := make([]byte, 16)
	if _, err := rand.Read(nonceBytes); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}
	nonce := hex.EncodeToString(nonceBytes)

	tuyaStr := accessID + accessToken + timestamp + nonce + stringToSign

	mac := hmac.New(sha256.New, []byte(accessSecret))
	mac.Write([]byte(tuyaStr))
	sign := strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))

	return &models.SignatureData{
		Sign:       sign,
		Timestamp:  timestamp,
		Nonce:      nonce,
		SignMethod: "HMAC-SHA256",
	},
}
