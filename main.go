package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kshard/fvecs"
)

func main() {

	file := ""
	result := ""

	fmt.Print("Enter path to file: ")

	// Reading vectors
	scanAmount, err := fmt.Scan(&file)
	if err != nil || scanAmount != 1 {
		fmt.Println("Console read error: ", err)
		return
	}

	r, err := os.Open(file)
	if err != nil {
		fmt.Println("File open error: ", err)
		return
	}

	if strings.Contains(file, ".fvecs") {
		decoder := fvecs.NewDecoder[float32](r)
		for {
			vectorData, err := decoder.Read()
			if err != nil {
				break
			}
			result += fmt.Sprintf("%f, ", vectorData)
		}
	} else if strings.Contains(file, ".ivecs") {
		decoder := fvecs.NewDecoder[uint32](r)
		for {
			vectorData, err := decoder.Read()
			if err != nil {
				break
			}
			result += fmt.Sprintf("%d, ", vectorData)
		}
	} else if strings.Contains(file, ".bvecs") {
		decoder := fvecs.NewDecoder[byte](r)
		for {
			vectorData, err := decoder.Read()
			if err != nil {
				break
			}
			result += fmt.Sprintf("%c, ", vectorData)
		}
	} else {
		fmt.Println("File format not supported")
		return
	}

	result = result[:len(result)-2] //trim last comma

	os.WriteFile("result.txt", []byte(result), 0644)
}
