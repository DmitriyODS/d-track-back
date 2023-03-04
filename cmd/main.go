package main

import (
	"context"
	"errors"
	"gitlab.com/ddda/d-track/d-track-back/endpoints"
	v1 "gitlab.com/ddda/d-track/d-track-back/endpoints/v1"
	"gitlab.com/ddda/d-track/d-track-back/global"
	"gitlab.com/ddda/d-track/d-track-back/implementation"
	"gitlab.com/ddda/d-track/d-track-back/middleware"
	"gitlab.com/ddda/d-track/d-track-back/service"
	"gitlab.com/ddda/d-track/d-track-back/store"
	httpTransport "gitlab.com/ddda/d-track/d-track-back/transport/http"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting server ...")
	defer log.Println("Server done")

	// подключаемся к репозиторию
	storeClient := store.NewStore(global.Config.URLPathDB)
	defer storeClient.CloseConnect()

	// создаём сервис
	var svc service.Service
	svc = implementation.NewBasicService(storeClient)

	// добавляем мидлвари
	svc = middleware.AddAuthMiddleware(svc)
	svc = middleware.AddLoggerMiddleware(svc)

	// создаём конечные точки для сервиса
	// в создатель конечных точек передаём мидлварь обработки транзакций
	svcEpsV1 := v1.MakeEndpoints(svc, endpoints.TransactionMiddleware(storeClient))

	// создаём транспортный протокол Http
	router := httpTransport.NewHttpTransport(svcEpsV1)
	server := http.Server{
		Addr:    global.Config.SrvAddr,
		Handler: router,
	}

	// создаём слушатель сигналов, которые хотят нас прикрыть
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGKILL)
	defer signal.Stop(quit)

	// запускаем горутину с нашим сервером
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Err in server: %s\n", err)
		}
	}()

	// ждём сигналов о завершении работы
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
