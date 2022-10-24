package someservice

import (
	"errors"

	lg "github.com/whisper612/logger/pkg/logger"
)

type SomeService struct {
	logger lg.Logger
}

// NewService...
func NewService(logger lg.Logger) *SomeService {
	return &SomeService{
		logger: logger,
	}
}

func (s SomeService) Serve() {
	s.logger.SetDatePrefixFormat("")
	s.logger.Log("log!")
	s.logger.Warn("warn!")
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat("d-m H:m")
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat("d-m-Y")
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat("d-m")
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat("H:i:s")
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat("")
	s.logger.Error(errors.New("error!"))
}
