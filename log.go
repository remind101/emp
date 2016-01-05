package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/mgutz/ansi"
)

var cmdLog = &Command{
	Run:      runLog,
	Usage:    "log",
	NeedsApp: true,
	Category: "app",
	Short:    "stream app log lines",
	Long: `
Log prints the streaming application log.
   Examples:
    $ emp log -a acme-inc
    2013-10-17T00:17:35.066089+00:00 app[web.1]: Completed 302 Found in 0ms
    ...
`,
}

func runLog(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.PrintUsage()
		os.Exit(2)
	}

	appName := mustApp()
	endpoint := fmt.Sprintf("/apps/%s/log-sessions", appName)

	r, w := io.Pipe()
	scan := bufio.NewScanner(r)
	colo := newColorizer(os.Stdout)
	go func() {
		for scan.Scan() {
			_, err := colo.Writeln(scan.Text())
			must(err)
		}
	}()

	must(client.Post(w, endpoint, nil))
}

type colorizer struct {
	colors      map[string]string
	colorScheme []string
	filter      *regexp.Regexp
	writer      io.Writer
}

func newColorizer(writer io.Writer) *colorizer {
	return &colorizer{
		colors: make(map[string]string),
		colorScheme: []string{
			"cyan",
			"yellow",
			"green",
			"magenta",
			"red",
		},
		// filter: regexp.MustCompile(`(?s)^(.*?\[([\w-]+)(?:[\d\.]+)?\]:)(.*)?$`),
		// 2015-09-17T12:52:02 [v3.web]: 2015/09/18 06:25:39 GET - /
		filter: regexp.MustCompile(`(.*)?(\[v3\.web\]:)(.*)?`),
		writer: writer,
	}
}

func (c *colorizer) resolve(p string) string {
	if color, ok := c.colors[p]; ok {
		return color
	}

	color := c.colorScheme[len(c.colors)%len(c.colorScheme)]
	c.colors[p] = color
	return color
}

func (c *colorizer) Writeln(p string) (n int, err error) {
	if c.filter.MatchString(p) {
		submatches := c.filter.FindStringSubmatch(p)
		return fmt.Fprintln(c.writer, ansi.Color(submatches[1]+submatches[2], "green")+ansi.ColorCode("reset")+submatches[3])
	}

	return fmt.Fprintln(c.writer, p)
}
