package logger

import (
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

			extras := map[string]interface{}{
				"method":         req.Method,
				"status":         res.Status,
				"path":           req.Path,
				"remote_address": req.RemoteAddress,
				"client_ip":      req.ClientIP,
				"start_time":     startTime.UnixNano(),
				"end_time":       endTime.UnixNano(),
				"latency":        endTime.Sub(startTime),
				"bytes_in":       req.ContentLength,
				"bytes_out":      res.ContentLength,
			}

			if err != nil {
				extras["error"] = err.Error()
				air.ERROR(gc.Message, extras)
			} else {
				air.INFO(gc.Message, extras)
			}

			return err
		}
	}
}
