package sign

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
	"time"

	"nkwangwenfang.com/kit/transport/http/handler"
)

const timeLimit = 600

var keys = make(map[string]string)

// SetAppKey 设置app key
func SetAppKey(app, key string) {
	keys[app] = key
}

// Sign 为Handler提供签名功能
func Sign(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientSign := r.Header.Get("X-Signature")
		if clientSign == "" {
			handler.OutputError(w, "Header X-Signature not set", handler.ErrorTypeAuth)
			return
		}

		urlValues, err := handler.ParseURL(r)
		if err != nil {
			handler.OutputError(w, "parse url error", handler.ErrorTypeAuth)
			return
		}

		appKey := keys[urlValues.Get("app")]
		if appKey == "" {
			handler.OutputError(w, "app not set", handler.ErrorTypeAuth)
			return
		}

		timestamp, err := strconv.ParseInt(urlValues.Get("timestamp"), 10, 64)
		if err != nil {
			handler.OutputError(w, "timestamp invalid", handler.ErrorTypeAuth)
			return
		}

		now := time.Now().Unix()
		if now-timestamp > timeLimit || now-timestamp < -timeLimit {
			handler.OutputError(w, "signature expired", handler.ErrorTypeAuth)
			return
		}

		if urlValues.Get("nonce") == "" {
			handler.OutputError(w, "nonce not set", handler.ErrorTypeAuth)
			return
		}

		var buf bytes.Buffer
		buf.WriteString(appKey)
		buf.WriteString("|request_uri=")
		buf.WriteString(r.RequestURI)

		h1 := sha1.New()
		h1.Write(buf.Bytes())
		serverSign := hex.EncodeToString(h1.Sum(nil))
		h2 := sha1.New()
		h2.Write([]byte(serverSign))
		serverSign = hex.EncodeToString(h2.Sum(nil))
		serverSign = strings.ToUpper(serverSign)

		if clientSign != serverSign {
			handler.OutputError(w, "signature error", handler.ErrorTypeAuth)
			return
		}

		h.ServeHTTP(w, r)
	})
}
