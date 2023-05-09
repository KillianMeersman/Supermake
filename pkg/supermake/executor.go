package supermake

type Executor interface {
	Execute(executable Executable) error
}

type DockerEnvironment struct {
	Image      string
	Entrypoint string
}

func (d *DockerEnvironment) Execute(executable Executable) error {
	return nil
}

type LocalEnvironment struct{}

func (l *LocalEnvironment) Execute(executable Executable) error {
	return nil
}
