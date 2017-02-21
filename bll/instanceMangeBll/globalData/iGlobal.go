package globalData

type IGlobal interface {
	ModuleName() string
	Init() []error
}
