package main

import (
	"fmt"
	"tfModuleDirectory/pkg/file_contents/gitignore"
	"tfModuleDirectory/pkg/terraform/terraformstructure"
)

func main() {
	if err := terraformstructure.CreateProjectStructure(); err != nil {
		fmt.Println(err)
		return
	}

	if err := gitignore.CreateGitIgnore(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Terraform project structure created successfully.")
}
