package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("computing")

	// Writable file
	f2, err := os.OpenFile("/tmp/mydata2.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logrus.Errorf("Open: %v", err)
		//panic(err)
	}
	defer f2.Close()

	f2.Write([]byte("hello 2"))

	// No permission
	f, err := os.OpenFile("/tmp/mydata.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		logrus.Errorf("Open: %v", err)
	}
	defer f.Close()

	if _, err = f.Write([]byte("hello")); err != nil {
		logrus.Errorf("Write: %v", err)
	}

	az, err := os.Open("/compute/.azure/sp_credentials.json")
	if err != nil {
		logrus.Errorf("Open: %v", err)
	}
	defer az.Close()

	bytes, err := ioutil.ReadAll(az)
	if err != nil {
		logrus.Errorf("Read az: %v", err)
	} else {
		logrus.Infof("File az has %v bytes", len(bytes))
	}

	time.Sleep(10 * time.Minute)
	fmt.Println("done")
}
