package tracker

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestFindStartingAirport(t *testing.T) {
	tests := []struct {
		name     string
		tickets  [][]string
		expected string
	}{
		{
			name: "single ticket",
			tickets: [][]string{
				{"SFO", "HKO"},
			},
			expected: "SFO",
		},
		{
			name: "multiple tickets",
			tickets: [][]string{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
				{"YUL", "YYZ"},
				{"HKO", "ORD"},
			},
			expected: "YUL",
		},
		{
			name: "multiple tickets with cycle",
			tickets: [][]string{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
				{"YUL", "YYZ"},
				{"HKO", "ORD"},
				{"ORD", "YUL"},
			},
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Tracker{}
			got := f.findStartingAirport(test.tickets)
			if got != test.expected {
				t.Errorf("unexpected starting airport: got %q, want %q", got, test.expected)
			}
		})
	}
}

func TestFindItineraryStartAndEnd(t *testing.T) {
	// Create a new FlightPathTracker instance.
	f := &Tracker{}

	// Define the test cases.
	testCases := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		{
			name: "single ticket",
			tickets: [][]string{
				{"SFO", "HKO"},
			},
			expected: []string{"SFO", "HKO"},
		},
		{
			name: "two tickets",
			tickets: [][]string{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
			},
			expected: []string{"YYZ", "HKO"},
		},
		{
			name: "three tickets",
			tickets: [][]string{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
				{"YUL", "YYZ"},
				{"HKO", "ORD"},
			},
			expected: []string{"YUL", "ORD"},
		},
	}

	// Iterate through the test cases.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the FindItineraryStartAndEnd method.
			result := f.FindItineraryStartAndEnd(tc.tickets)

			// Check if the result is what we expected.
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestFindItinerary(t *testing.T) {
	// Define the test cases.
	testCases := []struct {
		name     string
		tickets  [][]string
		expected []string
	}{
		{
			name: "single ticket",
			tickets: [][]string{
				{"SFO", "HKO"},
			},
			expected: []string{"SFO", "HKO"},
		},
		{
			name: "two tickets",
			tickets: [][]string{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
			},
			expected: []string{"YYZ", "SFO", "HKO"},
		},
		{
			name: "three tickets",
			tickets: [][]string{
				{"SFO", "HKO"},
				{"YYZ", "SFO"},
				{"YUL", "YYZ"},
				{"HKO", "ORD"},
			},
			expected: []string{"YUL", "YYZ", "SFO", "HKO", "ORD"},
		},
	}

	// Create a FlightPathTracker.
	f := &Tracker{}

	// Run the tests.
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := f.findItinerary(tc.tickets)
			// Check if the result is what we expected.
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func BenchmarkFindItineraryStartAndEnd(b *testing.B) {
	// Generate 100000 test tickets.
	tickets := make([][]string, 100000)
	for i := range tickets {
		tickets[i] = []string{strconv.Itoa(i), strconv.Itoa(i + 1)}
	}

	// Create a FlightPathTracker.
	f := Tracker{}

	// Run the benchmark.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.FindItineraryStartAndEnd(tickets)
	}
}

func BenchmarkFindItinerary(b *testing.B) {
	// Create a slice of tickets with 100000 elements.
	tickets := make([][]string, 100000)
	for i := 0; i < 100000; i++ {
		// Add a dummy ticket to the slice.
		tickets[i] = []string{"SFO", "HKO"}
	}

	// Create a FlightPathTracker object.
	f := &Tracker{}

	// Benchmark the findItinerary function.
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.findItinerary(tickets)
	}
}

func BenchmarkFindStartingAirport(b *testing.B) {
	// Generate 100000 tickets
	tickets := make([][]string, 100000)
	for i := 0; i < 100000; i++ {
		tickets[i] = []string{fmt.Sprintf("A%d", i), fmt.Sprintf("B%d", i)}
	}
	// Add a cycle to the tickets
	tickets[99998][1] = "A0"
	tickets[99999][1] = "A1"
	tickets[0][1] = "A99998"
	tickets[1][1] = "A99999"

	b.ResetTimer()
	f := &Tracker{}
	for i := 0; i < b.N; i++ {
		f.findStartingAirport(tickets)
	}
}
