package generators

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

// NormalizeOutputPath normalizes an output filename provided by the user.
// - treats leading "/build/" as the repo-local "./build/" directory
// - expands leading ~ to the user's home directory
// - ensures parent directories exist
func NormalizeOutputPath(filename string) (string, error) {
    // If user provided a path starting with "/build/" assume they meant the repo's build/ directory
    if strings.HasPrefix(filename, "/build/") {
        filename = filepath.Join(".", filename[1:])
    }

    // Expand ~ to home directory
    if strings.HasPrefix(filename, "~") {
        home, err := os.UserHomeDir()
        if err != nil {
            return "", fmt.Errorf("error expanding ~: %v", err)
        }
        filename = filepath.Join(home, filename[1:])
    }

    // Ensure parent directory exists
    dir := filepath.Dir(filename)
    if dir != "." {
        if err := os.MkdirAll(dir, 0o755); err != nil {
            return "", fmt.Errorf("error creating directories for output: %v", err)
        }
    }

    return filename, nil
}

func askOutput(defaultValue string) string {
	prompt := promptui.Prompt{
		Label:   "Output filename",
		Default: defaultValue,
	}
	
	output, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return defaultValue
	}
	
	return output
}