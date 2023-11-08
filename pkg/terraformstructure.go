package terraformstructure

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateProjectStructure sets up the base directories and Terraform files
func CreateProjectStructure() error {
	baseDirs := []string{
		"modules/aws",
		"tf_state",
	}
	files := []string{
		"README.md",
		"main.tf",
		"variables.tf",
		"locals.tf",
		"outputs.tf",
		"providers.tf",
		"terraform.tfvars",
	}

	// Create base directories and files in them
	for _, dir := range baseDirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %w", dir, err)
		}
		for _, file := range files {
			filePath := filepath.Join(dir, file)
			if err := createFile(filePath); err != nil {
				return fmt.Errorf("error creating file %s: %w", filePath, err)
			}
		}
	}

	return nil
}

// createFile creates an empty file at the specified path
func createFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	fmt.Printf("Created %s\n", filePath)
	return nil
}
