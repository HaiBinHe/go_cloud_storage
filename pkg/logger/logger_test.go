package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"testing"
)


func TestLogger_WithFields(t *testing.T) {
	l := NewLogger(&lumberjack.Logger{
		Filename: "testLog",
		MaxSize: 50,
		MaxAge: 10,
	}, "", log.LstdFlags)
	f := Fields{
		"id": "test",
		"name":"test",
	}
	log.Println(l)
	l = l.WithFields(f)
	log.Println(l)
}
func TestLogger_WithCallersFrames(t *testing.T) {
	l := NewLogger(&lumberjack.Logger{
		Filename: "testLog",
		MaxSize: 50,
		MaxAge: 10,
	}, "", log.LstdFlags)
	log.Println(l.callers)
	l = l.WithCallersFrames()
	log.Println(l.callers)

}