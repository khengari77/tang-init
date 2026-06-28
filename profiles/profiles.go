package profiles

type BoardProfile struct {
	ID           string
	Name         string
	FamilyString string
	DeviceString string
	ClockPin     int
	LEDPins      []int
	ButtonPins   []int
}

func All() []BoardProfile {
	return []BoardProfile{
		{
			ID:           "tangnano9k",
			Name:         "Tang Nano 9K",
			FamilyString: "GW1N-9C",
			DeviceString: "GW1NR-LV9QN88PC6/I5",
			ClockPin:     52,
			LEDPins:      []int{10, 11, 13, 14, 16, 38},
			ButtonPins:   []int{3, 4},
		},
		{
			ID:           "tangnano20k",
			Name:         "Tang Nano 20K",
			FamilyString: "GW2A-18C",
			DeviceString: "GW2AR-LV18QN88C8/I7",
			ClockPin:     4,
			LEDPins:      []int{15, 16, 17, 18, 19, 20},
			ButtonPins:   []int{88, 87},
		},
	}
}
