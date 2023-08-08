package dao

import (
	"fmt"
	"testing"
)

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本书，%v\n", k, v)
	}
}
