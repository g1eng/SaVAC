package main

import (
	"log"
	"os"

	"github.com/g1eng/savac/testutil/fake_vps"
)

func main() {
	log.Println("mock server starting...")
	err := fake_vps.StartFakeServer(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
}
