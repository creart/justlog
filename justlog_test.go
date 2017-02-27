package justlog

import (
	"testing"
)

var (
	logger = NewWithHandlers(NewConsoleHandler(DebugLevel), NewFileHandler("test.txt", ErrorLevel)).SetFormatters(NewFormatter("%{LEVEL}:%{TIME} -> %{MESSAGE}"))
)

func TestLogger(t *testing.T) {
	logger.Debug("yeah", "wow")
	logger.Error("Output to the file.")
}
