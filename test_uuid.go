package main

import (
	"github.com/google/uuid" 
	"fmt"
)

func IsUUID(str string) error {
	_, err := uuid.Parse(str)
	if err != nil {
		return err
	}
	return nil
}

func main(){
	err := IsUUID("abcd$qdw")
	if err != nil {
		fmt.Println("%v\n", err)
	} else {
		fmt.Println("success")
	}

	err = IsUUID("77eddd6c-d8d8-c046-efcd-513ae5013ccd")
	if err != nil {
		fmt.Println("%v\n", err)
	} else {
		fmt.Println("success")
	}
}
