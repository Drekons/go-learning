package main

import (
	"bytes"
	"context"
	"dz28/pkg/helper"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const proxyPort = "8080"

var (
	counter            = 0
	firstInstanceHost  = "http://localhost:8081"
	secondInstanceHost = "http://localhost:8082"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.AllowContentEncoding("Content-Type: application/json; charset=utf-8"))

	r.Post("/*", handleProxy)
	r.Get("/*", handleProxy)
	r.Put("/*", handleProxy)
	r.Delete("/*", handleProxy)

	server := &http.Server{Addr: ":" + proxyPort, Handler: r}
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	log.Println("Proxy running...")

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	<-serverCtx.Done()
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	instance := getInstance()

	body := bodyRead(r.Body)
	fmt.Printf("request | %s %s: %s\n", r.Method, r.RequestURI, body)

	resp := newRequest(r.Method, instance+r.RequestURI, body)
	defer helper.BodyCloseResponse(resp)

	responseBody := bodyRead(resp.Body)

	fmt.Printf("response | %s: %s\n", instance, responseBody)

	w.WriteHeader(resp.StatusCode)
	if err := helper.WriteJson(responseBody, w); err != nil {
		log.Println(err)
		return
	}
}

func getInstance() string {
	if counter == 0 {
		counter++
		return firstInstanceHost
	}

	counter--
	return secondInstanceHost
}

func newRequest(method, url, body string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return resp
}

func bodyRead(body io.ReadCloser) string {
	textBytes, err := ioutil.ReadAll(body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(textBytes)
}
