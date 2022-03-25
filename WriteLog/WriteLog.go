package WriteLog

import (
	"io"
	"os"
	"strconv"
	"time"
)

const (
	FILEPATH = "log/"
	FORMAT   = "20060102"
	LineFeed = "\r\n"
)

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

func WriteFile(fileName, msg string) error {
	var path = FILEPATH + time.Now().Format(FORMAT) + "/"
	if !IsExist(path) {
		return CreateDir(path)
	}

	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(path+fileName+"_"+time.Now().Format(FORMAT)+"_"+
		strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+
		strconv.Itoa(time.Now().Second())+".txt",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	_, err = io.WriteString(f, LineFeed+msg)
	defer f.Close()
	return err
}
