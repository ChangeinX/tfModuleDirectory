# Terraform Project Structure Generator

This tool helps in setting up a new Terraform project by creating a standard directory structure and necessary files, including a `.gitignore` tailored for Terraform projects.

## Getting Started

To get started with this tool, clone the repository to your local machine or download the source code.

### Prerequisites

Ensure you have Go installed on your system. You can check this by running:

```sh
go version
```
If Go is not installed, follow the installation instructions for your platform at Go's official site.

## Installation
Clone the repository:
```sh
git clone https://github.com/ChangeinX/tfModuleDirectory.git
cd tfModuleDirectory
```
Build the tool:
```sh
go build -o tfmodgen
```

## Add to System Path (Optional)

Optionally, you can add the tool to your system path to call it from the cli. 