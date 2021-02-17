package Factory

type Ak47 struct {
	gun
}
func NewAk47() IGun{
	return &Ak47{
		gun: gun{
			name : "ak47",
			power : 3,
		},
	}
}
