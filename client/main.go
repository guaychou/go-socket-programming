package main

import (
	"bufio"
	"encoding/json"
	"github.com/guaychou/go-socket-programming/config"
	"log"
	"math/rand"
	"net"
	"github.com/guaychou/go-socket-programming/models"
	"os"
	"fmt"
	"strings"
)

func main(){
	message:=models.MessageArray{}
	conn, err := net.Dial("tcp", "127.0.0.1"+config.SERVER_PORT)
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