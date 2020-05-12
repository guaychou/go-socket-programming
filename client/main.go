package main

import (
	"bufio"
	"encoding/json"
	"github.com/guaychou/go-socket-programming/config"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"github.com/guaychou/go-socket-programming/models"
	"os"
	"fmt"
	"strings"
	"time"
)
var timer int = 0

func init (){
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
}
func main(){
	message:=models.MessageArray{}
	conn, err := dialServer()
	if err != nil {
		log.Fatal("Server not found")
	}
	for {
		nama,nim:=input()
		text:=models.Data{
			Id: randomInt(1,501),
			Nama: nama,
			Nim: nim,
		}
		message.AddMessage(text)
		jsonText,err:=json.Marshal(message)
		if err!=nil{
			log.Fatal(err)
		}
		jsonText=append(jsonText,0x0A)
		conn.Write(jsonText)
		message.Message=nil
	}
}

func dialServer()(net.Conn,error){
	conn,err:=net.Dial("tcp", "127.0.0.1"+config.SERVER_PORT)
	if err!=nil{
		log.WithFields(log.Fields{
			"Status":"Down",
		}).Info("Server is not reachable")
		if timer==5{
			log.Fatal("Server down, Exiting . . . ")
		}
		timer++
		time.Sleep(5 * time.Second)
		return dialServer()
	}
	return conn,err
}

func input()(string,string){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama: ")
	nama, _ := reader.ReadString('\n')
	nama=strings.TrimSuffix(nama,"\n")
	fmt.Print("Nim: ")
	nim, _ := reader.ReadString('\n')
	nim=strings.TrimSuffix(nim,"\n")
	return nama,nim
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}