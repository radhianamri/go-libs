package scanner

import (
	"bufio"
	"os"
)

type Scanner struct {
	*bufio.Scanner
	File *os.File
}

func (s *Scanner) Close() error {
	return s.File.Close()
}

func New(filename string, maxBufferSize int) (*Scanner, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	if maxBufferSize > 0 {
		scanner.Buffer(make([]byte, 2), maxBufferSize)
	}

	return &Scanner{
		File:    f,
		Scanner: scanner,
	}, nil
}

type LineProcessor struct {
	Scanner
	err   error
	funcs []func(*string) error
	text  string
}

func (s *LineProcessor) Scan() bool {
	if s.Scanner.Scan() && s.err == nil {
		s.text = s.Scanner.Text()
		for i := range s.funcs {
			s.err = s.funcs[i](&s.text)
			if s.err != nil {
				return false
			}
		}
		return true
	}
	return false
}

func (s *LineProcessor) Text() string {
	return s.text
}

func (s *LineProcessor) Err() error {
	if s.err != nil {
		return s.err
	}
	return s.Scanner.Err()
}

func NewLineProcesser(filename string, maxBufferSize int, lineProcessors ...func(*string) error) (*LineProcessor, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	if maxBufferSize > 0 {
		scanner.Buffer(make([]byte, maxBufferSize), maxBufferSize)
	}

	return &LineProcessor{
		Scanner: Scanner{
			File:    f,
			Scanner: scanner,
		},
		funcs: lineProcessors,
	}, nil
}
