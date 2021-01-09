/*
Copyright 2020 RS4
@Author: Weny Xu
@Date: 2021/01/07 5:33
*/

/*
	Module for logger
*/
package logger

import (
	"log"
)

// Logger provides an abstract interface for logging from Reporters.
// Applications can provide their own implementation of this interface to adapt
// reporters logging to whatever logging library they prefer (stdlib log,
// logrus, go-logging, etc).
type Logger interface {
	// Error logs a message at error priority
	Error(msg string)

	// Infof logs a message at info priority
	Infof(msg string, args ...interface{})
}

// StdLogger is implementation of the Logger interface that delegates to default `log` package
var StdLogger = &logger{
	errorFunc: ErrorFunc,
	infofFunc: InfofFunc,
}

type logger struct {
	errorFunc func(msg string)
	infofFunc func(msg string, args ...interface{})
}

func (l logger) Error(msg string) {
	l.errorFunc(msg)
}

// Infof logs a message at info priority
func (l logger) Infof(msg string, args ...interface{}) {
	l.infofFunc(msg, args...)
}

// NullLogger is implementation of the Logger interface that delegates to default `log` package
var NullLogger = &logger{
	errorFunc: EmptyErrorFunc,
	infofFunc: EmptyInfofFunc,
}

func EmptyErrorFunc(msg string) {

}
func EmptyInfofFunc(msg string, args ...interface{}) {

}

func ErrorFunc(msg string) {
	log.Printf("ERROR: %s", msg)
}
func InfofFunc(msg string, args ...interface{}) {
	log.Printf(msg, args...)
}
