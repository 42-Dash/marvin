# Dash - Marvin

## Overview

**Dash** is a Go-based framework designed to facilitate the launch and management of programming competitions. It automates repository creation, manages participants, handles submissions, and provides tools for analysis and visualization of results.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
  - [.env File](#env-file)
  - [participants.json](#participantsjson)
  - [maps.json](#mapsjson)
- [Usage](#usage)
  - [CLI Menu](#cli-menu)
- [Running the Visualizer](#running-the-visualizer)
- [Project Structure](#project-structure)
- [Go Version & Dependencies](#go-version--dependencies)
- [Taskfile](#taskfile)

## Features

- **Automated Repository Management:** Create GitHub repositories from a template.
- **Participant Handling:** Manage teams and their associated competitions.
- **Submission Analysis:** Clone and evaluate participant submissions.
- **Trace Uploading:** Push test traces to specific branches for review.
- **Result Generation:** Compile logs into comprehensive JSON results for visualization.
- **Docker Integration:** Safe execution of user programs in isolated environments.
- **Interactive CLI:** User-friendly menu for managing various tasks.
- **Visualizer:** JavaScript-based tool for displaying competition results.

## Prerequisites

- **Go:** Version 1.22
- **Docker:** For running the visualizer and executing user programs securely.
- **GitHub Access Token:** Must have permissions to create repositories and manage access.
- **Organization:** The GitHub organization where repositories will be created must already exist and contain a `template-marvin` repository.

## Installation

1. **Clone the Repository:**
    ```bash
    git clone https://github.com/your_org/dashinette.git
    cd dashinette
    ```

2. **Install Dependencies:**
    ```bash
    go mod download
    ```

## Configuration

### `.env` File

Create a `.env` file in the root directory based on the `.env.example`:

```env
ORGANIZATION_NAME=your_github_org
GITHUB_ACCESS_TOKEN=your_access_token
```

- **ORGANIZATION_NAME:** Name of the GitHub organization where repositories will be created.
- **GITHUB_ACCESS_TOKEN:** Token with permissions to create repositories and manage collaborators.

#### `participants.json` configuration file
Defines the teams participating in the competition.

```json
{
    "teams": [
        {
            "name": "The-Avengers",
            "members": ["IronMan", "Thor777"],
			"league": "open"
        },
        {
            "name": "Pirates-of-the-Caribbean",
            "members": ["jack-sparrow", "wturner"],
			"league": "open"
        },
        {
            "name": "The-Justice-League",
            "members": ["BatmanIsHere", "SuPeRmAn"],
			"league": "rookie"
        }
    ]
}
```
- **name:** Team name.
- **members:** List of team members.
- **league:** Competition league (`open` or `rookie`).

#### `maps.json` configuration file
Specifies the maps used for generating traces for tests and final submittions.

```json
{
    "rookieleague": [
        {
            "path": "dashes/marvin/maps/rookieleague/amongus.txt",
            "name": "Impostor",
            "timeout": 2
        },
        ...
    ],
    "openleague": [
        {
            "path": "dashes/marvin/maps/openleague/sipmson.txt",
            "name": "You know him",
            "timeout": 3
        },
        ...
    ]
}
```

- **path:** File path to the map.
- **name:** Map name.
- **timeout:** Execution timeout in seconds.

## Usage
### CLI Menu
Launch the CLI with:
```bash
go run main.go
```

#### Menu Options:

1. **Create GitHub Repositories by template:** Initializes repositories for each team based on the `template-marvin` repo.
2. **Update README files with Subjects:** Pushes updated subjects to team repositories.
3. **Grant Collaborator Write Access:** Adds collaborators with write permissions to repositories.
4. **Configure Repositories as Read-Only:** Sets repositories to read-only access.
5. **Clone and Analyze Submissions to Generate Traces:** Clones student submissions and generates trace data.
6. **Parse and Upload Traces to 'traces' Branch:** Uploads generated traces to the `traces` branch of each repository.
7. **Parse Logs and Generate results.json:** Compiles logs into a `results.json` file for visualization.
8. **Exit:** Closes the CLI.

### Example Workflow

1. Create Repositories:
    - Select "Create GitHub Repositories by template."
2. Add collaborators:
    - Select "Grant Collaborator Write Access."
3. Analyze Submissions and give feedback
    - Select "Clone and Analyze Submissions to Generate Traces."
    - Select "Parse and Upload Traces to 'traces' Branch."
4. Analyze Submissions and prepare results for visualization:
    - Select "Clone and Analyze Submissions to Generate Traces."
    - Select "Parse Logs and Generate results.json."

All actions are logged in `app.log` for detailed tracking.

### Running the Visualizer
The visualizer is a Dockerized JavaScript application that displays competition results.

1. **Launch the Visualizer:**
```bash
go run scripts/marvin/visualiser/start/main.go generated_results.json
```

2. **Access the Visualizer:** Open your browser and navigate to `http://localhost:8080`.
```bash
.
├── README.md
├── Dockerfile
├── Taskfile.yml
├── app.log
├── cmd
│   └── marvin
│       ├── tester
│       ├── tests
│       └── marvin
├── config
│   ├── .env
│   ├── maps.json
│   └── participants.json
├── dashes
│   └── marvin
│       ├── README.md
│       ├── maps
│       ├── repos
│       ├── solutions
│       ├── traces
│       └── visualiser
│           ├── Dockerfile
│           ├── images
│           ├── index.html
│           ├── libraries
│           ├── results.json
│           └── src
├── go.mod
├── go.sum
├── internals
│   ├── cli
│   ├── grader
│   └── traces
├── main.go
├── pkg
│   ├── constants
│   ├── containerization
│   ├── github
│   ├── logger
│   └── parser
└── scripts
```

## Go Version & Dependencies
- Go Version: 1.22

**Dependencies:**
```go
module dashinette

go 1.22

require (
    github.com/docker/docker v27.3.1+incompatible
    github.com/joho/godotenv v1.5.1
    github.com/manifoldco/promptui v0.9.0
)
```

## Taskfile

The project utilizes a [`Taskfile.yml`](Taskfile.yml) to automate various development tasks. [Task](https://taskfile.dev/) is a simple task runner that helps streamline your workflow.

### Prerequisites

- **Task:** Ensure Task is installed on your system. Follow the [Task installation guide](https://taskfile.dev/installation/) to set it up.

### Available Tasks

- **marvin:** Build the marvin project.
- **marvin-tester:** Build the marvin project tester (for development).
- **marvin-docker:** Rebuild the Docker image for the marvin project.
- **clean:** Clean up the repository by removing build artifacts and logs.
- **delete-repos:** Remove the organization repositories. **Use with caution.**
- **test:** Run all tests and clean up afterwards.

### Usage

To execute a task, use the following command:

```bash
task <task-name>
```

### Examples:

- **Build the marvin project:**

```bash
task marvin
```

### Viewing Help
To view available commands and their descriptions:

```bash
task help
```

This will display all available tasks with their respective descriptions.
