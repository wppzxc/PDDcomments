package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGetCommentResult(t *testing.T) {
	rand.Seed(time.Now().Unix())
	AK = "J34XHETM2MM4NMHI2IGJSIBGOBULQYWC3ALGT246AZWWKT5W3RPQ1013f08"
	result := getCommentResult("9397982", "150", "", "", "")
	fmt.Println(result)
}
