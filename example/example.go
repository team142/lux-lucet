package main

import (
	"encoding/json"
	"errors"
	"github.com/team142/lux-lucet/lulu"
	"log"
)

func main() {

	//Start a Health Server
	healthServer := lulu.StartHealthServer()

	//Set initial state for each subsystem
	healthServer.UpdateOk("net/io")
	healthServer.UpdateOk("disk/io")
	healthServer.UpdateOk("queue-handler")

	//Query state
	state := healthServer.Query()
	//Marshal to json and log json string
	b, err := json.Marshal(state)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))

	//Pretend to run some subsystem
	runQueueHandler(healthServer)

	//Get the state and log the json string
	state = healthServer.Query()
	b, err = json.Marshal(state)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(b))

	//Blocking call - do last or `go lulu.StartRestServer...`
	lulu.StartRestServer(":9001", healthServer)

}

//Imaginary subsystem
func runQueueHandler(healthServer *lulu.HealthServer) {

	// ...
	// ...
	// ...

	//Some work goes wrong here
	err := someWork()
	if err != nil {
		//Example of an update to the system
		healthServer.Update("queue-handler", false, err.Error())
		return
	}

}

func someWork() error {
	return errors.New("some error")
}
