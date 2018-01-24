package main

import (
	"log"

	a "github.com/dqian96/alexandria/archive"
	d "github.com/dqian96/alexandria/director"
	"github.com/dqian96/alexandria/utils"

	"os"

	server "github.com/dqian96/alexandria/server"
)

const (
	defaultCustomerServer   = ":8080"
	defaultDirectorPort     = ":8081"
	heartBeatTimeoutSeconds = 180
)

func main() {
	// args, port := os.Args, defaultPort
	// log.Println("Command line arguments: ", args)
	// if len(args) > 1 {
	// 	port = ":" + args[1]
	// }

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Print("Starting Alexandria!")
	log.Print("Received command line arguments: ", os.Args)

	log.Print("Creating Archive")
	a, err := a.NewArchive(10, 10, 10, 1) // returns a pointer to archive struct, which is an Archive interface
	if err != nil {
		log.Fatalf("Fail to create Archive: %v", err)
	}

	log.Printf("Starting Director server on port %s", defaultDirectorPort)
	d := d.NewDirector(make([]string, 0, 0), heartBeatTimeoutSeconds, "how to register?", a)
	go func() {
		if err := d.Serve(defaultDirectorPort); err != nil {
			utils.Error("").Fatalf("Fail to start Director: %v", err)
		}
	}()

	log.Printf("Starting customer server on port %s", defaultCustomerServer)
	s := server.NewServer(d)
	if err := s.Serve(defaultCustomerServer); err != nil {
		utils.Error("").Fatalf("Fail to start the customer server: %v", err)
	}

}
