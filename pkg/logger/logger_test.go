package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"testing"
)

func TestLogger_Output(t *testing.T) {
	l := NewLogger(os.Stdout, "", log.LstdFlags)
	l.Output(LevelInfo, "fianf")
}

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