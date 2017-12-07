// Package goruntime is a runtime for the Go command
package gorun

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/micro/go-run"
)

type goRuntime struct {
	// GOPATH
	Path string
	// Go Command
	Cmd string
}

func (g *goRuntime) Fetch(url string, opts ...run.FetchOption) (*run.Source, error) {
	var options run.FetchOptions
	for _, o := range opts {
		o(&options)
	}

	args := []string{"get", "-d"}

	if options.Update {
		args = append(args, "-u")
	}

	cmd := exec.Command(g.Cmd, append(args, url)...)
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return &run.Source{
		URL: url,
		Dir: filepath.Join(g.Path, "src", url),
	}, nil
}

func (g *goRuntime) Build(src *run.Source) (*run.Binary, error) {
	cmd := exec.Command(g.Cmd, "install", src.URL)
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return &run.Binary{
		Path:   filepath.Join(g.Path, "bin", filepath.Base(src.URL)),
		Source: src,
	}, nil
}

func (g *goRuntime) Exec(bin *run.Binary) (*run.Process, error) {
	cmd := exec.Command(bin.Path)
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	return &run.Process{
		ID:     fmt.Sprintf("%d", cmd.Process.Pid),
		Binary: bin,
	}, nil
}

func (g *goRuntime) Kill(proc *run.Process) error {
	pid, err := strconv.Atoi(proc.ID)
	if err != nil {
		return err
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	return p.Kill()
}

func (g *goRuntime) Wait(proc *run.Process) error {
	pid, err := strconv.Atoi(proc.ID)
	if err != nil {
		return err
	}

	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}

	ps, err := p.Wait()
	if err != nil {
		return err
	}

	if ps.Success() {
		return nil
	}

	return errors.New(ps.String())
}

// whichGo locates the go command
func whichGo() string {
	// check GOROOT
	if gr := os.Getenv("GOROOT"); len(gr) > 0 {
		return filepath.Join(gr, "bin", "go")
	}

	// check path
	for _, p := range filepath.SplitList(os.Getenv("PATH")) {
		bin := filepath.Join(p, "go")
		if _, err := os.Stat(bin); err == nil {
			return bin
		}
	}

	// best effort
	return "go"
}

func NewRuntime() run.Runtime {
	cmd := whichGo()
	path := os.Getenv("GOPATH")

	// point of no return
	if len(cmd) == 0 {
		panic("Could not find Go executable")
	}

	// set path if not exists
	if len(path) == 0 {
		path = os.TempDir()
		os.Setenv("GOPATH", path)
	}

	return &goRuntime{
		Cmd:  cmd,
		Path: path,
	}
}
