package modelData

type IModel interface {
	ModuleName() string
	Init() []error
	Check() []error
	Convert() []error
}
