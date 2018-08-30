package space

import (
	"errors"
	"pass"
)

type Space struct {
	ticket *pass.Ticket
	number int
	next, prev *Space
	list *SpacesQueue
}

type SpacesQueue struct {
	front *Space
	back *Space
	size int
}

// Returns true if the space is occupied. Otherwise returns false.
func (space *Space) isOccupied() bool {
	return space.ticket != nil
}

// Returns true if the provided currSpace has a previous space. Otherwise returns false.
func (space *Space) hasPrev() bool {
	return space.prev != nil
}

// Returns true if the provided currSpace has a next space. Otherwise returns false.
func (space *Space) hasNext() bool {
	return space.next != nil
}

// If the providedspaceToRemove is in the list, removes the provided spaceToRemove and returns the
// spaceToRemove and nil. Otherwise, returns nil and an error
func (spaces *SpacesQueue) RemoveSpace(spaceToRemove *Space) (*Space, error){
	if spaceToRemove.list != spaces {
		return nil, errors.New("space must be in list")
	}
	if !spaceToRemove.hasNext() {
		return spaces.Dequeue() // spaceToRemove is last on list, so dequeue it
	}
	spaceToRemove.next.prev = spaceToRemove.prev
	if spaceToRemove.hasPrev() {
		spaceToRemove.prev.next = spaceToRemove.next
	}
	spaces.size--
	return spaceToRemove, nil
}

// Add the provided newSpace to the list
func (spaces *SpacesQueue) Enqueue(newSpace *Space) {
	newSpace.next = spaces.back
	if spaces.IsEmpty() {
		spaces.front = newSpace
	} else {
		spaces.back.prev = newSpace
	}
	spaces.back = newSpace
	newSpace.list = spaces
	spaces.size++
}

// If the list is not empty, removes and returns the space at the front of the list. Otherwise,
// returns nil and an error
func (spaces *SpacesQueue) Dequeue() (*Space, error) {
	if spaces.IsEmpty() {
		return nil, errors.New("no spaces to dequeue")
	}
	spaceToDequeue := spaces.front
	spaces.front = spaces.front.prev
	if spaces.front != nil {
		spaces.front.next, spaceToDequeue.prev = nil, nil // TODO: better way to do this in Go?
	}
	spaces.size--
	return spaceToDequeue, nil
}

// Returns true if the list is empty. Otherwise returns false.
func (spaces *SpacesQueue) IsEmpty() bool {
	return spaces.size == 0
}

func NewSpacesList() *SpacesQueue {
	return &SpacesQueue{}
}

// TODO: very unsure this will work
func EnqueueNSpaces(spacesStartNum, numOfSpaces int) (*SpacesQueue, error) {
	// TODO: the logic for numOfSpaces is repeated in levels. Fix??
	if spacesStartNum < 0 || numOfSpaces < 0 {
		return nil, errors.New("must have positive space numbers and number of spaces")
	}
	queueOfSpaces := &SpacesQueue{}
	for i := spacesStartNum; i <= numOfSpaces; i++ {
		queueOfSpaces.Enqueue(&Space{number: i}) // TODO: will this work?
	}
	return queueOfSpaces, nil
}


