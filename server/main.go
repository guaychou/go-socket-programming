package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/guaychou/go-socket-programming/config"
	"github.com/guaychou/go-socket-programming/models"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	done:=make(chan os.Signal,1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ln, err := net.Listen("tcp", config.SERVER_PORT)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started listening on port "+config.SERVER_PORT)
	log.Info("Press ctrl+c to exit")
	go func() {
		for {
			conn, err := ln.Accept() //Accepting connection
			if err != nil {
				log.Fatal(err)
			}
			log.Info("New client connected: "+conn.RemoteAddr().String())
			go func() {
				for  {
					err:=handleConnection(conn) // Handling Connection
					if err!=nil{
						if err==io.EOF{
							log.Info("Client disconeected: "+ conn.RemoteAddr().String())
							break
						}
						log.Fatal(err)
					}
				}
			}()
		}
	}()
	<-done
	log.Info("Server stopped . . . ")
}

func handleConnection(conn net.Conn) error {
	data :=models.MessageArray{}
	scanner:=bufio.NewReader(conn)
	message,err:=scanner.ReadBytes(0x0A)
	if err!=nil{
		return err
	}
	err=json.Unmarshal(message,&data)
	if err!=nil{
		return err
	}
	fmt.Println("======================= DATA =============================")
	fmt.Println("Nama Mahasiswa: ", data.Message[0].Nama)
	fmt.Println("Nim Mahasiswa: ", data.Message[0].Nim)
	fmt.Println("ID client: ", data.Message[0].Id)
	fmt.Println("======================== END =============================\n")
	return nil
}