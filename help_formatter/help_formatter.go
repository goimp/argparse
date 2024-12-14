package help_formatter

import (
	"argparse"
	"fmt"
	"regexp"
	"strings"
)

// HelpFormatter handles the generation of help messages and formatting.
type HelpFormatter struct {
	Prog              string
	IndentIncrement   int
	MaxHelpPosition   int
	Width             int
	CurrentIndent     int
	Level             int
	ActionMaxLength   int
	RootSection       *Section
	CurrentSection    *Section
	WhitespaceMatcher *regexp.Regexp
	LongBreakMatcher  *regexp.Regexp
}

// NewHelpFormatter creates a new instance of HelpFormatter
func NewHelpFormatter(prog string, indentIncrement, maxHelpPosition, width int) *HelpFormatter {
	// Default width setting
	if width == 0 {
		width = getTerminalWidth() - 2
	}

	// Ensure maxHelpPosition is not too large
	maxHelpPosition = min(maxHelpPosition, max(width-20, indentIncrement*2))

	// Create a new HelpFormatter
	return &HelpFormatter{
		Prog:              prog,
		IndentIncrement:   indentIncrement,
		MaxHelpPosition:   maxHelpPosition,
		Width:             width,
		CurrentIndent:     0,
		Level:             0,
		ActionMaxLength:   0,
		RootSection:       &Section{Formatter: nil}, // Root section
		CurrentSection:    &Section{Formatter: nil}, // Current section
		WhitespaceMatcher: regexp.MustCompile(`\s+`),
		LongBreakMatcher:  regexp.MustCompile(`\n\n\n+`),
	}
}

// getTerminalWidth attempts to get the terminal's width, similar to shutil.get_terminal_size()
func getTerminalWidth() int {
	return 80
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the larger of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FormatHelp formats the help message for the command.
func (hf *HelpFormatter) FormatHelp() string {
	// Implement formatting logic here (simplified for demonstration)
	return fmt.Sprintf("Usage: %s [options]", hf.Prog)
}

// Repr returns a string representation of the HelpFormatter
func (hf *HelpFormatter) Repr() string {
	return fmt.Sprintf("%s(indentIncrement=%d, maxHelpPosition=%d, width=%d)", hf.Prog, hf.IndentIncrement, hf.MaxHelpPosition, hf.Width)
}

// _indent increments the current indent and level
func (hf *HelpFormatter) indent() {
	hf.CurrentIndent += hf.IndentIncrement
	hf.Level++
}

// _dedent decrements the current indent and level
func (hf *HelpFormatter) dedent() {
	hf.CurrentIndent -= hf.IndentIncrement
	if hf.CurrentIndent < 0 {
		panic("Indent decreased below 0.")
	}
	hf.Level--
}

// ========================
// Message building methods
// ========================

// startSection starts a new section with a given heading.
func (hf *HelpFormatter) startSection(heading string) {
	hf.indent()
	section := &Section{
		Formatter: hf,
		Parent:    hf.CurrentSection,
		Heading:   heading,
		Items:     []SectionItem{},
	}
	hf.CurrentSection.Items = append(hf.CurrentSection.Items, SectionItem{
		FormatFunc: section.formatHelp,
		Args:       nil,
	})
	hf.CurrentSection = section
}

// endSection ends the current section and dedents.
func (hf *HelpFormatter) endSection() {
	hf.CurrentSection = hf.CurrentSection.Parent
	hf.dedent()
}

// addText adds a plain text item to the current section.
func (hf *HelpFormatter) addText(text string) {
	if text != argparse.SUPPRESS && text != "" {
		hf.CurrentSection.Items = append(hf.CurrentSection.Items, SectionItem{
			FormatFunc: hf.formatText,
			Args:       []interface{}{text},
		})
	}
}

// addUsage adds a usage item to the current section.
func (hf *HelpFormatter) addUsage(usage string, actions []string, groups []string, prefix string) {
	if usage != argparse.SUPPRESS {
		// Example: Use the data to format the usage string in a more complex manner
		usageText := fmt.Sprintf("Usage: %s %s", usage, strings.Join(actions, " "))
		hf.CurrentSection.Items = append(hf.CurrentSection.Items, SectionItem{
			FormatFunc: hf.formatUsage,
			Args:       []interface{}{usageText},
		})
	}
}

// formatText formats the text item.
func (hf *HelpFormatter) formatText(args ...interface{}) string {
	return fmt.Sprintf("%s", args[0])
}

// formatUsage formats the usage item.
func (hf *HelpFormatter) formatUsage(args ...interface{}) string {
	return fmt.Sprintf("%s", args[0])
}
