// Package log provides a logger interface and a global logger instance.
package log

// The Logger interface is used to log messages.
// Use this abstraction to avoid coupling your code to a specific logging library.
// [args] is a variadic parameter that can be used to pass additional arguments to the log message in the form of key-value pairs.
// e.g. log.Info("message", "key1", "value1", "key2", "value2").
type Logger interface {
	InitLogger(levelStr string)    // Initialize the logger with the given level
	Debug(msg string, args ...any) // Log a debug message
	Info(msg string, args ...any)  // Log an info message
	Warn(msg string, args ...any)  // Log a warning message
	Error(msg string, args ...any) // Log an error message
}

var L Logger = &zapLogger{} // Global logger
