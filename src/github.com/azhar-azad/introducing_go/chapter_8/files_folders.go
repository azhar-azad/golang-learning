package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func filesFoldersDemo() {

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Can't open file")
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Can't stat the file")
		return
	}

	// read the file
	bs := make([]byte, int(stat.Size()))
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("Can't read file")
		return
	}

	str := string(bs)
	fmt.Println(str)

	// shorter way
	bs, err = ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Can't read file")
		return
	}

	str = string(bs)
	fmt.Println(str)
}
