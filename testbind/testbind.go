package testbind

// Providers provides apis to access the container cloud platform
type Providers struct {
	Bind *Bind
}

// Constructor a new Providers
func New(headers map[string]string) *Providers {
	providers := &Providers{
		Bind: &Bind{},
	}

	return providers
}
