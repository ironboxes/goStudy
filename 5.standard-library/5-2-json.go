package main

import (
	"encoding/json"
	"fmt"
)

type ResponseData struct {
	Data struct {
		Items []Body `json:"items"`
		TotalCount int64 `json:"total_count"`
	} `json:"data"`
	Message    string `json:"message"`
	ResultCode int    `json:"result_code"`
}

type Body struct {
	ID int `json:"_id"`
}

func main() {
	//test()
	test1()
}

func test() {
	jsonStr := `{"data":{"items":[{"_id":2}],"total_count":1},"message":"","result_code":200}`
	var responseData ResponseData
	err := json.Unmarshal([]byte(jsonStr), &responseData)
	if err != nil {
		fmt.Println("parseJson error:"+err.Error())
		return
	}
	fmt.Println(responseData)
}

func test1() {
	r := ResponseData{
		Data:struct{
			Items []Body `json:"items"`
			TotalCount int64 `json:"total_count"`
		}{
			Items: []Body {
				{ID: 1},
				{ID: 2},
			},
			TotalCount: 1,
		},
		Message: "",
		ResultCode: 200,
	}
	resBytes, err := json.Marshal(r)
	if err != nil {
		fmt.Println("convertJson error: " + err.Error())
	}
	fmt.Println(string(resBytes))
}