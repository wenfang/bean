package counter

import (
	"net/http"

	"nkwangwenfang.com/kit/expvar"
)

func Counter(name string, h http.Handler) http.Handler {
	c := expvar.NewCounter(name)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.Add(1)
		h.ServeHTTP(w, r)
	})
}
