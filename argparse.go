package argparse

const (
	SUPPRESS               = "==SUPPRESS=="
	OPTIONAL               = "?"
	ZERO_OR_MORE           = "*"
	ONE_OR_MORE            = "+"
	PARSER                 = "A..."
	REMAINDER              = "..."
	UNRECOGNIZED_ARGS_ATTR = "_unrecognized_args"
)

func ProgName(prog string) string {
	return ""
}
