//go:build !solution

package hotelbusiness

import (
	"maps"
	"slices"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	if len(guests) == 0 {
		return nil
	}

	changesInGuests := make(map[int]int)
	for _, guest := range guests {
		changesInGuests[guest.CheckInDate]++
		changesInGuests[guest.CheckOutDate]--
	}

	sortedChangeDates := slices.Collect(maps.Keys(changesInGuests))
	slices.Sort(sortedChangeDates)

	currentGuests := 0
	res := make([]Load, 0, len(sortedChangeDates))
	for _, date := range sortedChangeDates {
		if changesInGuests[date] != 0 {
			currentGuests += changesInGuests[date]
			res = append(res, Load{date, currentGuests})
		}
	}
	return res
}
