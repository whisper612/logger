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
	s.logger.SetDatePrefixFormat(lg.DateFull)
	s.logger.Log("log!")
	s.logger.Warn("warn!")
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat(lg.DayMonthHoursMinute)
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat(lg.DayMonthYear)
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat(lg.DayMonth)
	s.logger.Error(errors.New("error!"))

	s.logger.SetDatePrefixFormat(lg.HoursMinuteSeconds)
	s.logger.Error(errors.New("error!"))
}
