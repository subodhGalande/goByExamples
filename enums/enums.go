package main

import "fmt"

type ServerState int

const (
	stateIdle ServerState = iota //iota is a special keyword which generates successive values automatically like 0,1,2

	stateConnected
	stateError
	stateRetrying
)

var stateName = map[ServerState]string{
	stateIdle:      "idle",
	stateConnected: "connected",
	stateError:     "error",
	stateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {

	ns := transition(stateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

}

func transition(s ServerState) ServerState {
	switch s {
	case stateIdle:
		return stateConnected
	case stateConnected, stateRetrying:
		return stateIdle
	case stateError:
		return stateError
	default:
		panic(fmt.Errorf("unknow error &s", s))
	}
}
