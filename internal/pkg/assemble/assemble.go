package assemble

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"time"
)

const (
	apiKey    = "38e8643fb0dc04e8d65b99994d3dafff"
	// secretKey = "10a01dcf33762d3a204cb96429918ff6"
	secretKey = "10a01dcf33762d3a204cb96429918ff6"
	zingMp3BaseURL = "https://zingmp3.vn/api"
)

// ZingMp3SongAPIs will return list of apis 
func ZingMp3SongAPIs(codes []string, nameURL string) []string {
	var apis []string
	for _, code := range codes {
		now := time.Now()
		timestamp := now.Unix()
		apiURL := zingMp3BaseURL + nameURL;
		idAndCtimeParams := fmt.Sprintf("id=%s&ctime=%d", code, timestamp)

		h256 := sha256.New()
		h256.Write([]byte(fmt.Sprintf("ctime=%did=%s", timestamp, code)))

		mac := hmac.New(sha512.New,  []byte(secretKey))
		mac.Write([]byte(fmt.Sprintf("%s%x", nameURL, h256.Sum(nil))))
		zingMp3SigParam := mac.Sum(nil)

		apis = append(apis, fmt.Sprintf("%s?%s&sig=%x&api_key=%s", apiURL, idAndCtimeParams, zingMp3SigParam, apiKey))
	}

	return apis
}