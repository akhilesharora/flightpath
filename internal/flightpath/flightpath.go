package flightpath

import (
	"container/list"
)

type Tracker struct{}

type TrackerSrv interface {
	findStartingAirport(tickets [][]string) string
	findItinerary(tickets [][]string) []string
	FindItineraryStartAndEnd(tickets [][]string) []string
}

// findStartingAirport finds the starting airport for the given
// tickets by counting the incoming flights for each airport and
// returning the first airport that has no incoming flights.
func (f *Tracker) findStartingAirport(tickets [][]string) string {
	// Map to store the incoming flight counts for each airport
	incomingCounts := make(map[string]int)
	// Iterate through the tickets and count the incoming flights
	for _, ticket := range tickets {
		incomingCounts[ticket[1]]++
	}
	// Iterate through the tickets again and find the starting airport
	for _, ticket := range tickets {
		if incomingCounts[ticket[0]] == 0 {
			return ticket[0]
		}
	}
	// This shouldn't happen since there must be a starting airport
	return ""
}

// findItinerary finds a valid itinerary for the given tickets
// using a topological sort algorithm. It returns a slice
// of strings representing the airports visited in order.
func (f *Tracker) findItinerary(tickets [][]string) []string {
	// Find the starting airport
	startingAirport := f.findStartingAirport(tickets)

	// Build the directed graph of airports and tickets.
	incomingCounts := make(map[string]int)
	adjList := make(map[string][]string)
	for _, ticket := range tickets {
		incomingCounts[ticket[1]]++
		adjList[ticket[0]] = append(adjList[ticket[0]], ticket[1])
	}

	// Perform a topological sort to find a valid itinerary.
	var itinerary []string
	queue := list.New()
	queue.PushBack(startingAirport)
	for queue.Len() > 0 {
		from := queue.Remove(queue.Front()).(string)
		itinerary = append(itinerary, from)
		for _, to := range adjList[from] {
			incomingCounts[to]--
			if incomingCounts[to] == 0 {
				queue.PushBack(to)
			}
		}
	}

	// Check if there is a valid itinerary.
	if len(itinerary) != len(tickets)+1 {
		return []string{}
	}
	return itinerary
}

// FindItineraryStartAndEnd finds the final itinerary for the given tickets
// by calling the findItinerary method and returning the starting and
// ending airports if a valid itinerary is found, or an empty slice
// if no valid itinerary is found.
func (f *Tracker) FindItineraryStartAndEnd(tickets [][]string) []string {
	itinerary := f.findItinerary(tickets)
	if len(itinerary) == 0 {
		return []string{}
	}
	src, dest := itinerary[0], itinerary[len(itinerary)-1]
	return []string{src, dest}
}
