package logger

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type level int8

const PanicLevel level = 0

func init() {
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func Log(err error, level level) {
	if err != nil {
		errNew := errors.New(err.Error())
		switch level {
		case PanicLevel:
			log.Panic().Stack().Err(errNew).Send()
		}
	}
}
