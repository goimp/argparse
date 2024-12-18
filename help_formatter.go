package argparse

import (
	"fmt"
	"reflect"
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
	AddUsage(string, []ActionInterface, []ActionsContainerInterface, string)

	AddArgument(ActionInterface)
	AddArguments([]ActionInterface)

	FormatHelp(...any) string
	JoinParts_([]string) string
	FormatUsage_(string, []ActionInterface, []ActionsContainerInterface, string) string
	FormatActionsUsage_([]ActionInterface, []ActionsContainerInterface) string
	GetActionsUsageParts_([]ActionInterface, []ActionsContainerInterface) []string

	FormatText_(string) string
	FormatAction_(ActionInterface) string
	FormatActionInvocation_(ActionInterface) string

	MetaVarFormatter_(ActionInterface, string) func(int) []string
	FormatArgs_(ActionInterface, string) string
	ExpandHelp_(ActionInterface, string) string
	IterIndentedSubactions_(ActionInterface) []ActionInterface
	SplitLines_(string, int) []string
	FillText_(string, int, string) string

	GetHelpString_(action ActionInterface) string
	GetDefaultMetaVarForOptional_(action ActionInterface) string
	GetDefaultMetaVarForPositional_(action ActionInterface) string
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
	if s.Heading != SUPPRESS && s.Heading != "" {
		currentIndent := s.Formatter.Struct().CurrentIndent_
		headingText := formatKeys("%(heading)s:", map[string]any{"heading": s.Heading})
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
	Width_            int
	CurrentIndent_    int
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
		Width_:          width,
		CurrentIndent_:  0,
		Level:           0,
		ActionMaxLength: 0,

		WhitespaceMatcher: regexp.MustCompile(`\s+`),
		LongBreakMatcher:  regexp.MustCompile(`\n\n\n+`),
	}

	rootSection := &Section_{
		Formatter: formatter,
		Parent:    nil,
	}

	formatter.RootSection = rootSection
	formatter.CurrentSection = formatter.RootSection

	return formatter
}

func (hf *HelpFormatter) Struct() *HelpFormatter {
	return hf
}

func (hf *HelpFormatter) Indent_() {
	hf.CurrentIndent_ += hf.IndentIncrement
	hf.Level++
}

func (hf *HelpFormatter) Dedent_() {
	hf.CurrentIndent_ -= hf.IndentIncrement
	if hf.CurrentIndent_ < 0 {
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
	if text != SUPPRESS && text != "" {
		hf.AddItem_(hf.FormatHelp, text)
	}
}

func (hf *HelpFormatter) AddUsage(usage string, actions []ActionInterface, groups []ActionsContainerInterface, prefix string) {
	if usage != SUPPRESS {
		args := []any{usage, actions, groups, prefix}
		hf.AddItem_(hf.FormatHelp, args...)
	}
}

func (hf *HelpFormatter) AddArgument(action ActionInterface) {
	if action.Struct().Help != SUPPRESS {

		// find all invocations
		getInvocation := hf.FormatActionInvocation_
		invocationLengths := []int{len(getInvocation(action)) + hf.CurrentIndent_}
		for _, subaction := range hf.IterIndentedSubactions_(action) {
			invocationLengths = append(invocationLengths, len(getInvocation(subaction))+hf.CurrentIndent_)
		}

		// update the maximum item length
		actionLength := max(invocationLengths...)
		hf.ActionMaxLength = max(hf.ActionMaxLength, actionLength)
	}
}

func (hf *HelpFormatter) AddArguments(actions []ActionInterface) {
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
		if part != "" && part != SUPPRESS {
			parts = append(parts, part)
		}
	}
	return strings.Join(parts, "")
}

func (hf *HelpFormatter) FormatUsage_(usage string, actions []ActionInterface, groups []ActionsContainerInterface, prefix string) string {
	if prefix != "" {
		prefix = "usage: "
	}

	// if usage is specified, use that
	if usage != "" {
		usage = fmt.Sprintf(usage, hf.Prog_)
	} else if usage == "" && len(actions) == 0 {
		// if no optionals or positionals are available, usage is just prog
		usage = formatKeys("%(prog)s", map[string]any{"prog": hf.Prog_})
	} else if usage == "" {
		prog := formatKeys("%(prog)s", map[string]any{"prog": hf.Prog_})

		// split optionals from positionals
		optionals := []ActionInterface{}
		positionals := []ActionInterface{}

		for _, action := range actions {
			if action.Struct().OptionStrings != nil {
				optionals = append(optionals, action)
			} else {
				positionals = append(positionals, action)
			}
		}

		// build full usage string
		format := hf.FormatActionsUsage_
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
		textWidth := hf.Width_ - hf.CurrentIndent_
		if len(prefix)+len(usage) > textWidth {
			// break usage into wrappable parts
			optParts := hf.GetActionsUsageParts_(optionals, groups)
			posParts := hf.GetActionsUsageParts_(positionals, groups)

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
				parts := append(optParts, posParts...)
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

func (hf *HelpFormatter) FormatActionsUsage_(actions []ActionInterface, groups []ActionsContainerInterface) string {
	return strings.Join(hf.GetActionsUsageParts_(actions, groups), " ")
}

func (hf *HelpFormatter) GetActionsUsageParts_(actions []ActionInterface, groups []ActionsContainerInterface) []string {
	// FIXME: Not done
	return []string{}
}

func (hf *HelpFormatter) FormatText_(text string) string {
	text = formatKeys(text, map[string]any{"prog": hf.Prog_})
	textWidth := max(hf.Width_-hf.CurrentIndent_, 11)
	indent := strings.Repeat(" ", hf.CurrentIndent_)
	return hf.FillText_(text, textWidth, indent)
}

func (hf *HelpFormatter) FormatAction_(action ActionInterface) string {
	// FIXME: Not done
	return ""
}

func (hf *HelpFormatter) FormatActionInvocation_(action ActionInterface) string {
	// FIXME: Not done
	return ""
}

func (hf *HelpFormatter) MetaVarFormatter_(action ActionInterface, defaultMetaVar string) func(int) []string {
	// FIXME: Not done
	// return nil
	var result string
	if action.Struct().MetaVar != "" {
		result = action.Struct().MetaVar.(string)
	} else if action.Struct().Choices != nil {
		var strChoices []string
		for _, choice := range action.Struct().Choices {
			strChoices = append(strChoices, fmt.Sprintf("%s", choice))
		}
		result = fmt.Sprintf("%s", strings.Join(strChoices, ","))
	} else {
		result = defaultMetaVar
	}

	// format := func(tupleSize int) interface{} {
	// 	// Check if result is already a tuple (slice in Go)

	// 	result
	// }
	// return format

	format := func(tupleSize int) []string {
		return []string{result}
	}
	return format
}

func (hf *HelpFormatter) FormatArgs_(action ActionInterface, defaultMetaVar string) string {
	// FIXME: Not done
	return ""
}

func (hf *HelpFormatter) ExpandHelp_(action ActionInterface, defaultMetaVar string) string {
	helpString := hf.GetHelpString_(action)
	if !strings.Contains(helpString, "%") {
		return helpString
	}
	// params := map[string]interface{}{
	// 	"prog": hf.Prog_,
	// }
	// for key, value := range action.Extra {
	// 	params[key] = value
	// }
	// FIXME: Not done
	return ""
}

func (hf *HelpFormatter) IterIndentedSubactions_(action ActionInterface) []ActionInterface {
	getSubActions := action.GetSubActions
	subactions := getSubActions()
	if subactions != nil {
		hf.Indent_()
		subactions = append(subactions, subactions...)
		hf.Dedent_()
	}
	return subactions
}

func (hf *HelpFormatter) SplitLines_(text string, width int) []string {
	// FIXME: Not done
	// if width <= 0 {
	// 	panic("wrapText: width must be greater than 0")
	// }

	var lines []string
	// var line strings.Builder
	// currentWidth := 0

	// words := strings.Fields(text) // Split text into words
	// for _, word := range words {
	// 	wordLen := utf8.RuneCountInString(word)

	// 	// If adding the word exceeds the width, finalize the current line
	// 	if currentWidth+wordLen+1 > width && currentWidth > 0 {
	// 		lines = append(lines, line.String())
	// 		line.Reset()
	// 		currentWidth = 0
	// 	}

	// 	// Add the word to the line
	// 	if currentWidth > 0 {
	// 		line.WriteString(" ")
	// 		currentWidth++
	// 	}
	// 	line.WriteString(word)
	// 	currentWidth += wordLen
	// }

	// // Add the last line if there's remaining text
	// if line.Len() > 0 {
	// 	lines = append(lines, line.String())
	// }

	return lines
}

func (hf *HelpFormatter) FillText_(text string, width int, indent string) string {
	// NOTE: 'indent' omited

	textWrap := func(text string, width int) string {
		if width <= 0 {
			return text
		}

		var result strings.Builder
		var line strings.Builder

		words := strings.Fields(text)

		for _, word := range words {
			// Check if adding the word exceeds the width
			if line.Len()+len(word)+1 > width { // +1 accounts for the space
				// Append the current line to the result
				result.WriteString(strings.TrimSpace(line.String()) + "\n")
				line.Reset()
			}
			// Append the word to the current line
			line.WriteString(word + " ")
		}

		// Add any remaining text in the line to the result
		if line.Len() > 0 {
			result.WriteString(strings.TrimSpace(line.String()))
		}

		return result.String()
	}

	return textWrap(text, width)
}

func (hf *HelpFormatter) GetHelpString_(action ActionInterface) string {
	return action.Struct().Help
}

func (hf *HelpFormatter) GetDefaultMetaVarForOptional_(action ActionInterface) string {
	return strings.ToUpper(action.Struct().Dest)
}

func (hf *HelpFormatter) GetDefaultMetaVarForPositional_(action ActionInterface) string {
	return action.Struct().Dest
}

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

type RawDescriptionHelpFormatter struct {
	*HelpFormatter
}

func (fh *RawDescriptionHelpFormatter) FillText_(text string, width int, indent string) string {
	var builder strings.Builder

	// Normalize line endings for consistent processing
	lines := strings.SplitAfter(text, "\n") // Split while retaining line breaks

	for _, line := range lines {
		builder.WriteString(indent) // Add the indent
		builder.WriteString(line)   // Add the original line (with line break if present)
	}

	return builder.String()
}

type RawTextHelpFormatter struct {
	*HelpFormatter
}

func (fh *RawTextHelpFormatter) FillText_(text string, width int, indent string) string {
	var builder strings.Builder

	// Normalize line endings for consistent processing
	lines := strings.SplitAfter(text, "\n") // Split while retaining line breaks

	for _, line := range lines {
		builder.WriteString(indent) // Add the indent
		builder.WriteString(line)   // Add the original line (with line break if present)
	}

	return builder.String()
}

func (fh *RawTextHelpFormatter) SplitLines_(text string, width int) []string {
	return strings.SplitAfter(text, "\n")
}

type ArgumentDefaultsHelpFormatter struct {
	*HelpFormatter
}

func (fh *ArgumentDefaultsHelpFormatter) GetHelpString_(action ActionInterface) string {
	// Extract the action structure for easier reference

	containsHelper := func(slice []string, value string) bool {
		for _, v := range slice {
			if v == value {
				return true
			}
		}
		return false
	}

	actionStruct := action.Struct()

	// Initialize help text
	help := actionStruct.Help
	if help == "" {
		help = ""
	}

	// Check if help already contains %(default)
	if !strings.Contains(help, "%(default)") {
		// Add default value if applicable
		if actionStruct.Default != SUPPRESS {
			defaultingNargs := []string{OPTIONAL, ZERO_OR_MORE}

			// Handle Nargs as either string or number
			switch t := actionStruct.Nargs.(type) {
			case string:
				// If Nargs is a string, check if it matches defaultingNargs
				if len(actionStruct.OptionStrings) > 0 || containsHelper(defaultingNargs, t) {
					help += fmt.Sprintf(" (default: %s)", actionStruct.Default)
				}
			case int:
				// If Nargs is a number, no need to compare with defaultingNargs
				if len(actionStruct.OptionStrings) > 0 {
					help += fmt.Sprintf(" (default: %s)", actionStruct.Default)
				}
			default:
				// Handle unexpected types gracefully (optional)
				// help += fmt.Sprintf(" (default: %s)", actionStruct.Default)
			}
		}
	}

	return help
}

type MetaVarTypeHelpFormatter struct {
	*HelpFormatter
}

func (fh *MetaVarTypeHelpFormatter) GetDefaultMetaVarForOptional_(action ActionInterface) string {
	return reflect.TypeOf(action.Struct().Type).String()
}

func (fh *MetaVarTypeHelpFormatter) GetDefaultMetaVarForPositional_(action ActionInterface) string {
	return reflect.TypeOf(action.Struct().Type).String()
}

func formatKeys(str string, values map[string]any) string {
	// Create a regex pattern to match placeholders like %(key) or %(key)s
	re := regexp.MustCompile(`%\((\w+)\)(s?)`)
	// Iterate over matches
	return re.ReplaceAllStringFunc(str, func(match string) string {
		// Extract the key from the match (exclude the surrounding '%(' and ')')
		key := match[2 : len(match)-2]
		// Check if the key exists in the values map
		if value, exists := values[key]; exists {
			// Replace with the formatted value (using %v to convert to string)
			return fmt.Sprintf("%v", value)
		}
		// If the key is not found, return the match unchanged
		return match
	})
}
