package runtime_util

import (
	"fmt"
	"runtime"
	"strings"
)

// GetFunctionName returns the name of the calling function
// skip: number of stack frames to skip (0 = current function, 1 = caller, etc.)
func GetFunctionName(skip int) string {
	pc, _, _, _ := runtime.Caller(skip + 1)
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return "[Unknown]"
	}

	// Get full function name (e.g., "math_quest_be/internal/auth/usecase.(*Usecase).SignUp")
	fullName := fn.Name()

	// Extract just the function name (e.g., "SignUp")
	parts := strings.Split(fullName, ".")
	if len(parts) > 0 {
		funcName := parts[len(parts)-1]
		return fmt.Sprintf("[%s]", funcName)
	}

	return "[Unknown]"
}
