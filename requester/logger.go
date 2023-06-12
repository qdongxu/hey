package requester

import (
	"github.com/rs/zerolog"
	"os"
)

var logger zerolog.Logger
var f *os.File

func InitLogger(doLog bool, filepath string) *os.File {
	if !doLog {
		logger = zerolog.Nop()
		return nil
	}

	var err error
	f, err = os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger = zerolog.New(f)
	return f
}

func Logger() *zerolog.Logger {
	return &logger
}
