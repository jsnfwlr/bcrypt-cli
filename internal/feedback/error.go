package feedback

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func EnableCaller() {
	callerInLogs = true
}

var callerInLogs = false

// Make sure we the file and line number of the caller don't point to this file
func getCaller(skip int) (string, int, bool) {
	var file string
	var line int
	var ok bool

	_, curFile, _, _ := runtime.Caller(0)
	for {
		_, file, line, ok = runtime.Caller(skip)
		if file != curFile {
			break
		}
		skip++
	}
	return file, line, ok
}

func HandleFatalErr(err error) {
	if err == nil {
		return
	}
	if callerInLogs {
		file, line, ok := getCaller(1)
		if ok {
			fmt.Fprintf(feedbackDestination.err, "%s:%d - Error: %+v\n", file, line, err)
			os.Exit(1)
		}
	}
	fmt.Fprintf(feedbackDestination.err, "Error: %+v\n", err)
	os.Exit(1)
}

func HandleErr(err error) bool {
	if err == nil {
		return false
	}
	if callerInLogs {
		file, line, ok := getCaller(1)
		if ok {
			fmt.Fprintf(feedbackDestination.err, "%s:%d - Error: %+v\n", file, line, err)
			return true
		}
	}
	fmt.Fprintf(feedbackDestination.err, "Error: %+v\n", err)
	return true
}

func HandleWErr(format string, err error, args ...any) bool {
	if err == nil {
		return false
	}
	comboArgs := append([]any{err}, args...)
	switch {
	case format == "" && len(args) == 0:
		return HandleErr(err)
	case format != "" && len(args) == 0:
		return HandleErr(fmt.Errorf(format, err))
	case format == "" && len(args) > 0:
		format = fmt.Sprintf("%s%s", "%w", strings.Repeat(" %v", len(comboArgs)-1))
		return HandleErr(fmt.Errorf(format, comboArgs...))
	default:
		return HandleErr(fmt.Errorf(format, comboArgs...))
	}
}

func HandleFatalWErr(format string, err error, args ...any) {
	if err == nil {
		return
	}
	comboArgs := append([]any{err}, args...)
	switch {
	case format == "" && len(args) == 0:
		HandleFatalErr(err)
	case format != "" && len(args) == 0:
		HandleFatalErr(fmt.Errorf(format, err))
	case format == "" && len(args) > 0:
		format = fmt.Sprintf("%s%s", "%w", strings.Repeat(" %v", len(comboArgs)-1))
		HandleFatalErr(fmt.Errorf(format, comboArgs...))
	default:
		HandleFatalErr(fmt.Errorf(format, comboArgs...))
	}
}
