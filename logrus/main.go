package main

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var logger = log.New()

func main() {

	test1()

	test2()

	test3()

	test4()
}

func test1() {
	fmt.Println("---------------------test 1--------------------------")
	//init
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	//打印文件、行数
	log.SetReportCaller(true)

	log.Error("test log Error")
	log.Info("test log Info")
}

func test2() {
	fmt.Println("---------------------test 2--------------------------")
	//init
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	//打印文件、行数
	log.SetReportCaller(true)

	log.Error("test log Error")
	log.Info("test log Info")
}

func test3() {
	fmt.Println("---------------------test 3--------------------------")
	logger.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
		FullTimestamp:   true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)

	logger.WithFields(log.Fields{
		"name":  "lengchuan",
		"email": "lishuijun1992@gmail.com",
	}).Info("test logrus WithFields")
}

//hook
func test4() {
	fmt.Println("---------------------test 4--------------------------")
	logger.AddHook(newLfsHook(nil, 1024))

	for {
		logger.Info("test file hook")
	}
}

func newLfsHook(logLevel *string, maxRemainCnt uint) log.Hook {
	writer, err := rotatelogs.New(
		"log"+".%Y%m%d%H"+".log",
		// WithLinkName为最新的日志建立软连接,以方便随着找到当前日志文件
		rotatelogs.WithLinkName("log"),

		// WithRotationTime设置日志分割的时间,这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Second*3),

		// WithMaxAge和WithRotationCount二者只能设置一个,
		// WithMaxAge设置文件清理前的最长保存时间,
		// WithRotationCount设置文件清理前最多保存的个数.
		//rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationCount(maxRemainCnt),
	)

	if err != nil {
		log.Errorf("config local file system for logger error: %v", err)
	}

	log.SetLevel(log.WarnLevel)
	lfsHook := lfshook.NewHook(lfshook.WriterMap{
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}, &log.TextFormatter{DisableColors: true})

	return lfsHook
}
