package danymicData

type IDanymic interface {
	ModuleName() string
	Init() []error
	Check() []error
	Convert() []error
	Confirm()
}
