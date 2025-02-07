package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

var producerClient *azeventhubs.ProducerClient = nil
var LOG_EVENTS_ACTIVE = os.Getenv("LOG_EVENTS_ACTIVE")

func checkEnvironmentVars() (string, string) {
	EVENTHUB_CONNECTION_STRING, isavailable := os.LookupEnv("EVENT_HUB_CONNECTION_STRING")

	if !isavailable {
		log.Printf("Environment Variable Not Declared [EVENTHUB_CONNECTION_STRING] = [%s] }", EVENTHUB_CONNECTION_STRING)
	}

	EVENTHUB_NAME, isavailable := os.LookupEnv("EVENTHUB_NAME")

	if !isavailable {
		log.Printf("Environment Variable [EVENTHUB_NAME] = [%s] }", EVENTHUB_NAME)
	}
	return EVENTHUB_CONNECTION_STRING, EVENTHUB_NAME
}

func sendDataToEventHub(producerClient *azeventhubs.ProducerClient, bytedata []byte) {
	batch, err := producerClient.NewEventDataBatch(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	// See ExampleProducerClient_AddEventData for more information.
	err = batch.AddEventData(&azeventhubs.EventData{Body: bytedata}, nil)

	if err != nil {
		panic(err)
	}

	err = producerClient.SendEventDataBatch(context.TODO(), batch, nil)

	if err != nil {
		panic(err)
	}
}

func eventHandler(w http.ResponseWriter, r *http.Request) {
	EVENTHUB_CONNECTION_STRING, EVENTHUB_NAME := checkEnvironmentVars()

	if EVENTHUB_CONNECTION_STRING == "" || EVENTHUB_NAME == "" {
		panic("Environment Variables not defined")
	}

	producerClient, err := azeventhubs.NewProducerClientFromConnectionString(EVENTHUB_CONNECTION_STRING, EVENTHUB_NAME, nil)

	if err != nil {
		log.Printf("Environment Variable [EVENTHUB_NAME] = [%s] }", EVENTHUB_NAME)
		panic(err)
	}

	defer producerClient.Close(context.TODO())

	bytedata, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	sendDataToEventHub(producerClient, bytedata)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/HttpBrokerTrigger", eventHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
