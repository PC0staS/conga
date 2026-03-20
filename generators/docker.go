package generators

import (
	"fmt"
	"os"

	"text/template"

	"github.com/manifoldco/promptui"
)
 const dockerComposeTemplate = `version: '{{ .Version }}'

services:
{{- range .Services }}

  {{ .Name }}:
    image: {{ .Image }}
{{- if .Ports }}
    ports:
{{- range .Ports }}
      - "{{ . }}"
{{- end }}
{{- end }}
{{- if .Volumes }}
    volumes:
{{- range .Volumes }}
      - {{ . }}
{{- end }}
{{- end }}
{{- if .EnvFile }}
    env_file:
      - {{ .EnvFile }}
{{- else if .Environment }}
    environment:
{{- range $key, $value := .Environment }}
      {{ $key }}: {{ $value }}
{{- end }}
{{- end }}
{{- end }}
`

// DockerService represents a single Docker service
type DockerService struct {
	Name        string            // "web", "db", "redis"
	Image       string            // "nginx:latest", "postgres:15"
	Ports       []string          // ["8080:80", "5432:5432"]
	Volumes     []string          // ["/data:/var/lib/postgresql/data"]
	Environment map[string]string // {"POSTGRES_PASSWORD": "secret"}
	EnvFile     string            // ".env.db" or ".env"
}

// DockerComposeConfig is the complete docker-compose configuration
type DockerComposeConfig struct {
	Version  string
	Services []DockerService
}

// HandleDocker manages all docker commands
func HandleDocker(command string) {
	switch command {
	case "generate":
		GenerateDockerCompose()
	case "help":
		PrintDockerHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

// PrintDockerHelp shows help for Docker
func PrintDockerHelp() {
	fmt.Println(`
📘 CONGA Docker
===============

Command: conga docker generate

Interactive flow to create docker-compose.yml with:
- Custom services
- Image selection
- Port mapping
- Volume mounting
- Environment variables
- Environment files

Generates a 'docker-compose.yml' file ready to use.`)
}

// AskNumberOfServices asks how many services to configure
func AskNumberOfServices() int {
	prompt := promptui.Prompt{
		Label:   "Number of services",
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

	result, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	var num int
	fmt.Sscanf(result, "%d", &num)
	return num
}

// AskService asks for details of a single Docker service
func AskService(serviceNumber int) DockerService {
	fmt.Printf("\n📦 Service %d:\n", serviceNumber)

	// Ask for service name
	namePrompt := promptui.Prompt{
		Label:   "Service name (e.g., web, db, cache)",
		Default: "service",
	}
	name, err := namePrompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return DockerService{}
	}

	// Ask for image
	imagePrompt := promptui.Prompt{
		Label:   "Docker image (e.g., nginx:latest, postgres:15)",
		Default: "alpine:latest",
	}
	image, err := imagePrompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return DockerService{}
	}

	service := DockerService{
		Name:        name,
		Image:       image,
		Ports:       []string{},
		Volumes:     []string{},
		Environment: make(map[string]string),
		EnvFile:     "",
	}

	// Ask for ports
	portsPrompt := promptui.Select{
		Label: "Add port mappings?",
		Items: []string{"Yes", "No"},
	}
	_, addPorts, err := portsPrompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return service
	}

	if addPorts == "Yes" {
		service.Ports = AskPorts()
	}

	// Ask for volumes
	volumesPrompt := promptui.Select{
		Label: "Add volume mounts?",
		Items: []string{"Yes", "No"},
	}
	_, addVolumes, err := volumesPrompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return service
	}

	if addVolumes == "Yes" {
		service.Volumes = AskVolumes()
	}

	// Ask for environment file
	envFilePrompt := promptui.Select{
		Label: "Use environment file?",
		Items: []string{"Yes", "No"},
	}
	_, useEnvFile, err := envFilePrompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return service
	}

	if useEnvFile == "Yes" {
		service.EnvFile = AskEnvFile()
	} else {
		// Ask for environment variables only if not using env_file
		envPrompt := promptui.Select{
			Label: "Add environment variables?",
			Items: []string{"Yes", "No"},
		}
		_, addEnv, err := envPrompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			return service
		}

		if addEnv == "Yes" {
			service.Environment = AskEnvironmentVariables()
		}
	}

	return service
}

// AskPorts asks for port mappings
func AskPorts() []string {
	ports := []string{}

	for {
		portPrompt := promptui.Prompt{
			Label: "Port mapping (container:host, e.g., 5432:5432)",
		}
		port, err := portPrompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		if port == "" {
			break
		}

		ports = append(ports, port)

		continuePrompt := promptui.Select{
			Label: "Add another port?",
			Items: []string{"Yes", "No"},
		}
		_, cont, err := continuePrompt.Run()
		if err != nil || cont == "No" {
			break
		}
	}

	return ports
}

// AskVolumes asks for volume mounts
func AskVolumes() []string {
	volumes := []string{}

	for {
		volumePrompt := promptui.Prompt{
			Label: "Volume mount (container:host, e.g., /data:/var/lib/postgresql/data)",
		}
		volume, err := volumePrompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		if volume == "" {
			break
		}

		volumes = append(volumes, volume)

		continuePrompt := promptui.Select{
			Label: "Add another volume?",
			Items: []string{"Yes", "No"},
		}
		_, cont, err := continuePrompt.Run()
		if err != nil || cont == "No" {
			break
		}
	}

	return volumes
}

// AskEnvironmentVariables asks for environment variables
func AskEnvironmentVariables() map[string]string {
	envVars := make(map[string]string)

	for {
		keyPrompt := promptui.Prompt{
			Label: "Environment variable name (e.g., POSTGRES_PASSWORD)",
		}
		key, err := keyPrompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		if key == "" {
			break
		}

		valuePrompt := promptui.Prompt{
			Label: fmt.Sprintf("Value for %s", key),
		}
		value, err := valuePrompt.Run()
		if err != nil {
			fmt.Println("Error:", err)
			break
		}

		envVars[key] = value

		continuePrompt := promptui.Select{
			Label: "Add another variable?",
			Items: []string{"Yes", "No"},
		}
		_, cont, err := continuePrompt.Run()
		if err != nil || cont == "No" {
			break
		}
	}

	return envVars
}

// AskEnvFile asks for environment file
func AskEnvFile() string {
	prompt := promptui.Prompt{
		Label:   "Environment file path (e.g., .env, .env.db)",
		Default: "",
	}

	envFile, err := prompt.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}

	return envFile
}
// WriteDockerCompose writes the docker-compose configuration to a file
func WriteDockerCompose(config DockerComposeConfig, filename string) error {
	// Parse the template
	tmpl, err := template.New("docker-compose").Parse(dockerComposeTemplate)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	// Create the file
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Execute template and write to file
	err = tmpl.Execute(file, config)
	if err != nil {
		return fmt.Errorf("error executing template: %v", err)
	}

	return nil
}
// GenerateDockerCompose manages the entire Docker Compose interactive flow
func GenerateDockerCompose() {
	fmt.Println("\n🐳 Configuring Docker Compose...")

	// Ask for number of services
	numServices := AskNumberOfServices()

	config := DockerComposeConfig{
		Version:  "3.8",
		Services: []DockerService{},
	}

	// Ask for each service
	for i := 1; i <= numServices; i++ {
		service := AskService(i)
		config.Services = append(config.Services, service)
	}

	// Display captured configuration
	fmt.Printf("\n✅ Configuration captured:\n")
	fmt.Printf("   Version: %s\n", config.Version)
	fmt.Printf("   Services: %d\n\n", len(config.Services))

	for i, service := range config.Services {
		fmt.Printf("   Service %d: %s\n", i+1, service.Name)
		fmt.Printf("      → Image: %s\n", service.Image)
		if len(service.Ports) > 0 {
			fmt.Printf("      → Ports: %v\n", service.Ports)
		}
		if len(service.Volumes) > 0 {
			fmt.Printf("      → Volumes: %v\n", service.Volumes)
		}
		if service.EnvFile != "" {
			fmt.Printf("      → Env file: %s\n", service.EnvFile)
		}
		if len(service.Environment) > 0 {
			fmt.Printf("      → Environment variables: %d\n", len(service.Environment))
		}
	}

	// Generate the file
	filename := askOutput("docker-compose.yml")
	err := WriteDockerCompose(config, filename)
	if err != nil {
		fmt.Printf("\n❌ Error generating config: %v\n", err)
		return
	}

	fmt.Printf("\n✅ Generated: %s\n", filename)
	fmt.Println("\n🎉 Ready to use!")
}