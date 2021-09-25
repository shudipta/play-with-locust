package mockCore

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	mock "github.com/shudipta/play-with-locust/mock"
	"github.com/tamalsaha/go-oneliners"
)

/*
{
    'driver': driver_id,
    'status': 'assigned',
    'drivers_tried': drivers_tried,
    'assigned_at': current_time
}

{
    "order_id": 627949,
    "status": "assigned",
    "message": "Assignment success !"
}
200
*/

func assignDriver(req *http.Request) mock.Response {
	vars := mux.Vars(req)
	var reqBody CoreAssignedReq
	var err error

	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		return mock.NewResponse(http.StatusBadRequest,
			fmt.Sprintf(">>>>>>>>>>>>>> /users/%v/orders/%v/dispatch: error getting json data", vars["uid"], vars["oid"]))
	}
	oneliners.PrettyJson(reqBody)

	var respBody []byte
	if respBody, err = json.MarshalIndent(ResponseBody{
		OrderHash: vars["oid"],
		Status:    string(reqBody.Status),
		Message:   "Assignment success !",
	}, "", "  "); err != nil {
		return mock.NewResponse(http.StatusInternalServerError,
			fmt.Sprintf(">>>>>>>>>>>>>> /users/%v/orders/%v:/dispatch error marshalling json response", vars["uid"], vars["oid"]))
	}

	fmt.Println("fmt >>>>>>>\n", string(respBody))
	oneliners.PrettyJson(respBody)
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

func assignmentError(req *http.Request) mock.Response {
	vars := mux.Vars(req)
	var reqBody CoreAssignedReq
	var err error

	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		return mock.NewResponse(http.StatusBadRequest,
			fmt.Sprintf(">>>>>>>>>>>>>> /users/%v/orders/%v: error getting json data", vars["uid"], vars["oid"]))
	}
	oneliners.PrettyJson(reqBody)

	var respBody []byte
	if respBody, err = json.MarshalIndent(ResponseBody{
		OrderHash: "e6eb1f54-5c14-4274-a4c8-ffba17c12f81",
		Status:    "CANCELED",
		Message:   "Assignment failed",
	}, "", "  "); err != nil {
		return mock.NewResponse(http.StatusInternalServerError,
			fmt.Sprintf(">>>>>>>>>>>>>> /users/%v/orders/%v: error marshalling json response", vars["uid"], vars["oid"]))
	}

	fmt.Println("fmt >>>>>>>\n", string(respBody))
	oneliners.PrettyJson(respBody)
	return mock.NewResponse(http.StatusAccepted, string(respBody))
}

/*
{
    'driver': driver_id,
    'status': 'assigned',
    'drivers_tried': drivers_tried,
    'assigned_at': current_time
}

{
    "order_id": 627949,
    "status": "assigned",
    "message": "Update order success !"
}
200
*/

func updateOrder(req *http.Request) mock.Response {
	vars := mux.Vars(req)
	var reqBody CoreAssignedReq
	var err error

	defer req.Body.Close()
	if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		return mock.NewResponse(http.StatusBadRequest,
			fmt.Sprintf(">>>>>>>>>>>>>> /dispatch-orders/%s: error getting json data", vars["oid"]))
	}
	oneliners.PrettyJson(reqBody)

	var respBody []byte
	if respBody, err = json.MarshalIndent(ResponseBody{
		OrderHash: vars["oid"],
		Status:    string(reqBody.Status),
		Message:   "Update order success !",
	}, "", "  "); err != nil {
		return mock.NewResponse(http.StatusInternalServerError,
			fmt.Sprintf(">>>>>>>>>>>>>> /dispatch-orders/%s: error marshalling json response", vars["oid"]))
	}

	fmt.Println("fmt >>>>>>>\n", string(respBody))
	oneliners.PrettyJson(respBody)
	return mock.NewResponse(http.StatusAccepted, string(respBody))
}

// RunMockCore runs a mock server of core
func RunMockCore() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/users/{uid}/orders/{oid}/dispatch", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, assignDriver(req))
	}).Methods("POST")

	router.HandleFunc("/api/v1/users/{uid}/orders/{oid}", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, assignmentError(req))
	}).Methods("POST")

	router.HandleFunc("/dispatch-orders/{oid}", func(w http.ResponseWriter, req *http.Request) {
		mock.Respond(w, updateOrder(req))
	}).Methods("PATCH")

	fmt.Println("running the mock core ...")
	log.Fatalln(http.ListenAndServe(":8001", router))
}
