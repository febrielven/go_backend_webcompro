package middlewares

import (
	"net/http"

	"github.com/febrielven/go_backend_webcompro/api/utils"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.Debug(struct {
			Method  string `json:"method"`
			Url     string `json:"url"`
			Version string `json:"version"`
		}{
			Method:  r.Method,
			Url:     r.Host + r.RequestURI,
			Version: r.Proto,
		})

		next(w, r)
	}
}
