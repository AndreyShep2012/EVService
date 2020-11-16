package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	pbf "github.com/AndreyShep2012/EVService/shippingService/protobuf/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func main() {
	// Установим соединение с сервером
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не могу подключиться: %v", err)
	}
	defer conn.Close()
	client := pbf.NewShippingServiceClient(conn)

	// Передадим в обработку consignment.json,
	// иначе возьмём путь к файлу из аргументов командной строки
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Не возможно распарсить фаил: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Не удалось создать: %v", err)
	}
	log.Printf("Создан: %t", r.Created)
}

//Функция парсит переданный фаил
func parseFile(file string) (*pbf.Consignment, error) {
	var consignment *pbf.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}
