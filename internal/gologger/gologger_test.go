package gologger

import (
	"errors"
	"testing"
)

func testDebug() {
	Debug("test debug func")
}
func TestLogger(t *testing.T) {
	Logger.SetLevel(DebugLevel)
	WithField("field1", "value").WithField("field2", "value21").Info()
	Debugln("test debug")
	testDebug()
	Infoln("test info")
	Warnln("test warn")
	Errorln("test error")
	WithError(errors.New("test error")).Errorln("error")
}
