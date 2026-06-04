package main

import (
	// "bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("enter source path")
	// reader := bufio.NewReader(os.Stdin)
	// source, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("source read not done successfully", err)
	// }

	// fmt.Println(source)

	// data, err2 := os.ReadFile(strings.TrimSpace(source)) //using trimSpace as \n also added when hitting enter
	// if err2 != nil {
	// 	fmt.Println("problem in reading", err2)
	// }

	// fmt.Println("enter dest path")
	// dest, err3 := reader.ReadString('\n')
	// if err3 != nil {
	// 	fmt.Println("dest read not done successfully", err)
	// }
	// fmt.Println(dest)

	// err4 := os.WriteFile(strings.TrimSpace(dest), data, 0644)
	// if err4 != nil {
	// 	fmt.Println("error encountered while writing")
	// } else {
	// 	fmt.Println("file written :) ")
	// }

	// -> with Args

	args := os.Args
	// fmt.Println(args)
	if len(args) > 3 { // 3 because args[0] is by default allocated at C:\Users\Ankit\AppData\Local\go-build\fb\fb4196b5859aefc3851526e02ab7f40d67b32a995fa98279fcce1e735f60c2ea-d\main.exe
		fmt.Println("only 2 args are accepted")
	} else {
		source := args[1]
		dest := args[2]
		// fmt.Println(source,dest,strings.TrimSpace(source),strings.TrimSpace(dest))
		data, err := os.ReadFile(strings.TrimSpace(source))

		if err != nil {
			fmt.Println("error while reading ", err)
		} else {
			err2 := os.WriteFile(strings.TrimSpace(dest), data, 0644)
			if err2 != nil {
				fmt.Println("error while writing ", err2)
			} else {
				fmt.Println("filw write successfull")
			}
		}

	}
}

// algo
// source as path
// -> read the source
// -> dest as path
// -> create a file at dest
// -> write the data of source at dest
