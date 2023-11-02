package main

import (
	"fmt"
	"log"
	"os"

	"github.com/psanford/aws-account-id-from-key/awskey"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <key>", os.Args[0])
	}

	for _, k := range os.Args[1:] {
		key, err := awskey.Decode(k)
		if err != nil {
			log.Fatalf("Decode %s err: %s", k, err)
		}
		fmt.Println(key.AccountID)
	}
}
