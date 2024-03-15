package middlewares

import (
	"github.com/dunkbing/go-common/mux"
	"github.com/dunkbing/go-common/utils"
	"net/http"
)

func EnableCORS(origins []string) mux.Middleware {
	allowedOriginsMap := map[string]bool{}
	for _, v := range origins {
		url, _ := utils.CleanURL(v)
		allowedOriginsMap[url] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if allowedOriginsMap[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
			next.ServeHTTP(w, r)
		})
	}
}
