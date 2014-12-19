package paizaio

type respObj interface {
	getError() error
}

type Runner struct {
	ID     string `json:"id"`
	Status string `json:"status"`

	Error string `json:"error"`
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

	Error string `json:"error"`
}

var _ respObj = &Details{}

// Languages is a list of languages supported by paiza.io.
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

// Status codes are value of Runner.Status or Details.Status.
const (
	StatusRunning   = "running"
	StatusCompleted = "completed"
)

// Results are value of Details.Result.
const (
	ResultSuccess = "success"
	ResultFailure = "failure"
	ResultError   = "error"
)

func (r *Runner) getError() error {
	if r.Error == "" {
		return nil
	}
	return newAPIError(r.Error)
}

func (d *Details) getError() error {
	if d.Error == "" {
		return nil
	}
	return newAPIError(d.Error)
}
