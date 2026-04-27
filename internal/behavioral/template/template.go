package template

import (
	"fmt"
	"sort"
	"strings"
)

type Formatter interface {
	Header(columns []string) string
	Row(columns []string, values map[string]string) string
	Footer(rowCount int) string
}

type Exporter interface {
	Export(data []map[string]string) string
}

type BaseExporter struct {
	formatter Formatter
}

func NewBaseExporter(formatter Formatter) BaseExporter {
	return BaseExporter{formatter: formatter}
}

func (b *BaseExporter) Export(data []map[string]string) string {
	if len(data) == 0 {
		return ""
	}

	columns := extractColumns(data[0])
	sort.Strings(columns)

	var sb strings.Builder
	sb.WriteString(b.formatter.Header(columns))
	for _, row := range data {
		sb.WriteString(b.formatter.Row(columns, row))
	}
	sb.WriteString(b.formatter.Footer(len(data)))
	return sb.String()
}

func extractColumns(row map[string]string) []string {
	columns := make([]string, 0, len(row))
	for k := range row {
		columns = append(columns, k)
	}
	return columns
}

type CSVExporter struct {
	BaseExporter
}

func NewCSVExporter() *CSVExporter {
	e := &CSVExporter{}
	e.BaseExporter = NewBaseExporter(e)
	return e
}

func (e *CSVExporter) Header(columns []string) string {
	return strings.Join(columns, ",") + "\n"
}

func (e *CSVExporter) Row(columns []string, values map[string]string) string {
	vals := make([]string, len(columns))
	for i, col := range columns {
		vals[i] = values[col]
	}
	return strings.Join(vals, ",") + "\n"
}

func (e *CSVExporter) Footer(rowCount int) string {
	return ""
}

type JSONExporter struct {
	BaseExporter
}

func NewJSONExporter() *JSONExporter {
	e := &JSONExporter{}
	e.BaseExporter = NewBaseExporter(e)
	return e
}

func (e *JSONExporter) Header(columns []string) string {
	return "[\n"
}

func (e *JSONExporter) Row(columns []string, values map[string]string) string {
	pairs := make([]string, len(columns))
	for i, col := range columns {
		pairs[i] = fmt.Sprintf("    %q: %q", col, values[col])
	}
	return "  {\n" + strings.Join(pairs, ",\n") + "\n  },\n"
}

func (e *JSONExporter) Footer(rowCount int) string {
	return "]\n"
}

type XMLExporter struct {
	BaseExporter
	rootElement string
	rowElement  string
}

func NewXMLExporter(rootElement, rowElement string) *XMLExporter {
	e := &XMLExporter{rootElement: rootElement, rowElement: rowElement}
	e.BaseExporter = NewBaseExporter(e)
	return e
}

func (e *XMLExporter) Header(columns []string) string {
	return fmt.Sprintf("<%s>\n", e.rootElement)
}

func (e *XMLExporter) Row(columns []string, values map[string]string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("  <%s>\n", e.rowElement))
	for _, col := range columns {
		sb.WriteString(fmt.Sprintf("    <%s>%s</%s>\n", col, values[col], col))
	}
	sb.WriteString(fmt.Sprintf("  </%s>\n", e.rowElement))
	return sb.String()
}

func (e *XMLExporter) Footer(rowCount int) string {
	return fmt.Sprintf("</%s>\n", e.rootElement)
}
