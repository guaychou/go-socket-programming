package main

import (
	"bufio"
	"encoding/json"
	"log"
	"math/rand"
	"net"
	"github.com/guaychou/go-socket-programming/models"
	"os"
	"fmt"
	"strings"
)

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Server not found")
	}
	nama,nim:=input()
	text:=models.Message{
		Id: randomInt(1,101),
		Nama: nama,
		Nim: nim,
	}
	jsonText,err:=json.Marshal(text)
		conn.Write(jsonText)
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