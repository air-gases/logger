package logger

import (
	"time"

	"github.com/aofei/air"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GasConfig is a set of configurations for the `Gas`.
type GasConfig struct {
	Logger               *zerolog.Logger
	Message              string
	IncludeClientAddress bool

	Skippable func(*air.Request, *air.Response) bool
}

// Gas returns an `air.Gas` that is used to log ervery request based on the gc.
func Gas(gc GasConfig) air.Gas {
	if gc.Logger == nil {
		gc.Logger = &log.Logger
	}

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) (err error) {
			if gc.Skippable != nil && gc.Skippable(req, res) {
				return next(req, res)
			}

			startTime := time.Now()
			res.Defer(func() {
				endTime := time.Now()

				logEvent := gc.Logger.Err(err).
					Str("app_name", req.Air.AppName).
					Str(
						"remote_address",
						req.RemoteAddress(),
					)
				if gc.IncludeClientAddress {
					logEvent.Str(
						"client_address",
						req.ClientAddress(),
					)
				}

				logEvent.Str("method", req.Method).
					Str("path", req.Path).
					Int64("bytes_in", req.ContentLength).
					Int64("bytes_out", res.ContentLength).
					Int("status", res.Status).
					Time("start_time", startTime).
					Time("end_time", endTime).
					Dur("latency", endTime.Sub(startTime)).
					Msg(gc.Message)
			})

			return next(req, res)
		}
	}
}
