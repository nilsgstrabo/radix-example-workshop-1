package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	// fmt.Println("computing")

	// // Writable file
	// f2, err := os.OpenFile("/tmp/mydata2.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	logrus.Errorf("Open: %v", err)
	// 	//panic(err)
	// }
	// defer f2.Close()

	// f2.Write([]byte("hello 2"))

	// // No permission
	// f, err := os.OpenFile("/tmp/mydata.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	// if err != nil {
	// 	logrus.Errorf("Open: %v", err)
	// }
	// defer f.Close()

	// if _, err = f.Write([]byte("hello")); err != nil {
	// 	logrus.Errorf("Write: %v", err)
	// }

	// az, err := os.Open("/compute/.azure/sp_credentials.json")
	// if err != nil {
	// 	logrus.Errorf("Open: %v", err)
	// }
	// defer az.Close()

	// bytes, err := ioutil.ReadAll(az)
	// if err != nil {
	// 	logrus.Errorf("Read az: %v", err)
	// } else {
	// 	logrus.Infof("File az has %v bytes", len(bytes))
	// }

	// fmt.Println("done")
	srv := startServer()
	wg := sync.WaitGroup{}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	ticker := time.NewTicker(2 * time.Second).C

	timer := time.NewTimer(60 * time.Second).C
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			log.Infof("Exiting goroutine")
		}()
		for {
			select {
			case t := <-ticker:
				log.Infof("Received tick on %v", t)
			case t := <-timer:
				log.Infof("Timeout on %v", t)
				return
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Stopping")
	srv.Shutdown(context.Background())
	fmt.Println("Server stopped")
	wg.Wait()
}

func startServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}
	http.Handle("/", http.HandlerFunc(serve))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	return srv
}

func serve(w http.ResponseWriter, r *http.Request) {
	log.Infof("recevied request %v", r.RequestURI)
	io.WriteString(w, "hello world")
}
