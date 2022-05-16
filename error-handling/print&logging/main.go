// more examples in: https://github.com/GoesToEleven/go-programming

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hi :3")

	// first example ==================
	// _, err := os.Open("no-file.txt")
	// if err != nil {
	// 	//fmt.Println("err happened", err)
	// 	log.Println("err happened", err)
	// 	// log.Fatalln(err)
	// 	// panic(err)
	// }
	// ================================

	// second example =================
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
		// log.Fatalln(err)
		// panic(err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")

}
