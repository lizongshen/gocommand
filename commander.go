package gocommand

type Commander interface {
	Exec(args ...string) (int, string, error)
	ExecAsync(stdout chan string, args ...string) int
	ExecNoWait(args ...string) error
}
