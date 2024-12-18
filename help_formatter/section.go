package help_formatter

// import (
// 	"argparse"
// 	"fmt"
// 	"strings"
// )

// // Section represents a section of help text in the formatter.
// type Section struct {
// 	Formatter *HelpFormatter
// 	Parent    *Section
// 	Heading   string
// 	Items     []SectionItem
// }

// // SectionItem is a placeholder for the items in a section.
// // In this case, each item represents a function and its arguments for formatting.
// // SectionItem represents an item within a section, which could be text or other data.
// type SectionItem struct {
// 	FormatFunc func(args ...interface{}) string
// 	Args       []interface{}
// }

// // NewSection creates a new section with a given heading.
// func NewSection(formatter *HelpFormatter, parent *Section, heading string) *Section {
// 	return &Section{
// 		Formatter: formatter,
// 		Parent:    parent,
// 		Heading:   heading,
// 		Items:     []SectionItem{},
// 	}
// }

// // AddItem adds an item (function and arguments) to the section.
// func (s *Section) AddItem(funcToAdd func(args ...interface{}) string, args ...interface{}) {
// 	s.Items = append(s.Items, SectionItem{
// 		FormatFunc: funcToAdd,
// 		Args:       args,
// 	})
// }

// // Section methods to format the help text for a section.
// // formatHelp formats the help string for the section.
// func (s *Section) formatHelp(args ...interface{}) string {
// 	// Indent for the parent section
// 	if s.Parent != nil {
// 		s.Formatter.indent()
// 	}

// 	// Create a list of formatted items using the provided functions and arguments
// 	var formattedItems []string
// 	for _, item := range s.Items {
// 		formattedItems = append(formattedItems, item.FormatFunc(item.Args...))
// 	}

// 	// Dedent for the parent section after formatting
// 	if s.Parent != nil {
// 		s.Formatter.dedent()
// 	}

// 	// If no items were added, return an empty string
// 	if len(formattedItems) == 0 {
// 		return ""
// 	}

// 	// Format heading
// 	var heading string
// 	if s.Heading != "" && s.Heading != argparse.SUPPRESS {
// 		currentIndent := s.Formatter.CurrentIndent
// 		headingText := fmt.Sprintf("%s:", s.Heading)
// 		heading = fmt.Sprintf("%*s%s\n", currentIndent, "", headingText)
// 	}

// 	// Join the section with newline, heading, and formatted items
// 	return fmt.Sprintf("\n%s%s%s\n", heading, strings.Join(formattedItems, "\n"), "\n")
// }
