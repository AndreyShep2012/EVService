package main

import (
	"log"
	"net"

	// Импотртируем код протобуфера
	pbf "shippingService/protobuf/pb"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

//IRepository - интерфейс хранилища
type IRepository interface {
	Create(*pbf.Consignment) (*pbf.Consignment, error)
}

// Repository - структура для эмитации хранилища,
// после мы заменим её на настоящие хранилищем
type Repository struct {
	consignments []*pbf.Consignment
}

//Create - создаём новое хранилище
func (repo *Repository) Create(consignment *pbf.Consignment) (*pbf.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

// Служба должна реализовать все методы для удовлетворения сервиса
// которые мы определили в нашем определении протобуфа. Вы можете проверить интерфейсы
// в сгенерированном коде для точных сигнатур методов и т. д.
type service struct {
	repo IRepository
}

// CreateConsignment - мы создали только один метод для нашего сервиса,
// который является методом create, который принимает контекст и запрос
// потом они обрабатываются сервером gRPC.
func (s *service) CreateConsignment(ctx context.Context, req *pbf.Consignment) (*pbf.Response, error) {

	// Сохраним нашу партию в хранидище
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	// Возвращаем сообщение `Response`,
	// которое мы создали в нашем определнии пробуфа
	return &pbf.Response{Created: true, Consignment: consignment}, nil
}

func main() {

	repo := &Repository{}

	// Стартуем наш gRPC сервер для прослушивания tcp
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	// Зарегистрируйте нашу службу на сервере gRPC, это свяжет нашу
	// реализацию с кодом автогенерированного интерфейса для нашего
	// сообщения `Response`, которое мы создали в нашем протобуфе
	pbf.RegisterShippingServiceServer(s, &service{repo})

	// Регистрация службы ответов на сервере gRPC.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
