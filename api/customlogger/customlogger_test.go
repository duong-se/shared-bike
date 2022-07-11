package customlogger

import (
	"shared-bike/apperrors"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/stretchr/testify/suite"
)

type CustomLoggerTestSuite struct {
	suite.Suite
	logger *logger
}

func (s *CustomLoggerTestSuite) SetupTest() {
	e := echo.New()
	contextLogger := NewContextLogger(e.Logger)
	s.logger = contextLogger

}
func TestCustomLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomLoggerTestSuite))
}

func (s *CustomLoggerTestSuite) TestSetRequestID() {
	s.logger.SetHeader("mockHeader")
	s.logger.SetRequestID("requestID")
	s.Equal("requestID", s.logger.requestID)
}

func (s *CustomLoggerTestSuite) TestOutput() {
	result := s.logger.Output()
	s.NotNil(result)
}

func (s *CustomLoggerTestSuite) TestAndSetPrefixPrefix() {
	s.logger.SetPrefix("prefix")
	result := s.logger.Prefix()
	s.Equal(result, "prefix")
	s.logger.SetPrefix("prefix1")
	newResult := s.logger.Prefix()
	s.Equal(newResult, "prefix1")
}

func (s *CustomLoggerTestSuite) TestLogLevel() {
	s.logger.SetLevel(log.INFO)
	result := s.logger.Level()
	s.Equal(result, log.INFO)
	s.logger.SetLevel(log.DEBUG)
	newResult := s.logger.Level()
	s.Equal(newResult, log.DEBUG)
}

func (s *CustomLoggerTestSuite) TestPrint() {
	s.logger.Print("mockPrint")
	s.logger.Printf("%s", "mockPrint")
	s.logger.Printj(log.JSON{"a": "mockPrint", "foo": "bar"})
	s.logger.Debug("mockDebug")
	s.logger.Debugf("%s", "mockDebug")
	s.logger.Debugj(log.JSON{"a": "mockDebug", "foo": "bar"})
	s.logger.Info("mockInfo")
	s.logger.Infof("%s", "mockInfo")
	s.logger.Infoj(log.JSON{"a": "mockInfo", "foo": "bar"})
	s.logger.Warn("mockWarn")
	s.logger.Warnf("%s", "mockWarn")
	s.logger.Warnj(log.JSON{"a": "mockWarn", "foo": "bar"})
	s.logger.Error("mockError", apperrors.ErrInternalServerError)
	s.logger.Errorf("%s", "mockError")
	s.logger.Errorj(log.JSON{"a": "mockError", "foo": "bar"})
}
