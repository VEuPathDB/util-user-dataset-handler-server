package middle

import (
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
	"github.com/VEuPathDB/util-exporter-server/internal/service/rid"
	"net/http"
)

func RequestCtxProvider(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		id, err := rid.GenerateRID()
		if err != nil {
			log.Logger().WithField("endpoint", req.URL.Path).Error(err)
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = res.Write([]byte(fmt.Sprintf(simpleErrFmt, err)))
			return
		}

		req.Header[rid.RIDKey] = []string{string(id)}
		logger.AddFields(id, map[string]interface{}{
			"endpoint": req.URL.Path,
			"method": req.Method,
		})

		next.ServeHTTP(res, req)
	}
}
