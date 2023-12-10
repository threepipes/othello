package prompt

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Prompt struct {
	io *bufio.ReadWriter
}

func NewPrompt(r io.Reader, w io.Writer) *Prompt {
	return &Prompt{
		bufio.NewReadWriter(bufio.NewReader(r), bufio.NewWriter(w)),
	}
}

func (p Prompt) Choose(options []string, msg string) (int, error) {
	p.io.WriteString(msg + "\n")
	for i, o := range options {
		p.io.WriteString(fmt.Sprintf("[%d] %s\n", i, o))
	}
	p.io.Flush()
	for true {
		pick, err := p.io.ReadString('\n')
		pick = strings.TrimSpace(pick)
		if err != nil {
			return 0, fmt.Errorf("failed to read input")
		}
		c, err := strconv.Atoi(pick)
		if err != nil || c < 0 || c >= len(options) {
			p.io.WriteString(fmt.Sprintf("Invalid input: '%s'. Please pick from 0 to %d.\n", pick, len(options)-1))
			p.io.Flush()
			continue
		}
		return c, nil
	}
	return 0, fmt.Errorf("unknown error has occured")
}

func (p Prompt) InputStringRegexMatch(regex, msg string) ([]string, error) {
	p.io.WriteString(msg + "\n")
	p.io.Flush()
	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse regex '%s': %w", regex, err)
	}
	for true {
		s, err := p.io.ReadString('\n')
		s = strings.TrimSpace(s)
		if err != nil {
			return nil, fmt.Errorf("failed to read input")
		}
		match := r.FindStringSubmatch(s)
		if len(match) == 0 {
			p.io.WriteString(fmt.Sprintf("Invalid input: '%s'\n", s))
			p.io.WriteString(msg)
			p.io.Flush()
			continue
		}
		return match, nil
	}
	return nil, fmt.Errorf("unknown error has occured")
}
