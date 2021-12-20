package logger

import (
	log "github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"strings"
)

func Log() *log.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("cannot get context for logger")
	}
	filename := file[strings.LastIndex(file, "/")+1:] + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return log.WithField("file", filename).WithField("function", fn)
}
