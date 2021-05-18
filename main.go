package main

import (
	"log"
	"time"
)

func main(){
	ti := time.NewTicker(4 * time.Second)
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("starting")
	for {
		<-ti.C
		log.Println("tick")
	}

}
