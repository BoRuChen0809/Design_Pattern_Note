package main

import "fmt"

const (
	homeJerseyType = "hJersey"
	awayJerseyType = "aJersey"
)

type jerseyFactory struct {
	jerseymap map[string]jersey
}

var (
	jerseyFactorySingleInstance = &jerseyFactory{make(map[string]jersey)}
)

func getJerseyFactory() *jerseyFactory {
	return jerseyFactorySingleInstance
}

func (j *jerseyFactory) getJerseyByType(jerseytype string) (jersey, error) {
	if j.jerseymap[jerseytype] != nil {
		return j.jerseymap[jerseytype], nil
	}

	if jerseytype == homeJerseyType {
		j.jerseymap[jerseytype] = newHomeJersey()
		return j.jerseymap[jerseytype], nil
	}

	if jerseytype == awayJerseyType {
		j.jerseymap[jerseytype] = newAwayJersey()
		return j.jerseymap[jerseytype], nil
	}

	return nil, fmt.Errorf("wrong jersey type.")
}
