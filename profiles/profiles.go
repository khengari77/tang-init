package profiles

type HDMIPins struct {
	ClkP int
	D0P  int
	D1P  int
	D2P  int
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
			HDMIPins:     HDMIPins{ClkP: 28, D0P: 30, D1P: 32, D2P: 35},
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
			HDMIPins:     HDMIPins{ClkP: 69, D0P: 71, D1P: 73, D2P: 75},
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
			HDMIPins:     HDMIPins{ClkP: 33, D0P: 35, D1P: 37, D2P: 39},
		},
	}
}
