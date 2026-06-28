package profiles

type BoardProfile struct {
	ID           string
	Name         string
	FamilyString string
	DeviceString string
	ClockPin     int
	ClockFreq    int
	LEDPins      []int
	ButtonPins   []int
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
		},
	}
}
