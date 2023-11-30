package main

import "log"

func main() {
	// Main function of the Go program
	//
	// Summary:
	// This code snippet shows the main function of a Go program. It initializes variables for the server address, Redis address, password, database, and a Redis pubsub object. It then creates a new API server instance with the server address and Redis pubsub object. Finally, it runs the server and logs any errors.
	//
	// Example Usage:
	// 	listenAddr := ":3000"
	// 	redis_addr := "localhost:6379"
	// 	redis_pwd := ""
	// 	redis_db := 0
	// 	redis_pubsub := NewRedisPubSub(redis_addr, redis_pwd, redis_db)
	//
	// 	server := NewAPIServer(listenAddr, redis_pubsub)
	//
	// 	if err := server.Run(); err != nil {
	// 		log.Fatal(err)
	// 	}
	//
	// Code Analysis:
	// Inputs:
	// - listenAddr (string): The address on which the server should listen for incoming requests.
	// - redis_addr (string): The address of the Redis server.
	// - redis_pwd (string): The password for the Redis server.
	// - redis_db (int): The database number to use in the Redis server.
	//
	// Flow:
	// 1. Initialize the listenAddr variable with the value ":3000".
	// 2. Initialize the redis_addr variable with the value "localhost:6379".
	// 3. Initialize the redis_pwd variable with an empty string.
	// 4. Initialize the redis_db variable with the value 0.
	// 5. Create a new Redis pubsub object by calling the NewRedisPubSub function with the redis_addr, redis_pwd, and redis_db variables.
	// 6. Create a new API server object by calling the NewAPIServer function with the listenAddr and redis_pubsub variables.
	// 7. Run the server by calling the Run method on the server object.
	// 8. If an error occurs during server execution, log the error and exit the program.
	//
	// Outputs:
	// None.

	listenAddr := ":3000"
	redis_addr := "localhost:6379"
	redis_pwd := ""
	redis_db := 0
	redis_pubsub := NewRedisPubSub(redis_addr, redis_pwd, redis_db)

	server := NewAPIServer(listenAddr, redis_pubsub)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
