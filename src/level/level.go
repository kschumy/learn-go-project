package level

import (
	"space"
	"spacetypes"
	"errors"
)

type Level struct {
	numOfSpaces int // HACK: currently needed for faster check for hasAvailableSpaces. Find better way
	unavailableSpaces []*space.Space
	availableSpaces *map[int]*space.SpacesQueue
}

func (currLevel *Level) hasAvailableSpaces() bool {
	return currLevel.numOfSpaces != len(currLevel.unavailableSpaces)
}

func (currLevel *Level) hasAvailableSpacesOfType(spaceType int) bool {
	return currLevel.numOfSpaces != len(currLevel.unavailableSpaces)
}

// TODO: combine logic with CreateLevel and delete this?
// Returns true if one of the provided numsOfSpaces is greater than 0 and none of the numsOfSpaces
// are below 0. Otherwise, returns false
func hasValidNumOfSpacesForNewLevel(numOfSpacesByType map[spacetypes.SpaceType]int) bool {
	var hasAtLeastOneSpace bool
	for _, num := range numOfSpacesByType {
		if num < 0 {
			return false
		}
		if !hasAtLeastOneSpace {
			hasAtLeastOneSpace = true
		}
	}
	return hasAtLeastOneSpace
}

func CreateLevel(numOfSpacesByType map[spacetypes.SpaceType]int) (*Level, error) {
	if !hasValidNumOfSpacesForNewLevel(numOfSpacesByType) {
		return nil, errors.New("invalid number of parking spaces")
	}
	newLevel := &Level{availableSpaces: make(&map[int]*space.SpacesQueue),}
	//newLevel.availableSpaces

	for spaceType, num := range numOfSpacesByType {

	}
}