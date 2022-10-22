//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//* Use the accessGranted() and accessDenied() functions to display
//  informational messages
//* Access at any time: Admin, Manager
//* Access weekends: Contractor
//* Access weekdays: Member
//* Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Granted")
}

func accessDenied() {
	fmt.Println("Denied")
}

func isWeekday(day int) bool {
	return day < 5
}

func isMondayWednesdayFriday(day int) bool {
	return day == 0 || day == 2 || day == 4
}

func isAllowed(role int, day int) bool {
	if role <= Manager {
		return true
	}

	weekday := isWeekday(day)

	if !weekday {
		return role == Contractor
	}

	switch role {
	case Guest:
		return isMondayWednesdayFriday(day)
	case Member:
		return true
	default:
		return false
	}
}

func getDayAndRole() (int, int) {
	var day int
	var role int

	fmt.Println(`What day is it?
		Monday    = 0
		Tuesday   = 1
		Wednesday = 2
		Thursday  = 3
		Friday    = 4
		Saturday  = 5
		Sunday    = 6
	`)
	fmt.Scanln(&day)

	fmt.Println(`Who are you?
		Admin      = 10
		Manager    = 20
		Contractor = 30
		Member     = 40
		Guest      = 50
	`)
	fmt.Scanln(&role)

	return day, role
}

func main() {
	// The day and role. Change these to check your work.
	// today, role := Tuesday, Guest
	day, role := getDayAndRole()

	allowed := isAllowed(role, day)

	if allowed {
		accessGranted()
	} else {
		accessDenied()
	}
}
