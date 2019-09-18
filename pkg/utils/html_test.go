package utils

import (
	"fmt"
	"testing"
)

func TestGetGoodsPrice(t *testing.T) {
	itemId := "2518027055"
	str := "1234567890.abcxyz"
	fmt.Println([]byte(str))
	price := GetGoodsPrice(itemId, "0.5")
	fmt.Println(price)
}
