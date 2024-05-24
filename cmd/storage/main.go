package main

import (
	"fmt"

	"github.com/buzruk/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test.txt", []byte("test"))

	if err != nil {
		panic(err)
		//log.Fatal(err)
	}

	//fmt.Println(st, file)

	receivedFile, err := st.GetById(file.ID)
	if err != nil {
		panic(err)
		//log.Fatal(err)
	}

	fmt.Println(receivedFile.String())

}
