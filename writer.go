package zlog

const (
	// defines the type to write log entries to write as
	ErrorLog = 1
	InfoLog  = 2
)

// Writer wraps zlog so it can be used as a stdlib logger, so it can be passed
// into the http error log for example
type Writer struct {
	Logger  *Logger
	Logtype int
}

// Write writes to zlog logger
func (z *Writer) Write(p []byte) (n int, err error) {

	switch z.Logtype {
	case InfoLog:
		z.Logger.Info(string(p))

	case ErrorLog:
		z.Logger.Error(string(p))
	}
	return len(p), nil
}
