package feedback

import (
	"fmt"
	"io"
	"os"
)

var (
	maxNoise            Level = 3
	feedbackDestination       = feedbackDest{
		std: os.Stdout,
		err: os.Stderr,
	}
)

type Level int

const (
	levelCap Level = 4
	Required Level = 0
	Error    Level = 1
	Warning  Level = 2
	Info     Level = 3
	Debug    Level = 4
)

type feedbackDest struct {
	std io.Writer
	err io.Writer
}

func SuppressNoise(levels Level) {
	target := (levelCap - levels)
	if target < 0 {
		target = 0
	}
	maxNoise = target
}

func GetNoiseLimit() Level {
	return maxNoise
}

func SetDestination(args ...io.Writer) {
	if len(args) == 0 {
		return
	}
	if len(args) == 1 {
		feedbackDestination = feedbackDest{
			std: args[0],
			err: args[0],
		}
		return
	}

	feedbackDestination = feedbackDest{
		std: args[0],
		err: args[1],
	}
}

func Printf(level Level, format string, args ...interface{}) {
	if !exceedsLimit(level) {
		fmt.Fprintf(feedbackDestination.std, format, args...)
	}
}

func Print(level Level, format string, args ...interface{}) {
	if !exceedsLimit(level) {
		if format[len(format)-1:] != "\n" {
			format += "\n"
		}
		fmt.Fprintf(feedbackDestination.std, format, args...)
	}
}

func Println(level Level, args ...interface{}) {
	if !exceedsLimit(level) {
		fmt.Fprintln(feedbackDestination.std, args...)
	}
}

func exceedsLimit(level Level) bool {
	return level > maxNoise
}

func NoiseLevel(level int) Level {
	return Level(level)
}
