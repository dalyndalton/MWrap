package wrapper

import (
	"regexp"
)

var logRegex = regexp.MustCompile(`(\[[0-9:]*\]) \[([A-z(-| )#0-9]*)\/([A-z #]*)\]: (.*)`)

type LogLine struct {
	timestamp  string
	threadName string
	level      string
	output     string
}

func ParseToLogLine(line string) *LogLine {
	matches := logRegex.FindAllStringSubmatch(line, 4)

	if matches == nil {
		return &LogLine{
			timestamp:  "",
			threadName: "",
			level:      "",
			output:     line,
		}
	}

	return &LogLine{
		timestamp:  matches[0][1],
		threadName: matches[0][2],
		level:      matches[0][3],
		output:     matches[0][4],
	}
}

type Event string

const (
	Ready    Event = "running"
	Booting  Event = "starting..."
	Stopping Event = "stopped"
)

var eventToRegexp = map[Event]*regexp.Regexp{
	Ready:    regexp.MustCompile(`Done (?s)(.*)! For help, type "help"`),
	Booting:  regexp.MustCompile(`Starting net.minecraft.server.Main`),
	Stopping: regexp.MustCompile(`Stopping (.*) server`),
}

func LogParser(line string) (*LogLine, *Event) {
	ll := ParseToLogLine(line)
	for e, r := range eventToRegexp {
		if r.MatchString(ll.output) {
			return ll, &e
		}
	}
	return ll, nil
}
