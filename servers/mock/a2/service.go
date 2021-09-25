package mocka2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
{
    "order_id": 627949,
    "status": "success",
    "message": "task_13327397-72bb-4a2e-a0cd-713b3c731791"
}
202
*/

func callback(orderID string, userID int, exdrivers ...int) {
	offset, _ := strconv.Atoi(strings.Split(orderID, "-")[1])
	reqBody := CallbackReq{
		DriverID: intPtr(getUserOrDriver(offset, exdrivers...)),
		Location: struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		}{20.2345, 23.2353},
		UserID:  userID,
		OrderID: orderID,
	}
	jsonByte, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>>>> failed to mashal callback req body to json: ", err)
		return
	}
	fmt.Println(">>>>>>>>>>>>>>>> callback req body: ", string(jsonByte))

	body := bytes.NewReader(jsonByte)
	resp, err := http.Post("http://localhost:4200/v1/callback", "application/json", body)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>> ", fmt.Errorf("failed to request callback to dispatcher: %v", err))
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>> ", fmt.Errorf("failed to read the callback response body: %v", err))
		return
	}

	fmt.Printf(">>>>>>>>>>>> callback resp status code: %v\n", resp.StatusCode)
	fmt.Printf(">>>>>>>>>> callback response body: %v\n", string(bodyBytes))
	return
}

func timeout(orderID string, timeout time.Duration) {
	reqBody := TimeoutReq{
		OrderID: orderID,
		Timeout: timeout,
	}
	jsonByte, err := json.MarshalIndent(reqBody, "", "  ")
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>>>> failed to mashal timeout req body to json: ", err)
		return
	}
	fmt.Println(">>>>>>>>>>>>>>>> timeout req body: ", string(jsonByte))

	body := bytes.NewReader(jsonByte)
	resp, err := http.Post("http://localhost:4200/v1/timeout", "application/json", body)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>> ", fmt.Errorf("failed to request timeout to dispatcher: %v", err))
		return
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(">>>>>>>>>>>>>> ", fmt.Errorf("failed to read the timeout response body: %v", err))
		return
	}

	fmt.Printf(">>>>>>>>>>>> timeout resp status code: %v\n", resp.StatusCode)
	fmt.Printf(">>>>>>>>>> timeout response body: %v\n", string(bodyBytes))
	return
}
