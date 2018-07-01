package main

import (
	"practice/ch5/2018/0630/work-pool/service"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
	"io/ioutil"
)

type serviceS struct {
	ServerName string
	ServerIP   string
}

type serverslice struct {
	Servers []serviceS
}

var (
	MaxWorkers = 100
	MaxQueue   = 100
)

func main() {

	//启动服务
	serv := service.NewService(MaxWorkers, MaxQueue)

	serv.Start()
	defer serv.Stop()

	//处理海量任务
	http.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var jobs serverslice

		var body, _ = ioutil.ReadAll(r.Body)

		err := json.Unmarshal(body, &jobs)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		//处理任务
		for _, job := range jobs.Servers {
			serv.AddJob(job)
		}

		//处理完成
		w.WriteHeader(http.StatusOK)
	})

	//启动web服务
	log.Fatal(http.ListenAndServe(":8080", nil))

}
