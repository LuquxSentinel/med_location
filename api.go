package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sentinel/med_location/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

type apiServer struct {
	listenAddr   string
	router       *mux.Router
	pubsubclient PubSub
}

func NewAPIServer(listenAddr string, pubsubclient PubSub) *apiServer {
	return &apiServer{
		listenAddr:   listenAddr,
		router:       mux.NewRouter(),
		pubsubclient: pubsubclient,
	}
}

// Run starts the API server and sets up the routes for handling HTTP requests.
//
// Example Usage:
//
//	api := NewAPIServer(":8080", pubsubclient)
//	err := api.Run()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Inputs:
//
//	None
//
// Outputs:
//
//	An error if there is an issue starting the server.
func (api *apiServer) Run() error {

	api.router.HandleFunc("/publish", api.handleFunc(api.publish)).Methods("GET")
	api.router.HandleFunc("/subscribe/{channel}", api.handleFunc(api.subscribe)).Methods("GET")

	return http.ListenAndServe(api.listenAddr, api.router)
}

// handleFunc is a method that belongs to the apiServer struct. It takes an APIFunc function as a parameter and returns an http.HandlerFunc.
// The method creates a new context and returns an anonymous function that handles the HTTP request.
// Inside the anonymous function, it calls the APIFunc function with the provided context, http.ResponseWriter, and *http.Request parameters.
// If an error occurs, it calls the InternalServerError function to write an error response.
func (api *apiServer) handleFunc(fn APIFunc) http.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(ctx, w, r)
		if err != nil {
			// write error response on error
			InternalServerError(w, map[string]string{"error": err.Error()})
		}

	}
}

// publish handles an HTTP request by decoding the request body, publishing the message to a channel/topic, and writing a success response.
func (api *apiServer) publish(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// fetch message from request
	// decode request body
	msg := new(types.Message)
	if err := DecodeRequest(r.Body, msg); err != nil {
		return err
	}

	// write location to channel/topic
	data, _ := EncodeRequest(msg.Location)
	err := api.pubsubclient.Publish(ctx, msg.Channel, data)
	if err != nil {
		return err
	}
	// write response on success
	WriteResponse(w, map[string]string{"msg": "success"})
	return nil
}

// subscribe handles HTTP requests for subscribing to a specific channel.
// It sets the necessary headers for server-sent events and establishes a connection with the client.
// It then continuously receives messages from the pubsub client and writes the payload of each message to the response writer.
//
// Parameters:
// - ctx: The context.Context object for managing the lifecycle of the request.
// - w: The http.ResponseWriter object for writing the response.
// - r: The *http.Request object representing the incoming request.
//
// Returns:
// - An error if there is an issue with the channel or writing the response.
func (api *apiServer) subscribe(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// ctx := context.Background()

	channel := mux.Vars(r)["channel"]
	if channel == "" {
		return fmt.Errorf("invalid channel provided")
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// TODO : close channel

	flusher, _ := w.(http.Flusher)
	pubsub := api.pubsubclient.Subscribe(ctx, channel)
	for {
		location, err := pubsub.ReceiveMessage(ctx)
		fmt.Printf("%+v\n", location)
		if err != nil {
			log.Println(err.Error())
		}
		err = WriteResponse(w, location.Payload)
		if err != nil {
			log.Println(err)
		}
		// log.Println(write)
		flusher.Flush()
	}

}
