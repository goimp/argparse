package argparse

import (
	"argparse/attribute_holder"
	"fmt"
	"reflect"
	"strings"
)

// Namespace extends AttributeHolder to store attributes and provides functionality
// for comparison, existence check, and string representation.
type Namespace struct {
	attribute_holder.AttributeHolder
	attributes map[string]any
}

// NewNamespace creates a new Namespace with the given attributes.
func NewNamespace(attributes map[string]any) *Namespace {
	return &Namespace{
		attributes: attributes,
	}
}

// Set adds or updates an attribute in the Namespace.
func (n *Namespace) Set(name string, value any) {
	n.attributes[name] = value
}

// Get retrieves the value of an attribute from the Namespace.
func (n *Namespace) Get(name string) (any, bool) {
	val, found := n.attributes[name]
	return val, found
}

// Equals compares two Namespace objects for equality based on attribute names and values.
func (n *Namespace) Equals(other *Namespace) bool {
	if other == nil {
		return false
	}
	return reflect.DeepEqual(n.attributes, other.attributes)
}

// Contains checks if an attribute exists in the Namespace.
func (n *Namespace) Contains(key string) bool {
	_, found := n.attributes[key]
	return found
}

// Repr returns a string representation of the Namespace with its attributes.
func (n *Namespace) Repr() string {
	// Collect the attribute name and value pairs.
	var argStrings []string
	for key, value := range n.attributes {
		argStrings = append(argStrings, fmt.Sprintf("%s=%v", key, value))
	}

	return fmt.Sprintf("Namespace(%s)", strings.Join(argStrings, ", "))
}
