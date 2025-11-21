package gui

import (
	"fmt"
	"strings"
)

func NewPrinter() *Printer {
	return &Printer{
		sb:         &strings.Builder{},
		indent:     "",
		indentChar: "    ",
	}
}

type Printer struct {
	sb         *strings.Builder
	indent     string
	indentChar string
}

func (pr *Printer) Push() *Printer {
	pr.indent += pr.indentChar
	return pr
}

func (pr *Printer) Pop() *Printer {
	if len(pr.indent) > 0 {
		pr.indent = pr.indent[:len(pr.indent)-len(pr.indentChar)]
	}
	return pr
}

func (pr *Printer) PutNL() *Printer {
	pr.sb.WriteString("\n")
	return pr
}

func (pr *Printer) Put(t string, args ...interface{}) *Printer {
	if len(args) > 0 {
		t = fmt.Sprintf(t, args...)
	}

	if t != "" {
		if strings.Contains(t, "\n") {
			for _, line := range strings.Split(t, "\n") {
				if line == "" {
					pr.sb.WriteString("\n")
				} else {
					pr.sb.WriteString(fmt.Sprintf("%s%s\n", pr.indent, line))
				}
			}
		} else {
			pr.sb.WriteString(fmt.Sprintf("%s%s\n", pr.indent, t))
		}
	} else {
		pr.sb.WriteString("\n")
	}

	return pr
}

func (pr *Printer) Code() string {
	return strings.TrimRight(pr.sb.String(), "\n")
}
