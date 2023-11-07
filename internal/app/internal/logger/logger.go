package logger

type Logger interface {
	Debug(string)
	Info(string)
	Warn(string)
	Error(error)
	Fatal(string)
}
