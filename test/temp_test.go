package argparse_test

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Action represents an action with a key and destination
type Action struct {
	Key  string `json:"key"`
	Dest string `json:"dest"`
}

// Group represents a group of items, which can be other groups or actions
type Group struct {
	Name      string `json:"name"`
	Items     []any  `json:"items"`
	Exclusive bool   `json:"exclusive"`
}

// GroupJSON represents a JSON-compatible structure for the Group with Actions
type GroupJSON struct {
	Name      string `json:"name"`
	Items     []any  `json:"items"`
	Exclusive bool   `json:"exclusive"`
}

// ActionJSON represents a JSON-compatible structure for the Action
type ActionJSON struct {
	Key  string `json:"key"`
	Dest string `json:"dest"`
}

func (g *Group) String() string {
	return ""
}

// Recursive function to collect data in JSON format
func (g *Group) IterRecursiveToJSON() ([]byte, error) {
	var items []any
	for _, item := range g.Items {
		switch t := item.(type) {
		case *Group:
			// Recursively call for subgroups
			subgroupJSON, err := t.IterRecursiveToJSON()
			if err != nil {
				return nil, err
			}
			var subgroup interface{}
			err = json.Unmarshal(subgroupJSON, &subgroup)
			if err != nil {
				return nil, err
			}
			items = append(items, subgroup)
		case *Action:
			// Create ActionJSON to store it in a JSON-compatible format
			action := ActionJSON{Key: t.Key, Dest: t.Dest}
			items = append(items, action)
		default:
			// Handle unexpected types
			items = append(items, fmt.Sprintf("unknown: %v", t))
		}
	}

	// Create GroupJSON to structure the output for the group
	groupJSON := GroupJSON{Name: g.Name, Items: items, Exclusive: g.Exclusive}
	return json.MarshalIndent(groupJSON, "", "  ")
}

// SaveGroupToJSON saves the group data as JSON to a file
func SaveGroupToJSON(g *Group, filename string) error {
	data, err := g.IterRecursiveToJSON()
	if err != nil {
		return err
	}

	// Write the JSON data to a file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (g *Group) IterRecursive(indent int) {
	sep := "  "
	ind := strings.Repeat(sep, indent)
	fmt.Printf("%s%-10s:\t%-10s\n", ind, "Group", g.Name)
	for _, item := range g.Items {
		switch t := item.(type) {
		case *Group:
			t.IterRecursive(indent + 1)
		case *Action:
			fmt.Printf("%s%-10s:\t%-10s\n", ind+sep, "Group", g.Name)
			fmt.Printf("%s%-10s:\t%-10s\n", ind+sep+sep, t.Key, t.Dest)
			fmt.Printf("%s%-10s:\t%-10s\n", ind+sep+sep, t.Key, t.Dest)
		default:
			fmt.Printf("%s %-10s:\t%v\n", ind, "unknown", t)
		}
	}
}

func TestTemp(t *testing.T) {

	args := &Group{
		Items: []any{
			&Action{
				Key:  "-v",
				Dest: "action_v",
			},
			&Group{
				Items: []any{
					&Action{
						Key:  "-a",
						Dest: "action_a",
					},
				},
			},
			&Group{
				Items: []any{
					&Action{
						Key:  "-b",
						Dest: "action_b",
					},
					&Action{
						Key:  "-c",
						Dest: "action_c",
					},
				},
				Exclusive: true,
			},
			&Group{
				Items: []any{
					&Group{
						Items: []any{
							&Action{
								Key:  "-d",
								Dest: "action_d",
							},
							&Action{
								Key:  "-e",
								Dest: "action_e",
							},
							"undefined",
						},
					},
					&Group{
						Items: []any{},
					},
				},
			},
		},
	}

	// args.IterRecursive(0)
	SaveGroupToJSON(args, "group_data.json")

}
