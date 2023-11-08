package gitignore

import "os"

// CreateGitIgnore creates a .gitignore file with predefined contents
func CreateGitIgnore() error {
	const gitignoreContents = `# Local .terraform directories
	**/.terraform/*
	
	# .tfstate files
	*.tfstate
	*.tfstate.*
	
	# Crash log files
	crash.log
	
	# Exclude all .tfvars files, which might contain sensitive data, such as
	# password, private keys, and other secrets. These should not be part of
	# version control as they are data points which are potentially sensitive
	# and subject to change depending on the environment.
	*.tfvars
	
	# Ignore override files as they are usually used to override resources locally and so
	# are not checked in
	_override.tf
	_override.tf.json
	*.override.tf
	*.override.tf.json
	
	# Ignore CLI configuration files
	.terraformrc
	terraform.rc
	`
	return os.WriteFile(".gitignore", []byte(gitignoreContents), 0644)
}