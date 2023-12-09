# Med Location

This project is a Go application that serves as an API for publishing and subscribing to ambulance locations through a pubsub/redis and server-side events architecture.

# Usage

To start the API server, navigate to the project root directory and run the following command:

make run

Once the server is up and running, you can perform the following actions:
Publish Ambulance Locations

To publish ambulance locations, send a POST request to the following endpoint:

http://localhost:8080/publish

Include the necessary parameters in the request body to specify the ambulance location. For example:

json

{
   "type": "Point",
   "coordinates": [longitude, latitude]
}

Replace longitude and latitude with the actual coordinates.
Subscribe to Ambulance Channels

To subscribe to specific ambulance channels, make a GET request to the following endpoint:

bash

http://localhost:8080/subscribe/{channel}

Replace {channel} with the desired channel name. This endpoint will establish a server-sent event (SSE) connection, allowing you to receive real-time updates on ambulance locations for the specified channel.
Example

Here's an example of how to use the API:

    Start the server:

bash

make run

    Publish ambulance location:

Send a POST request to http://localhost:8080/publish with a JSON body containing the location:

json

{
   "type": "Point",
   "coordinates": [longitude, latitude]
}

    Subscribe to a channel:

Make a GET request to http://localhost:8080/subscribe/{channel} to receive real-time updates for the specified channel.

Feel free to explore and integrate this API into your application for efficient ambulance location tracking.
