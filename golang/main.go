package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	nsq "github.com/nsqio/go-nsq"
)

func main() {
	config := nsq.NewConfig()
	w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	for i := 0; i < 10; i++ {
		w.Publish("topic_string", []byte(fmt.Sprintf("string%d", i)))
	}

	jsonData := []string{}
	jsonData = append(jsonData, `
								{
								    "func":"BarFuncAdd",
								    "params":[
								        "0.5",
								        "0.51"
								    ]
								}
								`)

	jsonData = append(jsonData, `
								{
								    "func":"FooFuncSwap",
								    "params":[
								        "a",
								        "b"
								    ]
								}
								`)

	jsonData = append(jsonData, `
								{
								    "func":"FooFuncSwap",
								    "params":[
								        "1",
								        "2"
								    ]
								}
								`)

	jsonData = append(jsonData, `
								{
								    "func":"FakeMethod",
								    "params":[
								        "a",
								        "b"
								    ]
								}
								`)

	for _, j := range jsonData {
		w.Publish("topic_json", []byte(j))
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	w.Stop()
}
