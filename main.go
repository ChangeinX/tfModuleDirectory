package main

import (
	"fmt"
	"tfModuleDirectory/pkg/gitignore"
	"tfModuleDirectory/pkg/terraform"
)

func main() {
	if err := terraform.CreateProjectStructure(); err != nil {
		fmt.Println(err)
		return
	}

	if err := gitignore.CreateGitIgnore(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Terraform project structure created successfully.")
}
