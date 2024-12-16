package action

import (
	"argparse"
	"argparse/namespace"
	"fmt"
	"strings"
)

// SubParsersAction handles subcommands and their parsers
type SubParsersAction struct {
	Action         // Embedding Action to reuse functionality
	ProgPrefix     string
	ParserClass    func(kwargs any) (any, error)
	NameParserMap  map[string]any
	ChoicesActions []*ChoicesPseudoAction
	Deprecated     map[string]struct{}
}

type ChoicesPseudoAction struct {
	Action
}

// Constructor-like function to create a new ChoicesPseudoAction
func NewChoicesPseudoAction(name string, aliases []string, help string) *ChoicesPseudoAction {
	// Set metavar and dest based on the name
	metavar := name
	dest := name

	// If aliases are provided, append them to metavar
	if len(aliases) > 0 {
		metavar += fmt.Sprintf(" (%s)", strings.Join(aliases, ", "))
	}

	// Initialize Action struct via embedding
	return &ChoicesPseudoAction{
		Action: Action{
			OptionStrings: []string{}, // Empty slice as per the original code
			Dest:          dest,
			Help:          help,
			Metavar:       metavar,
		},
	}
}

// NewSubParsersAction creates a new SubParsersAction instance
func NewSubParsersAction(
	optionStrings []string,
	prog string,
	parserClass func(kwargs any) (any, error),
	dest string,
	required bool,
	help string,
	metavar string,
) (*SubParsersAction, error) {
	return &SubParsersAction{
		Action: Action{
			OptionStrings: optionStrings,
			Dest:          dest,
			Nargs:         argparse.PARSER,
			Required:      required,
			Help:          help,
			Metavar:       metavar,
		},
		ProgPrefix:     prog,
		ParserClass:    parserClass,
		NameParserMap:  make(map[string]any),
		ChoicesActions: []*ChoicesPseudoAction{},
		Deprecated:     make(map[string]struct{}),
	}, nil
}

func (p *SubParsersAction) AddParser(name string, deprecated bool, kwargs map[string]any) (any, error) {
	if _, exist := kwargs["prog"]; exist {
		delete(kwargs, "help")
	}
	aliases, exist := kwargs["aliases"]
	if !exist {
		aliases = map[string]string{}
	}

	aliasesLi, ok := aliases.([]string)
	if !ok {
		return nil, fmt.Errorf("wrong aliases type: %v", aliasesLi)
	}

	if exist {
		delete(kwargs, "aliases")
	}
	for _name := range p.NameParserMap {
		if name == _name {
			return nil, fmt.Errorf("conflicting subparser: %s", name)
		}
	}
	for _, alias := range aliasesLi {
		for _name := range p.NameParserMap {
			if alias == _name {
				return nil, fmt.Errorf("conflicting subparser alias: %s", alias)
			}
		}
	}

	// create a pseudo-action to hold the choice help
	var choiceAction *ChoicesPseudoAction = nil
	if help, exist := kwargs["help"]; exist {
		switch t := help.(type) {
		case string:
			choiceAction = NewChoicesPseudoAction(name, aliasesLi, help.(string))
		default:
			return nil, fmt.Errorf("wrong help type for action: %s, %v", name, t)
		}
	}

	// create the parser and add it to the map
	parser, err := p.ParserClass(kwargs)
	if err != nil {
		return nil, fmt.Errorf("error on parser creating")
	}
	if choiceAction != nil {
		// parser.CheckHelp(choiceAction)
		fmt.Printf("WARNING on CheckHelp")
	}
	p.NameParserMap[name] = parser

	//  make parser available under aliases also

	for _, alias := range aliasesLi {
		p.NameParserMap[alias] = parser
	}

	if deprecated {
		// p.Deprecated = append(p.Deprecated, name)
		fmt.Printf("WARNING on deprecated")
	}

	return nil, nil
}

func (p *SubParsersAction) GetSubactions() []*ChoicesPseudoAction {
	return p.ChoicesActions
}

func (p *SubParsersAction) Call(parser any, namespace namespace.Namespace, values []any, optionString string) {
	parserName := values[0].(string)
	argStrings := values[1:]

	// set the parser name if requested
	if p.Dest != argparse.SUPPRESS {
		namespace.Set(p.Dest, parserName)
	}

	// select the parser
	// subparser, exist := p.NameParserMap[parserName]
	// if !exist {
	// 	keys := make([]string, len(p.NameParserMap))
	// 	joinedChoices := strings.Join(keys, ", ")
	// 	panic(fmt.Sprintf("unknown parser %s (choices: %s)", joinedChoices, parserName))
	// }

	for depr := range p.Deprecated {
		if depr == parserName {
			// parser.Warning(fmt.Sprintf("command '%(parser_name)s' is deprecated", parserName))
			fmt.Printf("command %s is deprecated\n", parserName)
		}
	}

	// parse all the remaining options into the namespace
	// store any unrecognized options on the object, so that the top
	// level parser can decide what to do with them

	// In case this subparser defines new defaults, we parse them
	// in a new namespace object and then update the original
	// namespace for the relevant parts.

	// subnamespace, argStrings := subparser.parseKnownArgs(argStrings, nil)
	// for key, value := range subnamespace {
	// 	namespace.Set(key, value)
	// }

	if argStrings != nil {
		if unrec, exist := namespace.Get(argparse.UNRECOGNIZED_ARGS_ATTR); !exist {
			namespace.Set(argparse.UNRECOGNIZED_ARGS_ATTR, make(map[string]any))
		} else {
			switch t := unrec.(type) {
			case []any:
				newUnrec := unrec.([]any)
				newUnrec = append(newUnrec, argStrings)
				namespace.Set(argparse.UNRECOGNIZED_ARGS_ATTR, newUnrec)
			default:
				fmt.Printf("Wrong unrec type %v\n", t)
			}
		}
	}
}
