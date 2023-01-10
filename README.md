# Flight Path Tracker
The Flight Path Tracker is a tool for finding the final itinerary for a set of flights. It uses a topological sort algorithm to find a valid flight path from the given flights.

## Getting Started
To use the Flight Path Tracker, import the package and create a new FlightPathTracker struct. Then call the FindItineraryStartAndEnd method on the struct, passing in a 2D slice of strings representing the flights. The method will return a slice of strings representing the airports visited in order if a valid itinerary is found, or an empty slice if no valid itinerary is found.


## Installation

### Prerequisites
Go 1.15+

To install the Flight Path Tracker server, run the following command:

```shell
go get github.com/akhilesharora/flightpath
```


## Usage
To start the dockerized server, run the following command:

```shell
make docker-run
```
This will start the server listening on port 8080.

To query the server, send a POST request to the /calculate endpoint with a JSON array of flights in the request body. Each flight should be represented as a JSON array with two strings, the starting airport and the ending airport.

For example, to find the start and end airports for the flights [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]], send the following request:

## API

### POST /calculate
Find the start and end airports for the given flights.

Request

```http request
curl -d '{"flights": [["SFO", "HKO"], ["YYZ", "SFO"], ["YUL", "YYZ"], ["HKO", "ORD"]]}' -H "Content-Type: application/json" -X POST http://localhost:8080/calculate
```

The server will respond with a JSON array containing the start and end airports:


```http request
// The response will be a JSON array in the format ["<starting_airport>", "<ending_airport>"]
{"path":["SFO","EWR"]}
```

## Example
To use the Flight Path Tracker service, import the flightpath package and create a new FlightPathTracker struct. Then, call the FindItineraryStartAndEnd method on the struct, passing in a 2D slice of strings representing the tickets. The method will return a slice of strings representing the starting and ending airports of the flight path, or an empty slice if no valid itinerary could be found.

```go
package main

import (
	"fmt"

	flight "github.com/akhilesharora/flightpath/internal/flightpath"
)

func main() {
	f := &flight.FlightPathTracker{}
	tickets := [][]string{
		{"SFO", "HKO"},
		{"YYZ", "SFO"},
		{"YUL", "YYZ"},
		{"HKO", "ORD"},
	}
	itinerary := f.FindItineraryStartAndEnd(tickets)
	fmt.Println(itinerary)
	// Output: ["YUL", "ORD"]
}

```
## Testing
To run the tests for the Flight Path Tracker, navigate to the root of the project and run the following command:

```shell
make test
```

## Benchmarking
To run the benchmarks for the Flight Path Tracker, navigate to the root of the project and run the following command:

```shell
make benchmarks
```

## Complexity
The topological sort algorithm used in the findItinerary function has a time complexity of O(V+E), where V is the number of vertices (airports) and E is the number of edges (flights) in the graph. The space complexity is also O(V+E) as the graph is stored in the form of adjacency lists. The findStartingAirport function has a time complexity of O(E), as it iterates through all the tickets and counts the incoming flights for each airport. The space complexity is also O(E) as it stores the incoming flight counts for each airport in a map. The FindItineraryStartAndEnd function has a time and space complexity of O(1), as it simply retrieves the first and last elements of the itinerary slice.

Overall, the performance of this service is efficient for large inputs, with a linear time and space complexity.

## Future Plans
1. Add request tracing: By adding request tracing, you can track the flow of requests through the microservice and identify any bottlenecks or errors. This can be done using a distributed tracing tool like Zipkin or Jaeger.
2. Use a message queue: Currently, the microservice receives a list of flights and calculates the flight path in one request. If the number of flights is large, this could lead to long request times. To improve this, you could use a message queue like RabbitMQ or Kafka to process the flights asynchronously.
3. Use a cache: If the flight path calculations are computationally expensive, you could use a cache like Redis to store the results of previous calculations and serve them quickly for subsequent requests.
4. Use a load balancer: If the microservice is expected to receive a high volume of requests, it might be a good idea to use a load balancer like NGINX or HAProxy to distribute the requests across multiple instances of the microservice.
5. Improve error handling: Currently, the microservice returns a generic "internal server error" message if there is an error while processing the request. It would be a good idea to improve error handling by adding more specific error messages and logging the errors to help with debugging.
