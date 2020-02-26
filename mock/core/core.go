package mock_core

import (
    "encoding/json"
    "fmt"
    mock "github.com/shudipta/play-with-locust/mock"
    "log"
    "magic.pathao.com/food/dispatcher-v2/usecase/data"
    "magic.pathao.com/food/dispatcher-v2/usecase/services/core"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/tamalsaha/go-oneliners"
)

/*
{
    "order_id": 627949,
    "status": "success",
    "message": "task_13327397-72bb-4a2e-a0cd-713b3c731791"
}
202
*/

func allocates(req *http.Request) mock.Response {
    var reqBody data.AllocateReq
    var err error

    defer req.Body.Close()
    if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
        return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/allocates: error getting json data")
    }

    var respBody []byte
    if respBody, err = json.MarshalIndent(core.ResponseBody{
        OrderHash: "e6eb1f54-5c14-4274-a4c8-ffba17c12f81",
        Status:    "success",
        Message:   "task_13327397-72bb-4a2e-a0cd-713b3c731791",
    }, "", " "); err != nil {
        return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/allocates: error marshalling json response")
    }

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
    var reqBody data.CoreAssignedReq
    var err error

    defer req.Body.Close()
    if err = json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
        return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/cancel: error getting json data")
    }
    oneliners.PrettyJson(reqBody)

    var respBody []byte
    if respBody, err = json.MarshalIndent(core.ResponseBody{
        OrderHash: "e6eb1f54-5c14-4274-a4c8-ffba17c12f81",
        Status:    "CANCELED",
        Message:   "Assignment failed",
    }, "", " "); err != nil {
        return mock.NewResponse(http.StatusBadRequest, ">>>>>>>>>>>>>> /v1/cancel: error marshalling json response")
    }

    return mock.NewResponse(http.StatusAccepted, string(respBody))
}

func RunMockCore() {
    router := mux.NewRouter()

    router.HandleFunc("/user/{uid}/order/{oid}", func (w http.ResponseWriter, req *http.Request) {
        mock.Respond(w, assignmentError(req))
    }).Methods("POST")

    fmt.Println("running the mock core ...")
    log.Fatalln(http.ListenAndServe(":8001", router))
}