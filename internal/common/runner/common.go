package runner

import "github.com/kuritka/doge-action/internal/common/guards"

// CommonRunner is running all commands
type CommonRunner struct {
	command Runner
}

// NewCommonRunner creates new instance of CommonRunner
func NewCommonRunner(command Runner) *CommonRunner {
	return &CommonRunner{
		command,
	}
}

// MustRun runs command once and panics if command is broken
func (r *CommonRunner) MustRun() {
	err := r.command.Run()
	guards.Must(err, "command %s failed", r.command.String())
}
