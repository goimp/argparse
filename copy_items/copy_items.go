package copy_items

import (
	"reflect"
)

// CopyItems creates a copy of the provided input.
// If the input is nil, it returns an empty slice.
// If the input is a slice, it returns a shallow copy of the slice.
// For other types, it uses reflection to attempt a shallow copy.
func CopyItems(items any) any {
	if items == nil {
		return []any{}
	}

	// Handle case where items is a slice
	if reflect.TypeOf(items).Kind() == reflect.Slice {
		v := reflect.ValueOf(items)
		copy := reflect.MakeSlice(v.Type(), v.Len(), v.Cap())
		reflect.Copy(copy, v)
		return copy.Interface()
	}

	// Fallback to shallow copy for other types
	// Using reflection for generality
	return shallowCopy(items)
}

// shallowCopy performs a shallow copy of non-slice types.
func shallowCopy(item any) any {
	v := reflect.ValueOf(item)
	if !v.IsValid() || v.Kind() == reflect.Ptr && v.IsNil() {
		return nil
	}

	// Ensure we handle copyable types
	if v.Kind() == reflect.Struct || v.Kind() == reflect.Map || v.Kind() == reflect.Array {
		copy := reflect.New(v.Type()).Elem()
		copy.Set(v)
		return copy.Interface()
	}

	// Return the item directly for immutable types
	return item
}

// func main() {
// 	// Example usage

// 	// Copy a slice
// 	originalSlice := []int{1, 2, 3}
// 	copiedSlice := CopyItems(originalSlice).([]int)
// 	fmt.Println("Original Slice:", originalSlice)
// 	fmt.Println("Copied Slice:", copiedSlice)

// 	// Copy a non-slice value
// 	originalValue := map[string]int{"a": 1, "b": 2}
// 	copiedValue := CopyItems(originalValue).(map[string]int)
// 	fmt.Println("Original Map:", originalValue)
// 	fmt.Println("Copied Map:", copiedValue)

// 	// Ensure independence of slices
// 	copiedSlice[0] = 99
// 	fmt.Println("Modified Copied Slice:", copiedSlice)
// 	fmt.Println("Original Slice After Modification:", originalSlice)
// }
