package golang

import (
	"fmt"
)

func what(types interface{}) string {
	//return fmt.Sprintf("%T", types)

	switch v := types.(type) {
	case int:
		return fmt.Sprintf("%T", v)
	case string:
		return fmt.Sprintf("%T", v)
	case bool:
		return fmt.Sprintf("%T", v)
	default:
		return fmt.Sprintf("%T", v)
	}
}
