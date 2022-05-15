package app

type ConfigureFunc = func() error
type RunFunc = func() error
type PostRunFunc = func(ac *AppContext) error
type AppContext struct {
	PreRuns  []RunFunc
	Runs     []RunFunc
	PostRuns []PostRunFunc

	Name        string
	Short       string
	Description string
	Version     string
}

func NewAppContext() *AppContext {
	return &AppContext{
		PreRuns:  make([]RunFunc, 0),
		Runs:     make([]RunFunc, 0),
		PostRuns: make([]PostRunFunc, 0),
	}
}

func (a *AppContext) PreRun(action RunFunc) {
	a.PreRuns = append(a.PreRuns, action)
}

func (a *AppContext) Run(action RunFunc) {
	a.Runs = append(a.Runs, action)
}

func (a *AppContext) PostRun(action PostRunFunc) {
	a.PostRuns = append(a.PostRuns, action)
}
