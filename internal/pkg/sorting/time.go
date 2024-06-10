package sorting

import (
	"sort"
	"time"
)

type TimeSortable interface {
	GetTime() time.Time
}

func SortByTime[T TimeSortable](items []T) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].GetTime().Before(items[j].GetTime())
	})
}
