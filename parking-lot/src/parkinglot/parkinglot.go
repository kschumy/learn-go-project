// NOTE: ISSUES
//	1. large dependency on special and pass (booleans) throughout program and packages
//	2. concurrency problems. Need to learn about concurrency and resolve these problems
//	3. made a lot of assumptions along the way
//	4. pretty sure the linkedlist approach isn't the best for the
//	5. consider the use of Timer to resolve problem with expired passes. This solution should
// 		include channels. See docs: https://golang.org/pkg/time/#Timer
// 6. Look into time package having different time for local and the potental bugs from this.

package parkinglot

import (
	"time"
	"fmt"
	"errors"
	"space"
	"pass"
	"level"
	"spacetypes"
)

const (
	MonthlyPassCost      int = 42
	DefaultLevelCapacity int = 11
	SpecialWithoutPass int = 1
	RegularWithoutPass int = 2
	SpecialWithPass int = 3
	RegularyWithoutPass int = 4
)


// Design a parking lot with:
//	1. multi-floors,
//	2. gates,
//	3. pass holders,
//	4. one-time users,
//	5. special designated spots,
//	6. payment before leaving,
//	7. credit card and cash


type spaceType struct {
	isSpecial bool
	requiresPass bool
}

type gate struct {
	passes []*pass.Ticket
}

//type level struct {
//	numOfSpaces int // HACK: currently needed for faster check for hasAvailableSpaces. Find better way
//	unavailableSpaces []*space.Space
//	availableSpaces map[int]*space.SpacesQueue
//}
//
//func (currLevel *level) hasAvailableSpaces() bool {
//	return currLevel.numOfSpaces != len(currLevel.unavailableSpaces)
//}
//
//func (currLevel *level) hasAvailableSpacesOfType(spaceType int) bool {
//	return currLevel.numOfSpaces != len(currLevel.unavailableSpaces)
//}

type parkinglot struct {
	levels []*level.Level
	*gate
	pricePerHour int // TODO: move to gates
}

//// TODO: delete this when you know what you're doing
//func (lot *parkinglot) PrintLevels() {
//	fmt.Println(lot.levels)
//}

func (lot *parkinglot) AddNewLevel(numOfRegularSpaces, numOfSpecialSpaces int) error {
	// BUG: need to figure out how capacity works
	if len(lot.levels) == cap(lot.levels) {
		return errors.New("bug. Need to figure out how capacity works")
	}
	lot.levels[len(lot.levels)] = &level{
		occupiedSpaces: make(map[int]*space),
	}
	return nil
}

func NewLot(rate int, numOfSpacesByType map[spacetypes.SpaceType]int) (*parkinglot, error) {
	if rate < 1 {
		return nil, errors.New("hourly rate cannot be less than 1")
	}
	//newLot := &parkinglot{pricePerHour: hourlyRate, levels: make(map[string]*level)}
	newLot := &parkinglot{pricePerHour: rate, levels: make([]*level.Level, 0, DefaultLevelCapacity)}
	levelError := newLot.CreateNewLevel(numOfSpacesByType)
	//levelError := newLot.CreateNewLevel(numOfRegularSpaces, numOfSpecialSpaces, levelName)
	return newLot, levelError
}

//
func (lot parkinglot) BuyMonthlyPass(money int, needSpecial bool) (*pass, int) {
	// error if money < MonthlyPassCost
	return &pass{
		special: needSpecial,
		expires: time.Now().AddDate(0,1,0),
	}, MonthlyPassCost - money
}
