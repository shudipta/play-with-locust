package mocka2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	mock "github.com/shudipta/play-with-locust/mock"
	"github.com/tamalsaha/go-oneliners"
)

// AllocateDriver is a flag var that tells whether to allocate driver or not
var AllocateDriver bool

// PairLimit is the max number of orders a driver can accept within a specified time
var PairLimit int

var isCanceled map[string]bool

/*
{
    "order_id": 627949,
    "status": "success",
    "message": "task_13327397-72bb-4a2e-a0cd-713b3c731791"
}
202
*/
func allocates(req *http.Request) mock.Response {
	var reqBody AllocateReq
	var err error

	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/allocates: error getting json data")
	}
	oneliners.PrettyJson(reqBody)

	var respBody []byte
	if respBody, err = json.MarshalIndent(ResponseBody{
		OrderHash: reqBody.Order.Hash,
		Status:    "success",
		Message:   "task_13327397-72bb-4a2e-a0cd-713b3c731791",
	}, "", " "); err != nil {
		return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/allocates: error marshalling json response")
	}

	go func() {
		if isCanceled[reqBody.Order.Hash] {
			return
		}

		time.Sleep(reqBody.Request.Timeout)

		if AllocateDriver {
			driversStr := strings.Split(reqBody.Request.ExcludeDrivers, ",")
			exdrivers := make([]int, len(driversStr))
			for i := range driversStr {
				exdrivers[i], _ = strconv.Atoi(driversStr[i])
			}
			callback(reqBody.Order.Hash, reqBody.Order.OrderMeta.UserID, exdrivers...)
		} else {
			timeout(reqBody.Order.Hash, reqBody.Request.Timeout)
		}
	}()

	return mock.NewResponse(http.StatusAccepted, string(respBody))
}

/*
{
   "hash": "asd",
   "user_id": 123,
   "service_type": "food"
}

{
    "order_id": 627949,
    "status": "success",
    "message": "A2 Request cancelled"
}
200
*/

func release(req *http.Request) mock.Response {
	var reqBody ReleaseSpecs
	var err error

	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/release: error getting json data")
	}
	oneliners.PrettyJson(reqBody)

	isCanceled[reqBody.OrderHash] = true

	var respBody []byte
	if respBody, err = json.MarshalIndent(ResponseBody{
		OrderHash: reqBody.OrderHash,
		Status:    "success",
		Message:   "A2 Request release",
	}, "", " "); err != nil {
		return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/release: error marshalling json response")
	}

	return mock.NewResponse(http.StatusAccepted, string(respBody))
}

func cancel(req *http.Request) mock.Response {
	var reqBody CancelSpecs
	var err error

	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/cancel: error getting json data")
	}
	oneliners.PrettyJson(reqBody)

	isCanceled[reqBody.OrderHash] = true

	var respBody []byte
	if respBody, err = json.MarshalIndent(ResponseBody{
		OrderHash: reqBody.OrderHash,
		Status:    "success",
		Message:   "A2 Request cancelled",
	}, "", " "); err != nil {
		return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/cancel: error marshalling json response")
	}

	return mock.NewResponse(http.StatusAccepted, string(respBody))
}

// RunMockA2 runs mock a2
func RunMockA2() {
	isCanceled = make(map[string]bool)
	router := mux.NewRouter()

	router.HandleFunc("/v1/allocates", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, allocates(req))
	}).Methods("POST")

	router.HandleFunc("/v1/release", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, release(req))
	}).Methods("PATCH")

	router.HandleFunc("/v1/cancel", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, cancel(req))
	}).Methods("PATCH")

	fmt.Println("running the mock a2 ...")
	log.Fatalln(http.ListenAndServe(":8002", router))
}
