package top

type LoggerInterface interface {
	Debugf(format string, v ...interface{})

	Infof(format string, v ...interface{})

	Warnf(format string, v ...interface{}) error

	Errorf(format string, v ...interface{}) error

	//Criticalf(format string, v ...interface{}) error

	Debug(v ...interface{})

	Info(v ...interface{})

	Warn(v ...interface{}) error

	Error(v ...interface{}) error

	//Critical(v ...interface{}) error
}
