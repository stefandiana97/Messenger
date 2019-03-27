package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func main(){

	serverFrontend();
	serverBackend();
	//http.HandleFunc()

	err := http.ListenAndServe("localhost:1324",nil)
	if err != nil {
		log.Fatalf("could not start webserver: %v ", err)
	}

}


func serverFrontend(){
	fileserver := http.FileServer(http.Dir("./frontend/build/"))
	http.Handle("/",fileserver)

}

//var clients []Client //[]*Client pointer


var clients []*Client //[]*Client pointer

type Client struct{
	//ID int
	Conn *websocket.Conn
	Signal chan string
}

func serverBackend(){

	//upgrader := websocket.Upgrader{}
	//http.HandleFunc("")

	http.HandleFunc("/ws",func(w http.ResponseWriter, r *http.Request){
			upgrader:= websocket.Upgrader{}

			//client := Client{} //&Client{}
			client := &Client{}

			conn,err:=upgrader.Upgrade(w,r,nil)
			if err!= nil{
				log.Fatalf("could not upgrade connect: %v ", err)
			}

			client.Conn = conn
			clients = append(clients,client)

			err = client.Conn.WriteMessage(websocket.TextMessage,[]byte("{\"Text\": \"First message\"}"))
			err = client.Conn.WriteMessage(websocket.TextMessage,[]byte("{\"Text\": \"Second message\"}"))
			if err!= nil{
				log.Fatalf("could not send message: %v ", err)
			}

			//for key,client := range clients{
			//	_ = client.Conn.WriteMessage(messageType,msg)
			//}


			//var wg sync.WaitGroup

			//wg.Add(1)
			//asculta dupa mesaje Thread 1
			go func(){ //(wg sync.WaitGroup) {
				for {
					messageType, msg, _ := client.Conn.ReadMessage()

					for key, client := range clients {
						_ = client.Conn.WriteMessage(messageType, msg)
						fmt.Printf("key= %d\n", key)
					}

					//fmt.Printf("message received: %s (%d message type)", msg, messageType)
					//
					//_ = client.Conn.WriteMessage(messageType,msg)

					//if false==true {
					//	wg.Done()
					//	break
					//}
				}
			}() //(wg)

			//wg.Add(1)
			//sms Thread 2
			go func() {
				for {
					fmt.Println("working...")
					//wg.Done()
				}
			}()
			//wg.Wait()

			//go fmt.Println("whatever1")
			//go fmt.Println("whatever2")
			//go fmt.Println("whatever3")
			//go fmt.Println("whatever4")

			////asculta dupa mesaje Thread 1 Gheorghe
			//go func() {
			//	for {
			//		messageType, msg, _ := client.Conn.ReadMessage()
			//		fmt.Printf("message received: %s (%d message type)", msg, messageType)
			//
			//		for _, client := range clients {
			//			client.Signal <- string(msg)
			//			//msg := <- client.Signal
			//			//_ = client.Conn.WriteMessage(messageType, msg)
			//		}
			//	}
			//}()
			//
			//
			////trimite sms on signal Ion
			//go func(signal <- chan string) {
			//	for {
			//		msg := <-client.Signal
			//		fmt.Printf("message received: %s ", msg)
			//	}
			//}(client.Signal)
	})
}

