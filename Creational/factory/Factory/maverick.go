package Factory

type maverick struct {
	gun
}

func NewMaverick()IGun{
	return &maverick{
		gun{
			name : "maverick",
			power: 5,
		},
	}
}