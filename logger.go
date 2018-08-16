package logger

import (
	"time"

	"github.com/sheng/air"
)

// msg is the logger's main output message.
const msg = "finished request-response cycle"

// Gas is used to log every request.
func Gas(next air.Handler) air.Handler {
	return func(req *air.Request, res *air.Response) error {
		startTime := time.Now()
		err := next(req, res)
		endTime := time.Now()

		extras := map[string]interface{}{
			"remote_addr": req.RemoteAddr,
			"method":      req.Method,
			"path":        req.URL.Path,
			"start_time":  startTime.UnixNano(),
			"end_time":    endTime.UnixNano(),
			"latency":     endTime.Sub(startTime).String(),
			"bytes_in":    req.ContentLength,
			"bytes_out":   res.ContentLength,
		}

		if err != nil {
			if e, ok := err.(*air.Error); ok {
				extras["status_code"] = e.Code
			} else {
				extras["status_code"] = 500
			}

			extras["error"] = err.Error()
			air.ERROR(msg, extras)
		} else {
			extras["status_code"] = res.StatusCode
			air.INFO(msg, extras)
		}

		return err
	}
}
