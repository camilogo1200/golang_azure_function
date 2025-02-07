package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
)

var EVENTHUB_NAMESPACE = os.Getenv("EVENTHUB_NAMESPACE")
var EVENTHUB_NAME = os.Getenv("EVENTHUB_NAME")
var EVENTHUB_CONNECTION_STRING = os.Getenv("EVENT_HUB_CONNECTION_STRING")

func helloHandler(w http.ResponseWriter, r *http.Request) {

	//eventHubNamespace := EVENTHUB_NAME
	eventHubName := EVENTHUB_NAME

	defaultAzureCred, err := azidentity.NewDefaultAzureCredential(nil)

	if err != nil {
		panic(err)
	}

	connectionString := EVENTHUB_CONNECTION_STRING

	producerClient, err := azeventhubs.NewProducerClientFromConnectionString(connectionString, eventHubName, nil)

	if err != nil {
		panic(err)
	}

	defer producerClient.Close(context.TODO())
	err = batch.AddEventData(events[i], nil)
	producerClient.SendEventDataBatch(context.TODO(), batch, nil)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/HttpBrokerTrigger", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
