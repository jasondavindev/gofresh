package main

import (
	"log"
	"strings"
)

type Option struct {
	value       interface{}
	description string
}

type Options map[string]*Option

func NewOptions() Options {
	return Options{
		"dir": {"", "Watch directory"},
		"h":   {false, "Show help"},
		"cmd": {"", "Command to run after reload"},
	}
}

var options = NewOptions()

func (options Options) IsBool(flag string) bool {
	_, ok := options[flag].value.(bool)
	return ok
}

func (options Options) Has(flag string) bool {
	_, ok := options[flag]
	return ok
}

func (options Options) Get(flag string) *Option {
	if options.Has(flag) {
		return options[flag]
	}

	return nil
}

func (options Options) Bool(flag string) bool {
	if !options.Has(flag) {
		return false
	}

	v, _ := options[flag].value.(bool)
	return v
}

func (options Options) Parse(args []string) []string {
	dirArgs := make([]string, 0)

	for _, arg := range args {
		if arg[0] == '-' {
			tokens := strings.SplitN(arg, "=", 2)
			flag := tokens[0][1:]

			if !options.Has(flag) {
				log.Fatalf("Invalid option: '%v'\n", flag)
				continue
			}

			switch len(tokens) {
			case 1:
				if options.IsBool(flag) {
					options[flag].value = true
				}

			case 2:
				options[flag].value = tokens[1]
			default:
				continue
			}
		} else {
			dirArgs = append(dirArgs, arg)
		}
	}

	return dirArgs
}
