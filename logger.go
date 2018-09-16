package logger

import (
	"net"
	"strings"
	"time"

	"github.com/aofei/air"
)

// GasConfig is a set of configurations for the `Gas()`.
type GasConfig struct {
	Message string
}

// Gas returns an `air.Gas` that is used to log ervery request based on the gc.
func Gas(gc GasConfig) air.Gas {
	if gc.Message == "" {
		gc.Message = "finished request-response cycle"
	}

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			startTime := time.Now()
			err := next(req, res)
			endTime := time.Now()

			clientIP, _, _ := net.SplitHostPort(req.RemoteAddr)
			if xff := req.Headers["X-Forwarded-For"]; xff != "" {
				clientIP = strings.Split(xff, ", ")[0]
			} else if xrIP := req.Headers["X-Real-IP"]; xrIP != "" {
				clientIP = xrIP
			}

			extras := map[string]interface{}{
				"remote_addr": req.RemoteAddr,
				"client_ip":   clientIP,
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
				air.ERROR(gc.Message, extras)
			} else {
				extras["status_code"] = res.StatusCode
				air.INFO(gc.Message, extras)
			}

			return err
		}
	}
}
