package log

import (
    "testing"
)

type spyLogger struct {
    called bool
}

func (s *spyLogger) Debug(output string) {
    s.called = true
}

func (s *spyLogger) Info(output string) {
    s.called = true
}

func (s *spyLogger) Error(output string) {
    s.called = true
}

func TestVerifyDebugCalled(t *testing.T) {
    spyLog := spyLogger{called: false}
    logHandler := NewLogHandler(&spyLog)
    logHandler.Debug("testing debug called")
    if !spyLog.called {
        t.Errorf("Could not verify that Debug was called")
    }
}

func TestVerifyInfoCalled(t *testing.T) {
    spyLog := spyLogger{called: false}
    logHandler := NewLogHandler(&spyLog)
    logHandler.Info("testing info called")
    if !spyLog.called {
        t.Errorf("Could not verify that Info was called")
    }
}
func TestVerifyErrorCalled(t *testing.T) {
    spyLog := spyLogger{called: false}
    logHandler := NewLogHandler(&spyLog)
    logHandler.Error("testing error called")
    if !spyLog.called {
        t.Errorf("Could not verify that Error was called")
    }
}

