package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	infoLog *log.Logger
	warnLog *log.Logger
	errLog  *log.Logger

	ldir = flag.String("ldir", "./", "log dir")
)

func init() {

	flag.Parse()

	fi, err := os.OpenFile(fmt.Sprintf("%s/%s-info-%s-%d.log", *ldir, filepath.Base(os.Args[0]),
		time.Now().Format("200601021504"), os.Getpid()), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	infoLog = log.New(fi, "[INFO] ", log.LstdFlags|log.Llongfile)

	fw, err := os.OpenFile(fmt.Sprintf("%s/%s-warn-%s-%d.log", *ldir, filepath.Base(os.Args[0]),
		time.Now().Format("200601021504"), os.Getpid()), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	warnLog = log.New(fw, "[WARN] ", log.LstdFlags|log.Llongfile)

	fe, err := os.OpenFile(fmt.Sprintf("%s/%s-error-%s-%d.log", *ldir, filepath.Base(os.Args[0]),
		time.Now().Format("200601021504"), os.Getpid()), os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	errLog = log.New(fe, "[ERROR] ", log.LstdFlags|log.Llongfile)
}

func ErrorFatal(v ...interface{}) {
	errLog.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func ErrorFatalf(format string, v ...interface{}) {
	errLog.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func ErrorFatalln(v ...interface{}) {
	errLog.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func ErrorPanic(v ...interface{}) {
	s := fmt.Sprint(v...)
	errLog.Output(2, s)
	panic(s)
}

func ErrorPanicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	errLog.Output(2, s)
	panic(s)
}

func ErrorPanicln(v ...interface{}) {
	s := fmt.Sprint(v...)
	errLog.Output(2, s)
	panic(s)
}

func ErrorPrint(v ...interface{}) {
	errLog.Output(2, fmt.Sprint(v...))
}

func ErrorPrintf(format string, v ...interface{}) {
	errLog.Output(2, fmt.Sprintf(format, v...))
}

func ErrorPrintln(v ...interface{}) {
	errLog.Output(2, fmt.Sprint(v...))
}

func InfoPrint(v ...interface{}) {
	infoLog.Output(2, fmt.Sprint(v...))
}

func InfoPrintf(format string, v ...interface{}) {
	infoLog.Output(2, fmt.Sprintf(format, v...))
}

func InfoPrintln(v ...interface{}) {
	infoLog.Output(2, fmt.Sprint(v...))
}

func WarnPrint(v ...interface{}) {
	warnLog.Output(2, fmt.Sprint(v...))
}

func WarnPrintf(format string, v ...interface{}) {
	warnLog.Output(2, fmt.Sprintf(format, v...))
}

func WarnPrintln(v ...interface{}) {
	warnLog.Output(2, fmt.Sprint(v...))
}
