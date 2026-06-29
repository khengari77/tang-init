package profiles

type HDMIPins struct {
	ClkP int
	ClkN int
	D0P  int
	D0N  int
	D1P  int
	D1N  int
	D2P  int
	D2N  int
}

func (p HDMIPins) all() []struct{ name string; pin int } {
	return []struct{ name string; pin int }{
		{"tmds_clk_p", p.ClkP},
		{"tmds_clk_n", p.ClkN},
		{"tmds_d_p[0]", p.D0P},
		{"tmds_d_n[0]", p.D0N},
		{"tmds_d_p[1]", p.D1P},
		{"tmds_d_n[1]", p.D1N},
		{"tmds_d_p[2]", p.D2P},
		{"tmds_d_n[2]", p.D2N},
	}
}

type BoardProfile struct {
	ID           string
	Name         string
	FamilyString string
	DeviceString string
	ClockPin     int
	ClockFreq    int
	LEDPins      []int
	ButtonPins   []int
	HasHDMI      bool
	HDMIPins     HDMIPins
}

func All() []BoardProfile {
	return []BoardProfile{
		{
			ID:           "tangnano",
			Name:         "Tang Nano (original)",
			FamilyString: "GW1N-1",
			DeviceString: "GW1N-LV1QN48C6/I5",
			ClockPin:     35,
			ClockFreq:    24_000_000,
			LEDPins:      []int{16, 17, 18},
			ButtonPins:   []int{15},
			HasHDMI:      false,
		},
		{
			ID:           "tangnano1k",
			Name:         "Tang Nano 1K",
			FamilyString: "GW1NZ-1",
			DeviceString: "GW1NZ-LV1QN48C6/I5",
			ClockPin:     47,
			ClockFreq:    27_000_000,
			LEDPins:      []int{9, 11, 10},
			ButtonPins:   []int{13, 44},
			HasHDMI:      false,
		},
		{
			ID:           "tangnano4k",
			Name:         "Tang Nano 4K",
			FamilyString: "GW1NSR-4C",
			DeviceString: "GW1NSR-LV4CQN48PC6/I5",
			ClockPin:     45,
			ClockFreq:    27_000_000,
			LEDPins:      []int{10},
			ButtonPins:   []int{14, 15},
			HasHDMI:      true,
			HDMIPins:     HDMIPins{ClkP: 28, ClkN: 29, D0P: 30, D0N: 31, D1P: 32, D1N: 33, D2P: 35, D2N: 36},
		},
		{
			ID:           "tangnano9k",
			Name:         "Tang Nano 9K",
			FamilyString: "GW1NR-9C",
			DeviceString: "GW1NR-LV9QN88PC6/I5",
			ClockPin:     52,
			ClockFreq:    27_000_000,
			LEDPins:      []int{10, 11, 13, 14, 16, 38},
			ButtonPins:   []int{3, 4},
			HasHDMI:      true,
			HDMIPins:     HDMIPins{ClkP: 69, ClkN: 70, D0P: 71, D0N: 72, D1P: 73, D1N: 74, D2P: 75, D2N: 76},
		},
		{
			ID:           "tangnano20k",
			Name:         "Tang Nano 20K",
			FamilyString: "GW2A-18C",
			DeviceString: "GW2AR-LV18QN88C8/I7",
			ClockPin:     4,
			ClockFreq:    27_000_000,
			LEDPins:      []int{15, 16, 17, 18, 19, 20},
			ButtonPins:   []int{88, 87},
			HasHDMI:      true,
			HDMIPins:     HDMIPins{ClkP: 33, ClkN: 34, D0P: 35, D0N: 36, D1P: 37, D1N: 38, D2P: 39, D2N: 40},
		},
	}
}
