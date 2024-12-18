package help_formatter

import (
	"argparse"
	"fmt"
	"regexp"
	"strings"
)

type HelpFormatterInterface interface {
	Struct() *HelpFormatter
	Indent_()
	Dedent_()
	AddItem_(func(...any) string, ...any)

	StartSection(string)
	EndSection()
	AddText(string)
	AddUsage(string, []argparse.ActionInterface, []argparse.ActionsContainerInterface, string)

	AddArgument(argparse.ActionInterface)
	AddArguments([]argparse.ActionInterface)

	FormatHelp(...any) string
	JoinParts_([]string) string
}

const DefaultTerminalWidth = 80

func GetTerminalWidth() int {
	return DefaultTerminalWidth
}

type SectionItem_ struct {
	Func func(...any) string
	Args []any
}

type Section_ struct {
	Formatter HelpFormatterInterface
	Parent    *Section_
	Heading   string
	Items     []*SectionItem_
}

func (s *Section_) FormatHelp(...any) string {
	// format the indented section
	if s.Parent != nil {
		s.Formatter.Indent_()
	}
	join := s.Formatter.JoinParts_
	var helpStrings []string
	for _, item := range s.Items {
		helpStrings = append(helpStrings, item.Func(item.Args...))
	}
	itemHelp := join(helpStrings)
	if s.Parent != nil {
		s.Formatter.Dedent_()
	}

	// return nothing if the section was empty
	if itemHelp == "" {
		return ""
	}

	var heading string
	// add the heading if the section was non-empty
	if s.Heading != argparse.SUPPRESS && s.Heading != "" {
		currentIndent := s.Formatter.Struct().CurrentIndent
		headingText := fmt.Sprintf("%s:", s.Heading)
		heading = fmt.Sprintf("%*s%s\n", currentIndent, "", headingText)
	} else {
		heading = ""
	}

	// join the section-initial newline, the heading and the help
	return join([]string{"\n", heading, itemHelp, "\n"})
}

// HelpFormatter handles the generation of help messages and formatting.
type HelpFormatter struct {
	Prog_             string
	IndentIncrement   int
	MaxHelpPosition   int
	Width             int
	CurrentIndent     int
	Level             int
	ActionMaxLength   int
	RootSection       *Section_
	CurrentSection    *Section_
	WhitespaceMatcher *regexp.Regexp
	LongBreakMatcher  *regexp.Regexp
}

func NewHelpFormatter(prog string, indentIncrement, maxHelpPosition, width int) HelpFormatterInterface {

	// default setting for width
	if width == 0 {
		width = GetTerminalWidth()
		width -= 2
	}

	formatter := &HelpFormatter{
		Prog_:           prog,
		IndentIncrement: indentIncrement,
		MaxHelpPosition: min(maxHelpPosition, max(width-20, indentIncrement*2)),
		Width:           width,
		CurrentIndent:   0,
		Level:           0,
		ActionMaxLength: 0,
		// RootSection: ,
		// CurrentSection: ,
		WhitespaceMatcher: regexp.MustCompile(`\s+`),
		LongBreakMatcher:  regexp.MustCompile(`\n\n\n+`),
	}

	return formatter
}

func (hf *HelpFormatter) Struct() *HelpFormatter {
	return hf
}

func (hf *HelpFormatter) Indent_() {
	hf.CurrentIndent += hf.IndentIncrement
	hf.Level++
}

func (hf *HelpFormatter) Dedent_() {
	hf.CurrentIndent -= hf.IndentIncrement
	if hf.CurrentIndent < 0 {
		panic("Indent decreased below 0.")
	}
	hf.Level--
}

func (hf *HelpFormatter) AddItem_(func_ func(...any) string, args ...any) {
	section := &SectionItem_{
		Func: func_,
		Args: args,
	}
	hf.CurrentSection.Items = append(hf.CurrentSection.Items, section)
}

// ========================
// Message building methods
// ========================

func (hf *HelpFormatter) StartSection(heading string) {
	hf.Indent_()
	section := &Section_{
		Formatter: hf,
		Parent:    hf.CurrentSection,
		Heading:   heading,
	}
	hf.AddItem_(section.FormatHelp, nil)
	hf.CurrentSection = section
}

func (hf *HelpFormatter) EndSection() {
	hf.CurrentSection = hf.CurrentSection.Parent
	hf.Dedent_()
}

func (hf *HelpFormatter) AddText(text string) {
	if text != argparse.SUPPRESS && text != "" {
		hf.AddItem_(hf.FormatHelp, text)
	}
}

func (hf *HelpFormatter) AddUsage(usage string, actions []argparse.ActionInterface, groups []argparse.ActionsContainerInterface, prefix string) {
	if usage != argparse.SUPPRESS {
		args := []any{usage, actions, groups, prefix}
		hf.AddItem_(hf.FormatHelp, args...)
	}
}

func (hf *HelpFormatter) AddArgument(action argparse.ActionInterface) {
	if action.Struct().Help != argparse.SUPPRESS {

		// find all invocations
		getInvocation := hf.FormatActionInvocation_
		invocationLengths := []int{len(getInvocation(action)) + hf.CurrentIndent}
		for _, subaction := range hf.IterIndentedSubactions_(action) {
			invocationLengths = append(invocationLengths, len(getInvocation(subaction)+hf.CurrentIndent))
		}

		// update the maximum item length
		actionLength := max(invocationLengths...)
		hf.ActionMaxLength = max(hf.ActionMaxLength, actionLength)
	}
}

func (hf *HelpFormatter) AddArguments(actions []argparse.ActionInterface) {
	for _, action := range actions {
		hf.AddArgument(action)
	}
}

// =======================
// Help-formatting methods
// =======================

func (hf *HelpFormatter) FormatHelp(args ...any) string {
	help := hf.RootSection.FormatHelp()
	if help != "" {
		help = hf.LongBreakMatcher.ReplaceAllString(help, "\n\n")
		help = strings.Trim(help, "\n") + "\n"
	}
	return help
}

func (hf *HelpFormatter) JoinParts_(partStrings []string) string {
	parts := []string{}
	for _, part := range partStrings {
		if part != "" && part != argparse.SUPPRESS {
			parts = append(parts, part)
		}
	}
	return strings.Join(parts, "")
}

// FIXME: not done
func (hf *HelpFormatter) FormatUsage_(usage string, actions []argparse.ActionInterface, groups []argparse.ActionsContainerInterface, prefix string) string {
	if prefix != "" {
		prefix = "usage: "
	}

	// if usage is specified, use that
	if usage != "" {
		usage = fmt.Sprintf(usage, hf.Prog_)
	} else if usage == "" && len(actions) == 0 {
		// if no optionals or positionals are available, usage is just prog
		usage = fmt.Sprintf("%(prog)s", hf.Prog_)
	} else if usage == "" {
		prog := fmt.Sprintf("%(prog)s", hf.Prog_)

		// split optionals from positionals
		optionals := []argparse.ActionInterface{}
		positionals := []argparse.ActionInterface{}

		for _, action := range actions {
			if action.Struct().OptionStrings != nil {
				optionals = append(optionals, action)
			} else {
				positionals = append(positionals, action)
			}
		}

		// build full usage string
		format := hf.FormatActionUsage
		actionUsage := format(append(optionals, positionals...), groups)
		sList := []string{}
		if prog != "" {
			sList = append(sList, prog)
		}
		if actionUsage != "" {
			sList = append(sList, actionUsage)
		}
		usage = strings.Join(sList, " ")

		// wrap the usage parts if it's too long
		textWidth := hf.Width - hf.CurrentIndent
		if len(prefix)+len(usage) > textWidth {
			// break usage into wrappable parts
			optParts := hf.GetActionsUsageParts(optionals, groups)
			posParts := hf.GetActionsUsageParts(positionals, groups)

			// helper for wrapping lines
			getLines := func(parts []string, indent string, prefix string) []string {
				var lines []string
				var line []string
				var lineLen int
				indentLength := len(indent)
				if prefix != "" {
					lineLen = len(prefix) - 1
				} else {
					lineLen = indentLength - 1
				}
				for _, part := range parts {
					if lineLen+1+len(part) > textWidth && len(line) > 0 {
						lines = append(lines, indent+strings.Join(line, " "))
						line = []string{}
						lineLen = indentLength - 1
					}
					line = append(line, part)
					lineLen += len(part) + 1
				}
				if len(line) > 0 {
					lines = append(lines, indent+strings.Join(line, " "))
				}
				if prefix != "" {
					lines[0] = lines[0][indentLength:]
				}
				return lines
			}

			// if prog is short, follow it with optionals or positionals
			var lines []string
			if len(prefix)+len(prog) <= int(0.75*float64(textWidth)) {
				indent := strings.Repeat(" ", len(prefix)+len(prog)+1)
				if len(optParts) > 0 {
					lines = getLines(append([]string{prog}, optParts...), indent, prefix)
					lines = append(lines, getLines(posParts, indent, "")...)
				} else if len(posParts) > 0 {
					lines = getLines(append([]string{prog}, posParts...), indent, prefix)
				} else {
					lines = []string{prog}
				}
			} else {
				// if prog is long, put it on its own line
				indent := strings.Repeat(" ", len(prefix))
				parts := append(optParts, posParts)
				lines = getLines(parts, indent, "")
				if len(lines) > 1 {
					lines = []string{}
					lines = append(lines, getLines(optParts, indent, "")...)
					lines = append(lines, getLines(posParts, indent, "")...)
				}
				lines = append([]string{prog}, lines...)
			}

			// join lines into usage
			usage = strings.Join(lines, "\n")
		}
	}
	// prefix with 'usage:'
	return fmt.Sprintf("%s%s\n\n", prefix, usage)
}

// // NewHelpFormatter creates a new instance of HelpFormatter
// func NewHelpFormatter(prog string, indentIncrement, maxHelpPosition, width int) *HelpFormatter {
// 	// Default width setting
// 	if width == 0 {
// 		width = getTerminalWidth() - 2
// 	}

// 	// indentIncrement = 2 //

// 	// Ensure maxHelpPosition is not too large
// 	maxHelpPosition = min(maxHelpPosition, max(width-20, indentIncrement*2))

// 	// Create a new HelpFormatter
// 	return &HelpFormatter{
// 		Prog:              prog,
// 		IndentIncrement:   indentIncrement,
// 		MaxHelpPosition:   maxHelpPosition,
// 		Width:             width,
// 		CurrentIndent:     0,
// 		Level:             0,
// 		ActionMaxLength:   0,
// 		RootSection:       &Section{Formatter: nil}, // Root section
// 		CurrentSection:    &Section{Formatter: nil}, // Current section
// 		WhitespaceMatcher: regexp.MustCompile(`\s+`),
// 		LongBreakMatcher:  regexp.MustCompile(`\n\n\n+`),
// 	}
// }

// // getTerminalWidth attempts to get the terminal's width, similar to shutil.get_terminal_size()
// func getTerminalWidth() int {
// 	return 80
// }

// min returns the smaller of two integers
func min(nums ...int) int {
	if len(nums) == 0 {
		panic("min: no arguments provided")
	}

	minVal := nums[0]
	for _, num := range nums[1:] {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

// max returns the larger of two integers
func max(nums ...int) int {
	if len(nums) == 0 {
		panic("max: no arguments provided")
	}

	maxVal := nums[0]
	for _, num := range nums[1:] {
		if num > maxVal {
			maxVal = num
		}
	}
	return maxVal
}

// // FormatHelp formats the help message for the command.
// func (hf *HelpFormatter) FormatHelp() string {
// 	// Implement formatting logic here (simplified for demonstration)
// 	return fmt.Sprintf("Usage: %s [options]", hf.Prog)
// }

// // Repr returns a string representation of the HelpFormatter
// func (hf *HelpFormatter) Repr() string {
// 	return fmt.Sprintf("%s(indentIncrement=%d, maxHelpPosition=%d, width=%d)", hf.Prog, hf.IndentIncrement, hf.MaxHelpPosition, hf.Width)
// }

// // _indent increments the current indent and level
// func (hf *HelpFormatter) indent() {
// 	hf.CurrentIndent += hf.IndentIncrement
// 	hf.Level++
// }

// // _dedent decrements the current indent and level
// func (hf *HelpFormatter) dedent() {
// 	hf.CurrentIndent -= hf.IndentIncrement
// 	if hf.CurrentIndent < 0 {
// 		panic("Indent decreased below 0.")
// 	}
// 	hf.Level--
// }

// func (hf *HelpFormatter) addItem(funcToAdd func(args ...interface{}) string, args ...interface{}) {
// 	hf.CurrentSection.AddItem(funcToAdd, args)
// }

// // ========================
// // Message building methods
// // ========================

// // startSection starts a new section with a given heading.
// func (hf *HelpFormatter) startSection(heading string) {
// 	hf.indent()
// 	section := &Section{
// 		Formatter: hf,
// 		Parent:    hf.CurrentSection,
// 		Heading:   heading,
// 		Items:     []SectionItem{},
// 	}
// 	hf.CurrentSection.Items = append(hf.CurrentSection.Items, SectionItem{
// 		FormatFunc: section.formatHelp,
// 		Args:       nil,
// 	})
// 	hf.CurrentSection = section
// }

// // endSection ends the current section and dedents.
// func (hf *HelpFormatter) endSection() {
// 	hf.CurrentSection = hf.CurrentSection.Parent
// 	hf.dedent()
// }

// // addText adds a plain text item to the current section.
// func (hf *HelpFormatter) addText(text string) {
// 	if text != argparse.SUPPRESS && text != "" {
// 		hf.CurrentSection.Items = append(hf.CurrentSection.Items, SectionItem{
// 			FormatFunc: hf.formatText,
// 			Args:       []interface{}{text},
// 		})
// 	}
// }

// // addUsage adds a usage item to the current section.
// func (hf *HelpFormatter) addUsage(usage string, actions []string, groups []string, prefix string) {
// 	if usage != argparse.SUPPRESS {
// 		// Example: Use the data to format the usage string in a more complex manner
// 		usageText := fmt.Sprintf("Usage: %s %s", usage, strings.Join(actions, " "))
// 		hf.CurrentSection.Items = append(hf.CurrentSection.Items, SectionItem{
// 			FormatFunc: hf.formatUsage,
// 			Args:       []interface{}{usageText},
// 		})
// 	}
// }

// // formatText formats the text item.
// func (hf *HelpFormatter) formatText(args ...interface{}) string {
// 	return fmt.Sprintf("%s", args[0])
// }

// // formatUsage formats the usage item.
// func (hf *HelpFormatter) formatUsage(args ...interface{}) string {
// 	return fmt.Sprintf("%s", args[0])
// }
