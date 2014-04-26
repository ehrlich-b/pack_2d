package main

import (
	"fmt"
	"github.com/adotout/pack_2d"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	packer := pack_2d.Packer2d{}
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: pack_run [file name]")
		return
	}

	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	blocks := strings.TrimSpace(string(b))

	blocksSplit := strings.Split(blocks, " ")

	for _, value := range blocksSplit {
		xySplit := strings.Split(value, ",")
		width, _ := strconv.ParseInt(xySplit[0], 10, 0)
		height, _ := strconv.ParseInt(xySplit[1], 10, 0)
		packer.AddNewBlock(int(width), int(height), 0)
	}
	rBlocks, width, height := packer.Pack()
	pack_2d.PrintBlocks(rBlocks)
	fmt.Print("Width: ")
	fmt.Println(width)
	fmt.Print("Height: ")
	fmt.Println(height)
}
