package namespace

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
	attributes map[string]interface{}
}

// NewNamespace creates a new Namespace with the given attributes.
func NewNamespace(attributes map[string]interface{}) *Namespace {
	return &Namespace{
		attributes: attributes,
	}
}

// Set adds or updates an attribute in the Namespace.
func (n *Namespace) Set(name string, value interface{}) {
	n.attributes[name] = value
}

// Get retrieves the value of an attribute from the Namespace.
func (n *Namespace) Get(name string) (interface{}, bool) {
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

// // Example usage
// func main() {
// 	// Create a new Namespace instance
// 	ns := NewNamespace(map[string]interface{}{"key1": "value1", "key2": 42})

// 	// Print the string representation
// 	fmt.Println(ns.Repr()) // Output: Namespace(key1=value1, key2=42)

// 	// Compare namespaces for equality
// 	ns2 := NewNamespace(map[string]interface{}{"key1": "value1", "key2": 42})
// 	fmt.Println("Namespaces are equal:", ns.Equals(ns2)) // Output: true

// 	// Check if a key exists
// 	fmt.Println("Contains 'key1':", ns.Contains("key1")) // Output: true

// 	// Retrieve a value by key
// 	if val, found := ns.Get("key2"); found {
// 		fmt.Println("key2:", val) // Output: key2: 42
// 	}

// 	// Add a new attribute
// 	ns.Set("key3", "new value")
// 	fmt.Println("Updated Namespace:", ns.Repr()) // Output: Namespace(key1=value1, key2=42, key3=new value)
// }
