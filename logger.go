package logger

import (
	"time"

	"github.com/aofei/air"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// GasConfig is a set of configurations for the `Gas()`.
type GasConfig struct {
	Logger  *zerolog.Logger
	Message string
}

// Gas returns an `air.Gas` that is used to log ervery request based on the gc.
func Gas(gc GasConfig) air.Gas {
	if gc.Logger == nil {
		gc.Logger = &log.Logger
	}

	if gc.Message == "" {
		gc.Message = "finished request-response cycle"
	}

	return func(next air.Handler) air.Handler {
		return func(req *air.Request, res *air.Response) error {
			startTime := time.Now()
			err := next(req, res)
			endTime := time.Now()

			var logEvent *zerolog.Event
			if err != nil {
				logEvent = gc.Logger.Error().Err(err)
			} else {
				logEvent = gc.Logger.Info()
			}

			logEvent.
				Str("app_name", req.Air.AppName).
				Str("remote_address", req.RemoteAddress()).
				Str("client_address", req.ClientAddress()).
				Str("method", req.Method).
				Str("path", req.Path).
				Int64("bytes_in", req.ContentLength).
				Int64("bytes_out", res.ContentLength).
				Int("status", res.Status).
				Time("start_time", startTime).
				Time("end_time", endTime).
				Dur("latency", endTime.Sub(startTime)).
				Msg(gc.Message)

			return err
		}
	}
}
