package day_05

import (
	"sort"
)

type Map []Edge

func (m Map) Next(i int) int {
	m.SortBySrcAsc()
	for _, edge := range m {
		if i >= edge[1] && i <= edge.MaxSrcValue() {
			return edge[0] + i - edge[1]
		}
	}
	return i
}

func (m Map) SortByDestAsc() {
	sort.Slice(m, func(i, j int) bool {
		return m[i][0] < m[j][0]
	})
}

func (m Map) SortBySrcAsc() {
	sort.Slice(m, func(i, j int) bool {
		return m[i][1] < m[j][1]
	})
}

func FromSensibleUnits(edges [][3]int) Map {
	// shift, srcStart, srcEnd
	m := Map{}
	for _, edge := range edges {
		m = append(m, Edge{edge[1] + edge[0], edge[1], edge[2] - edge[1] + 1})
	}
	return m
}

func (m Map) MaxValueInSameSrcRange(givenValue, maxSrc int) int {
	m.SortBySrcAsc()
	for _, edge := range m {
		if edge[1] > givenValue {
			return edge[1] - 1
		}
		if edge[1] <= givenValue && edge.MaxSrcValue() >= givenValue {
			return edge.MaxSrcValue()
		}
	}
	return maxSrc
}

func (m Map) EdgeContaining(givenValue int) Edge {
	m.SortBySrcAsc()
	for _, edge := range m {
		if edge[1] > givenValue {
			return Edge{givenValue, givenValue, edge[1] - givenValue}
		}
		if edge[1] <= givenValue && edge.MaxSrcValue() >= givenValue {
			return edge
		}
	}
	return Edge{-1, -1, -1}
}

func (m1 Map) Merge(m2 Map) Map {
	m1.SortBySrcAsc()
	m2.SortBySrcAsc()
	// We're going to work upwards (by src value) through the map, defining new edges.
	merged := Map{}
	// We need to start with the lowest possible value that will be impacted by
	// either map, then work upwards until we've defined an edge for the max value that
	// will be impacted by either map.
	i := min(m1[0][1], m2[0][1])
	stop := max(m1[len(m1)-1].MaxSrcValue(), m2[len(m2)-1].MaxSrcValue())
	for i <= stop {
		// fmt.Printf("i is now: %v\n", i)
		m1Edge := m1.EdgeContaining(i)
		m1Shift := m1Edge[0] - m1Edge[1]
		// Map the value through m1 to its destination "a".
		// Find the max of the src edge in m2 that A falls into.
		// Map that value back into m1's units by applying m1's shift.
		a := m1.Next(i)
		maxM2Edge := m2.MaxValueInSameSrcRange(a, stop)
		maxM1Edge := min(m1Edge.MaxSrcValue(), maxM2Edge-m1Shift)

		// We now have a src edge that'll all map to the same dest edge.
		merged = append(merged, Edge{m2.Next(a), i, maxM1Edge - i + 1})
		i = maxM1Edge + 1
	}
	return merged
}

func MergeAll(maps []Map) Map {
	merged := maps[0]
	for _, map_ := range maps[1:] {
		merged = merged.Merge(map_)
	}
	return merged
}
