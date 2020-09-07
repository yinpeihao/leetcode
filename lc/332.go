package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
	//fmt.Println(findItinerary([][]string{{"JFK", "ATL"}, {"ATL", "JFK"},{"JFK", "ATL"}, {"ATL", "JFK"}}))
}

var entryMap = map[string][]string{}
var ticketLen int

func findItinerary(tickets [][]string) []string {
	ticketLen = len(tickets) + 1
	for _, ticket := range tickets {
		entryMap[ticket[0]] = append(entryMap[ticket[0]], ticket[1])
	}
	for k, v := range entryMap {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		entryMap[k] = v
	}
	return dfs([]string{"JFK"}, "JFK")
}

func dfs(journey []string, prev string) []string {
	//fmt.Println(journey)
	//time.Sleep(time.Millisecond*100)
	if len(journey) == ticketLen {
		return journey
	}
	for i, next := range entryMap[prev] {
		if next != "" {
			entryMap[prev][i] = ""
			x := dfs(append(journey, next), next)
			if len(x) == ticketLen {
				return x
			}
			entryMap[prev][i] = next
		}
	}
	return journey
}
