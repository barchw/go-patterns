/*
Package template -- Template Method

What is it?
Template Method is a behavioral design pattern that defines the skeleton of an algorithm in a base
method, delegating individual steps to interface implementations. The shared algorithm logic is
defined once, and the differences between variants are isolated in separate methods.

When to use?
  - When you have several algorithm variants that differ only in individual steps (e.g., export to CSV, JSON, XML).
  - When you want to enforce a specific order of algorithm steps while allowing different implementations.
  - When you want to avoid duplicating shared logic (e.g., iterating over data, building the result string).
  - When new variants should be added by implementing an interface, not by modifying existing code.

When NOT to use?
  - When the algorithm has no fixed structure -- if each variant requires a different order of steps.
  - When there is only one algorithm variant and you don't anticipate more -- the extra abstraction is unnecessary.
  - When algorithm steps are independent of each other -- it's better to use Strategy, which replaces the entire algorithm.

Tips and pitfalls:
  - Self-reference in the constructor: the exporter passes itself as the formatter (e.BaseExporter = NewBaseExporter(e)).
  - Column sorting ensures deterministic output, because map iteration in Go does not guarantee order.
  - An empty footer in CSV is valid -- not every format requires a footer.
  - Template Method vs Strategy: Strategy replaces the entire algorithm; Template Method replaces only individual steps.
*/
package template

import (
	"fmt"
	"sort"
	"strings"
)

// Formatter defines the interface for template steps: header, row, and footer of a data export.
type Formatter interface {
	Header(columns []string) string
	Row(columns []string, values map[string]string) string
	Footer(rowCount int) string
}

// Exporter defines the data exporter interface.
type Exporter interface {
	Export(data []map[string]string) string
}

// BaseExporter contains the template method Export, which defines the skeleton of the export algorithm.
type BaseExporter struct {
	formatter Formatter
}

// NewBaseExporter creates a base exporter with the given step formatter.
func NewBaseExporter(formatter Formatter) BaseExporter {
	return BaseExporter{formatter: formatter}
}

// Export implements the template method: extracts columns, generates the header, rows, and footer.
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

// CSVExporter exports data in CSV format. Embeds BaseExporter and implements Formatter.
type CSVExporter struct {
	BaseExporter
}

// NewCSVExporter creates a new CSV exporter, passing itself as the formatter.
func NewCSVExporter() *CSVExporter {
	e := &CSVExporter{}
	e.BaseExporter = NewBaseExporter(e)
	return e
}

// Header generates the CSV header -- column names separated by commas.
func (e *CSVExporter) Header(columns []string) string {
	return strings.Join(columns, ",") + "\n"
}

// Row generates a CSV row -- values separated by commas.
func (e *CSVExporter) Row(columns []string, values map[string]string) string {
	vals := make([]string, len(columns))
	for i, col := range columns {
		vals[i] = values[col]
	}
	return strings.Join(vals, ",") + "\n"
}

// Footer returns an empty footer -- CSV format does not require a footer.
func (e *CSVExporter) Footer(rowCount int) string {
	return ""
}

// JSONExporter exports data in JSON format. Embeds BaseExporter and implements Formatter.
type JSONExporter struct {
	BaseExporter
}

// NewJSONExporter creates a new JSON exporter, passing itself as the formatter.
func NewJSONExporter() *JSONExporter {
	e := &JSONExporter{}
	e.BaseExporter = NewBaseExporter(e)
	return e
}

// Header generates the beginning of a JSON array.
func (e *JSONExporter) Header(columns []string) string {
	return "[\n"
}

// Row generates a JSON object with key-value pairs for each column.
func (e *JSONExporter) Row(columns []string, values map[string]string) string {
	pairs := make([]string, len(columns))
	for i, col := range columns {
		pairs[i] = fmt.Sprintf("    %q: %q", col, values[col])
	}
	return "  {\n" + strings.Join(pairs, ",\n") + "\n  },\n"
}

// Footer generates the end of a JSON array.
func (e *JSONExporter) Footer(rowCount int) string {
	return "]\n"
}

// XMLExporter exports data in XML format with configurable element names.
type XMLExporter struct {
	BaseExporter
	rootElement string
	rowElement  string
}

// NewXMLExporter creates a new XML exporter with the given root and row element names.
func NewXMLExporter(rootElement, rowElement string) *XMLExporter {
	e := &XMLExporter{rootElement: rootElement, rowElement: rowElement}
	e.BaseExporter = NewBaseExporter(e)
	return e
}

// Header generates the opening tag of the XML root element.
func (e *XMLExporter) Header(columns []string) string {
	return fmt.Sprintf("<%s>\n", e.rootElement)
}

// Row generates an XML row element with fields for each column.
func (e *XMLExporter) Row(columns []string, values map[string]string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("  <%s>\n", e.rowElement))
	for _, col := range columns {
		sb.WriteString(fmt.Sprintf("    <%s>%s</%s>\n", col, values[col], col))
	}
	sb.WriteString(fmt.Sprintf("  </%s>\n", e.rowElement))
	return sb.String()
}

// Footer generates the closing tag of the XML root element.
func (e *XMLExporter) Footer(rowCount int) string {
	return fmt.Sprintf("</%s>\n", e.rootElement)
}
