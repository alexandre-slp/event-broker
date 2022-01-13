//Contracts are interfaces that are necessary to make decoupled code.
//Using interfaces to communicate between services facilitate all the testing process.

package infra

//Database All database interactors must follow this contract
type Database interface {
	addEvent()
	updateEvent()
	removeEvent()
	healthCheck()
}
