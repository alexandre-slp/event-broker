//Contracts are interfaces that are necessary to make decoupled code.
//Using interfaces to communicate between services facilitate all the testing process.

package infra

type database interface {
	addEvent()
	updateEvent()
	removeEvent()
	healthCheck()
}
