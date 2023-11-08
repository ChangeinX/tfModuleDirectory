package main

import (
	"fmt"
	"github.com/ChangeinX/tfModuleDirectory/pkg/file_contents/gitignore"
	"github.com/ChangeinX/tfModuleDirectory/pkg/terraform/terraformstructure"
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
