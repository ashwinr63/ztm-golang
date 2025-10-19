package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumebr int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumebr: 2}
		ella = Passenger{Name: "Ella", TicketNumebr: 3}
	)

	fmt.Println(bill, ella)

	var heidi Passenger
	heidi.Name = "Heidi"
	heidi.TicketNumebr = 4
	fmt.Println(heidi)

	casey.Boarded = true
	bill.Boarded = true
	ella.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill Boarded the bus")
	}
	if casey.Boarded {
		fmt.Println(casey.Name, "has boarded the bus")
	}

	heidi.Boarded = true
	bus := Bus{heidi}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}

