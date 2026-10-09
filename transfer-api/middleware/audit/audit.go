package audit

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type inspectingWriter struct {
	hrw        http.ResponseWriter
	statusCode int
}

var _ http.ResponseWriter = (*inspectingWriter)(nil)

func (iw *inspectingWriter) Header() http.Header {
	return iw.hrw.Header()
}

func (iw *inspectingWriter) Write(b []byte) (int, error) {
	if iw.statusCode == 0 {
		iw.WriteHeader(http.StatusOK)
	}

	return iw.hrw.Write(b)
}

func (iw *inspectingWriter) WriteHeader(statusCode int) {
	iw.statusCode = statusCode
	iw.hrw.WriteHeader(statusCode)
}

func WithAudit(l *zap.SugaredLogger, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		iw := inspectingWriter{
			hrw:        w,
			statusCode: 0,
		}

		next.ServeHTTP(&iw, r)

		totalServeTimeMs := time.Now().UnixMilli() - start.UnixMilli()

		l.Infow("request completed",
			"path", r.URL.Path,
			"method", r.Method,
			"statusCode", iw.statusCode,
			"duration", totalServeTimeMs)
	})
}
