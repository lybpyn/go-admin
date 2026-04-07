package filesign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

var ErrInvalidPath = errors.New("invalid path")
var ErrInvalidExpireSeconds = errors.New("invalid expire seconds")

func GenerateSignedURL(baseURL, relativePath, secret string, expiresIn time.Duration) (string, error) {
	normalizedPath, err := NormalizeRelativePath(relativePath)
	if err != nil {
		return "", err
	}

	expires := time.Now().Add(expiresIn).Unix()
	signature := generateSignature(normalizedPath, expires, secret)

	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	u.Path = "/static"
	query := u.Query()
	query.Set("path", normalizedPath)
	query.Set("expires", strconv.FormatInt(expires, 10))
	query.Set("signature", signature)
	u.RawQuery = query.Encode()

	return u.String(), nil
}

// GenerateTemporaryAccessURL converts historical file_url/file_path into a temporary signed URL.
func GenerateTemporaryAccessURL(baseURL, rawPathOrURL, secret string, expireSeconds int64) (string, error) {
	if expireSeconds <= 0 {
		return "", ErrInvalidExpireSeconds
	}

	relativePath, ok := ExtractRelativePath(rawPathOrURL)
	if !ok {
		return "", ErrInvalidPath
	}

	return GenerateSignedURL(baseURL, relativePath, secret, time.Duration(expireSeconds)*time.Second)
}

func ValidateSignedURL(relativePath, expiresStr, signature, secret string) bool {
	if expiresStr == "" || signature == "" {
		return false
	}

	normalizedPath, err := NormalizeRelativePath(relativePath)
	if err != nil {
		return false
	}

	expires, err := strconv.ParseInt(expiresStr, 10, 64)
	if err != nil || time.Now().Unix() > expires {
		return false
	}

	expectedHex := generateSignature(normalizedPath, expires, secret)
	expectedRaw, err := hex.DecodeString(expectedHex)
	if err != nil {
		return false
	}

	providedRaw, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	return hmac.Equal(expectedRaw, providedRaw)
}

func NormalizeRelativePath(p string) (string, error) {
	p = strings.TrimSpace(strings.ReplaceAll(p, "\\", "/"))
	if p == "" || strings.Contains(p, "\x00") {
		return "", ErrInvalidPath
	}

	cleaned := path.Clean("/" + p)
	cleaned = strings.TrimPrefix(cleaned, "/")
	if cleaned == "" || cleaned == "." || cleaned == ".." {
		return "", ErrInvalidPath
	}

	return cleaned, nil
}

func ExtractRelativePath(raw string) (string, bool) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", false
	}

	if strings.HasPrefix(raw, "//") {
		return extractFromAbsoluteURL("http:" + raw)
	}

	if strings.Contains(raw, "://") {
		return extractFromAbsoluteURL(raw)
	}

	if looksLikeHostPath(raw) {
		return extractFromAbsoluteURL("http://" + raw)
	}

	if strings.Contains(raw, "?") {
		u, err := url.Parse(raw)
		if err == nil {
			if queryPath := u.Query().Get("path"); queryPath != "" {
				normalized, normalizeErr := NormalizeRelativePath(queryPath)
				return normalized, normalizeErr == nil
			}
			return extractFromPath(u.Path)
		}
	}

	return extractFromPath(raw)
}

func extractFromAbsoluteURL(raw string) (string, bool) {
	u, err := url.Parse(raw)
	if err != nil {
		return "", false
	}

	if queryPath := u.Query().Get("path"); queryPath != "" {
		normalized, normalizeErr := NormalizeRelativePath(queryPath)
		return normalized, normalizeErr == nil
	}

	return extractFromPath(u.Path)
}

func looksLikeHostPath(raw string) bool {
	if raw == "" || strings.HasPrefix(raw, "/") {
		return false
	}

	head := raw
	if idx := strings.IndexAny(raw, "/?"); idx >= 0 {
		head = raw[:idx]
	}

	return strings.Contains(head, ".") || strings.Contains(head, ":")
}

func extractFromPath(rawPath string) (string, bool) {
	pathValue := strings.TrimSpace(strings.ReplaceAll(rawPath, "\\", "/"))
	if pathValue == "" {
		return "", false
	}

	if strings.HasPrefix(pathValue, "/static/") {
		pathValue = strings.TrimPrefix(pathValue, "/static/")
	} else if strings.HasPrefix(pathValue, "static/") {
		pathValue = strings.TrimPrefix(pathValue, "static/")
	} else if strings.HasPrefix(pathValue, "/") {
		pathValue = strings.TrimPrefix(pathValue, "/")
	}

	normalized, err := NormalizeRelativePath(pathValue)
	if err != nil {
		return "", false
	}

	return normalized, true
}

func generateSignature(relativePath string, expires int64, secret string) string {
	message := relativePath + "?expires=" + strconv.FormatInt(expires, 10)

	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}
