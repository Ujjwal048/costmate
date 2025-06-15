package aws

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Profile represents an AWS profile configuration
type Profile struct {
	Name string
	Role string
}

// GetAvailableProfiles reads the AWS credentials file and returns a list of available profiles
func GetAvailableProfiles() ([]Profile, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}

	credentialsPath := filepath.Join(homeDir, ".aws", "credentials")
	file, err := os.Open(credentialsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open AWS credentials file: %v", err)
	}
	defer file.Close()

	var profiles []Profile
	scanner := bufio.NewScanner(file)
	currentProfile := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentProfile = strings.Trim(line, "[]")
			profiles = append(profiles, Profile{Name: currentProfile})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading AWS credentials file: %v", err)
	}

	return profiles, nil
}

// SwitchProfile sets the AWS_PROFILE environment variable to the specified profile
func SwitchProfile(profileName string) error {
	if err := os.Setenv("AWS_PROFILE", profileName); err != nil {
		return fmt.Errorf("failed to set AWS_PROFILE environment variable: %v", err)
	}
	return nil
}

// GetCurrentProfile returns the currently active AWS profile
func GetCurrentProfile() string {
	return  "default"
}
