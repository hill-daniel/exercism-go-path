// provides two for one functionality.
package twofer

import (
	"fmt"
)

// returns "One for <name>, one for me." string or "One for you..." if name is empty.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %v, one for me.", name)
}
