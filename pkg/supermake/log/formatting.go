package log

import (
	"fmt"
	"strings"
)

type TabulatedRow struct {
	fields    []string
	sizes     []int
	lengths   []int
	delimiter string
}

const elipsis = "..."

func TabulateRow() TabulatedRow {
	return TabulatedRow{
		fields:    make([]string, 0),
		sizes:     make([]int, 0),
		lengths:   make([]int, 0),
		delimiter: " ",
	}
}

func (t TabulatedRow) Delimiter(delimiter string) TabulatedRow {
	return TabulatedRow{
		t.fields,
		t.sizes,
		t.lengths,
		delimiter,
	}
}

func (t TabulatedRow) Field(text string, width int, formatting ...ShellFormatting) TabulatedRow {
	if width > 0 && len(text) > width {
		text = text[:width-len(elipsis)] + elipsis
	}

	length := len(text)

	if len(formatting) > 0 {
		text = FormatText(text, formatting...)
	}

	return TabulatedRow{
		append(t.fields, text),
		append(t.sizes, width),
		append(t.lengths, length),
		t.delimiter,
	}
}

func (t TabulatedRow) String() string {
	builder := strings.Builder{}

	delimiter := ""
	for i := 0; i < len(t.fields); i++ {
		field := t.fields[i]
		width := t.sizes[i]
		length := t.lengths[i]

		padding := ""
		if width > 0 {
			padding = strings.Repeat(" ", width-length)
		}
		builder.WriteString(fmt.Sprintf("%s%s%s", delimiter, field, padding))
		delimiter = t.delimiter
	}

	builder.WriteRune('\n')
	return builder.String()
}
