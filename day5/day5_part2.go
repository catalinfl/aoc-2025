package day5

import (
	"fmt"
	"sort"
)

func (s *Solution) Day5part2() {
	interval, _ := getInput()
	ids := getIdsRange(interval)

	sort.Slice(ids, func(i, j int) bool {
		return ids[i].idStart < ids[j].idStart
	})

	fmt.Println(ids)

	merged := make([]id, 0, len(ids))

	for _, cur := range ids {
		if len(merged) == 0 || cur.idStart > merged[len(merged)-1].idEnd {
			merged = append(merged, cur)
			continue
		}

		if cur.idEnd > merged[len(merged)-1].idEnd {
			merged[len(merged)-1].idEnd = cur.idEnd
		}
	}

	counter := 0
	for _, m := range merged {
		counter += m.idEnd - m.idStart + 1
	}

	fmt.Println(counter)
}
