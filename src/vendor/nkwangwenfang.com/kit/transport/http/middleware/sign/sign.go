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
			handler.OutputError(w, handler.ErrorReason{Reason: "Header X-Signature not set"}, handler.ErrorAuth)
			return
		}

		appKey := keys[r.FormValue("app")]
		if appKey == "" {
			handler.OutputError(w, handler.ErrorReason{Reason: "app not set"}, handler.ErrorAuth)
			return
		}

		timestamp, err := strconv.ParseInt(r.FormValue("timestamp"), 10, 64)
		if err != nil {
			handler.OutputError(w, handler.ErrorReason{Reason: "timestamp invalid"}, handler.ErrorAuth)
			return
		}

		now := time.Now().Unix()
		if now-timestamp > timeLimit || now-timestamp < -timeLimit {
			handler.OutputError(w, handler.ErrorReason{Reason: "signature expired"}, handler.ErrorAuth)
			return
		}

		if r.FormValue("nonce") == "" {
			handler.OutputError(w, handler.ErrorReason{Reason: "nonce not set"}, handler.ErrorAuth)
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
			handler.OutputError(w, nil, handler.ErrorAuth)
			return
		}

		h.ServeHTTP(w, r)
	})
}
