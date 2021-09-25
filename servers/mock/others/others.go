package mockothers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mock "github.com/shudipta/play-with-locust/mock"
)

func pushNotification(req *http.Request) mock.Response {
	return mock.NewResponse(http.StatusAccepted, string("success"))
}

func sendSocketMessage(req *http.Request) mock.Response {
	return mock.NewResponse(http.StatusAccepted, string("success"))
}

func saveMeta(req *http.Request) mock.Response {
	return mock.NewResponse(http.StatusAccepted, string("success"))
}

// RunMockKuego runs a mock server of kuego
func RunMockKuego() {
	router := mux.NewRouter()

	router.HandleFunc("/notifications", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, pushNotification(req))
	}).Methods("POST")

	fmt.Println("running the mock kuego ...")
	log.Fatalln(http.ListenAndServe(":8003", router))
}

// RunMockSocketServer runs a mock server of socket
func RunMockSocketServer() {
	router := mux.NewRouter()

	router.HandleFunc("/send", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, sendSocketMessage(req))
	}).Methods("POST")

	fmt.Println("running the mock socket server ...")
	log.Fatalln(http.ListenAndServe(":8004", router))
}

// RunMockMeta runs a mock server of meta
func RunMockMeta() {
	router := mux.NewRouter()

	router.HandleFunc("/event", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, saveMeta(req))
	}).Methods("POST")

	fmt.Println("running the mock meta ...")
	log.Fatalln(http.ListenAndServe(":8005", router))
}
