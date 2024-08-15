package core

type UnitStats struct {
	Speed   float64
	Attack  int
	Defence int
	Health  int
	// how to deal with archers and wizards?

}

type Stats struct {
	unitStats map[string]UnitStats
}

func NewStats() *Stats {
	return &Stats{
		unitStats: map[string]UnitStats{
			"blue-soldier": {
				Speed: 0.4,
			},
			"blue-archer": {
				Speed: 0.45,
			},
			"red-knight": {
				Speed: 0.7,
			},
			"wizard": {
				Speed: 0.2,
			},
		},
	}
}

func (r *Stats) GetUnitStats(key string) UnitStats {
	s, ok := r.unitStats[key]
	if !ok {
		panic("missing units stats:" + key)
	}
	return s
}
