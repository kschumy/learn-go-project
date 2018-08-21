// NOTE: See note of issues involving time in parkinglot.go

package pass

import (
	"time"
	"space"
)

const rateForExpiredPass int = 2

type Ticket interface {
	getAmountOwed()
}

type oneTimeTicket struct {
	assignedSpace *space.Space
	pricePerHour int
	startTime time.Time
}

type passTicket struct {
	assignedSpace *space.Space
	startDate time.Time
	expireDate time.Time
	cost int
}

func (ticket *passTicket) IsExpired() bool {
	return ticket.expireDate.After(time.Now())
}

// BUG: this almost certainly does not work as expected. Issues is likely with time
// TODO: consider use of time.Round: https://golang.org/pkg/time/#Time.Round
func (ticket *oneTimeTicket) getAmountOwed() int {
	return int(time.Since(ticket.startTime).Hours()) * ticket.pricePerHour
}

// BUG: this almost certainly does not work as expected. Issues is likely with time
func (ticket *passTicket) getAmountOwed() int {
	if !ticket.IsExpired() {
		return 0
	}
	perHourCost := ticket.cost / int(ticket.expireDate.Sub(ticket.startDate)) * rateForExpiredPass
	return perHourCost * int(time.Since(ticket.startDate).Hours())
}

