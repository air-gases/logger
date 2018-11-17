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
				"remote_address": req.RemoteAddress(),
				"client_address": req.ClientAddress(),
				"method":         req.Method,
				"path":           req.Path,
				"bytes_in":       req.ContentLength,
				"bytes_out":      res.ContentLength,
				"status":         res.Status,
				"start_time":     startTime.UnixNano(),
				"end_time":       endTime.UnixNano(),
				"latency":        endTime.Sub(startTime),
			}

			if err != nil {
				extras["error"] = err.Error()
				req.Air.ERROR(gc.Message, extras)
			} else {
				req.Air.INFO(gc.Message, extras)
			}

			return err
		}
	}
}
