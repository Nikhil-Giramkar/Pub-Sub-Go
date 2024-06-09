package implementations

type video struct {
	title       string
	description string
}

func (v video) GetTitle() string {
	return v.title
}
func (v video) GetDescription() string {
	return v.description
}
