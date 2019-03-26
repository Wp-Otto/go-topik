package main

import (
	"bufio"
	"fmt"
	"github.com/Wp-Otto/go-topik"
	"log"
	"os"
)

func main() {

	f, err := os.Open("testdata/domains.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(f)

	tk := topik.New(10)

	for scanner.Scan() {

		item := scanner.Text()
		//fmt.Println(item)
		tk.Insert(item)
		//fmt.Println(tk)

	}

	if err := scanner.Err(); err != nil {
		log.Println("error during scan: ", err)
	}


	fmt.Println(len(tk.Keys))
	fmt.Println(tk.Get())

}
