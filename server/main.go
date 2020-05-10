package main

import (
	"github.com/guaychou/go-socket-programming/models"
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	done:=make(chan os.Signal,1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Server started . . . ")
	log.Info("Press ctrl+c to exit")
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			go handleConnection(conn)
		}
	}()
	<-done
	log.Info("Server stopeed . . . ")
}

func handleConnection(conn net.Conn)  {
	var mahasiswa models.Message
	message, _ := bufio.NewReader(conn).ReadBytes(0x7d)
	err:=json.Unmarshal(message, &mahasiswa)
	if err!=nil{
		log.Fatal("Error disini",err)
	}
	fmt.Println("======================= DATA ============================")
	fmt.Println("Nama Mahasiswa: ", mahasiswa.Nama)
	fmt.Println("Nim Mahasiswa: ", mahasiswa.Nim)
	fmt.Println("ID client: ", mahasiswa.Id)
	fmt.Println("======================== END =============================\n")
}