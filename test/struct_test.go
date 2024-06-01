package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type ApiResponse struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func TestStruct(t *testing.T) {
	fmt.Println(unsafe.Alignof(ApiResponse{}))
	fmt.Println(unsafe.Sizeof(ApiResponse{}))
	s := ApiResponse{}
	fmt.Println(unsafe.Sizeof(s.Success))
	fmt.Println(unsafe.Sizeof(s.Code))
	fmt.Println(unsafe.Sizeof(s.Message))
	fmt.Println(unsafe.Sizeof(s.Data))
}
