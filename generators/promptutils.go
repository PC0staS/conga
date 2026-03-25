package generators

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

// AskDomain asks for the main domain (shared)
func AskDomain() string {
    prompt := promptui.Prompt{
        Label:   "Main domain",
        Default: "localhost",
    }

    domain, err := runPromptOrExit(prompt)
    if err != nil {
        fmt.Println("Error:", err)
        return ""
    }

    return domain
}

// AskGenerateCerts asks if we should generate certs with mkcert (shared)
func AskGenerateCerts() bool {
    prompt := promptui.Select{
        Label: "Generate local certificates with mkcert? (Needs sudo for system dirs, or will generate locally with warning)",
        Items: []string{"Yes", "No"},
    }

    _, result, err := runSelectOrExit(prompt)
    if err != nil {
        fmt.Println("Error:", err)
        return false
    }

    return result == "Yes"
}

// GenerateCerts runs mkcert to produce cert and key for domain into outDir (shared)
func GenerateCerts(domain, outDir string) error {
    if _, err := exec.LookPath("mkcert"); err != nil {
        return fmt.Errorf("mkcert not found in PATH: %v", err)
    }

    if err := os.MkdirAll(outDir, 0755); err != nil {
        return fmt.Errorf("cannot create out dir: %v", err)
    }

    cmd := exec.Command("mkcert", "-install")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("mkcert -install failed: %v", err)
    }

    certFile := filepath.Join(outDir, domain+".crt")
    keyFile := filepath.Join(outDir, domain+".key")

    hosts := []string{domain, "localhost", "127.0.0.1"}
    args := append([]string{"-cert-file", certFile, "-key-file", keyFile}, hosts...)
    cmd = exec.Command("mkcert", args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        return fmt.Errorf("mkcert generation failed: %v", err)
    }

    return nil
}

// AskNumberOfRoutes asks how many routes to configure (shared)
func AskNumberOfRoutes() int {
    prompt := promptui.Prompt{
        Label:   "Number of routes (paths)",
        Default: "1",
        Validate: func(input string) error {
            num := 0
            _, err := fmt.Sscanf(input, "%d", &num)
            if err != nil {
                return fmt.Errorf("must be a number")
            }
            if num < 1 {
                return fmt.Errorf("must be at least 1")
            }
            return nil
        },
    }

    result, err := runPromptOrExit(prompt)
    if err != nil {
        fmt.Println("Error:", err)
        return 0
    }

    var num int
    fmt.Sscanf(result, "%d", &num)
    return num
}

// AskTLS asks if using TLS and optionally where certs live. Returns (useTLS, certPath).
// Pass a sensible defaultCertPath (without suffix) like "/etc/nginx/certs/example.com".
func AskTLS(domain, defaultCertPath string) (bool, string) {
    prompt := promptui.Select{
        Label: "Use TLS (provide cert/key)?",
        Items: []string{"Yes", "No"},
    }

    _, result, err := runSelectOrExit(prompt)
    if err != nil {
        fmt.Println("Error:", err)
        return false, ""
    }

    if result == "Yes" {
        if AskGenerateCerts() {
            // Try to generate certs directly in the desired destination directory.
            desiredDir := filepath.Dir(defaultCertPath)
            // Attempt to create destination dir (may fail if no permissions)
            if err := os.MkdirAll(desiredDir, 0o755); err == nil {
                if err := GenerateCerts(domain, desiredDir); err == nil {
                    fmt.Printf("\n✅ Certificates created at %s (referenced as %s.crt/%s.key)\n", desiredDir, defaultCertPath, defaultCertPath)
                    return true, defaultCertPath
                }
                // Fall through to attempt local generation below
            }

            // Fallback: generate in local ./certs if writing to desiredDir failed
            outDir := filepath.Join(".", "certs")
            if err := GenerateCerts(domain, outDir); err != nil {
                fmt.Printf("\n❌ Error generating certs: %v\n", err)
                fmt.Println("Continuing; ensure certs exist at the specified cert path.")
            } else {
                // Inform user that certs were created locally but config will reference defaultCertPath
                fmt.Printf("\n⚠️ Certificates created at %s but configuration references %s.crt/%s.key.\n", outDir, defaultCertPath, defaultCertPath)
                fmt.Println("If you need the files at the referenced path, move them with sudo or run conga with appropriate permissions.")
                return true, defaultCertPath
            }
        }

        certPrompt := promptui.Prompt{
            Label:   "Cert path without suffix (e.g. /etc/ssl/certs/example.com)",
            Default: defaultCertPath,
        }
        certPath, err := runPromptOrExit(certPrompt)
        if err != nil {
            fmt.Println("Error:", err)
            return true, defaultCertPath
        }
        return true, certPath
    }

    return false, ""
}
