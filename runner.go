package paizaio

type respObj interface {
	getError() error
}

type Runner struct {
	ID     string `json:"id"`
	Status string `json:"status"`

	err string `json:"error"`
}

var _ respObj = &Runner{}

type Details struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Note     string `json:"note"`
	Status   string `json:"status"`

	BuildStdout   string `json:"build_stdout"`
	BuildStderr   string `json:"build_stderr"`
	BuildExitCode int    `json:"build_exit_code"`
	BuildTime     string `json:"build_time"`
	BuildMemory   int    `json:"build_memory"`
	BuildResult   string `json:"build_result"`

	Stdout   string `json:"stdout"`
	Stderr   string `json:"stderr"`
	ExitCode int    `json:"exit"`
	Time     string `json:"time"`
	Memory   int    `json:"memory"`
	Result   string `json:"result"`

	err string `json:"error"`
}

var _ respObj = &Details{}

var Languages = []string{
	"c",
	"cpp",
	"objective-c",
	"java",
	"scala",
	"csharp",
	"go",
	"haskell",
	"erlang",
	"perl",
	"python",
	"python3",
	"ruby",
	"php",
	"bash",
	"r",
	"javascript",
	"coffeescript",
	"vb",
	"cobol",
	"fsharp",
	"d",
	"clojure",
	"mysql",
}

const (
	StatusRunning   = "running"
	StatusCompleted = "completed"

	ResultSuccess = "success"
	ResultFailure = "failure"
	ResultError   = "error"
)

func (r *Runner) getError() error {
	if r.err == "" {
		return nil
	}
	return newAPIError(r.err)
}

func (d *Details) getError() error {
	if d.err == "" {
		return nil
	}
	return newAPIError(d.err)
}
