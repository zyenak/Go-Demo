# Go Block Chain Project

This project is a simple implementation of a blockchain in Go. It allows you to manage blocks, create new blocks, and modify existing blocks.

## Getting Started

Follow these instructions to set up and run the project on your local machine.

### Prerequisites

Make sure you have Go installed on your system. You can download and install it from [here](https://golang.org/dl/).

### Installation

To install the project, clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/blockchain.git
```

## Usage of Git Commands

### Stage & Snapshot

- `git status`: Check the status of the working directory.
- `git add [file]`: Stage changes for the next commit.
- `git commit -m "[message]": Commit staged changes with a descriptive message.

### Setup & Init

- `git config --global user.name "[firstname lastname]": Set your name for version history.
- `git config --global user.email "[valid-email]": Set your email for version history.

### Branch & Merge

- `git branch [branch-name]`: Create a new branch for development.
- `git checkout [branch]`: Switch to a different branch.
- `git merge [branch]`: Merge changes from another branch into the current branch.

### Share & Update

- `git remote add [alias] [url]`: Add a remote repository as an alias.
- `git fetch [alias]`: Fetch changes from a remote repository.
- `git push [alias] [branch]`: Push local changes to a remote repository.

### Tracking Path Changes

- `git rm [file]`: Remove a file from the project and stage the deletion.
- `git mv [existing-path] [new-path]`: Move or rename a file and stage the change.

### Temporary Commits

- `git stash`: Temporarily store modified files to work on something else.
- `git stash pop`: Retrieve stashed changes to continue working.

### Rewrite History

- `git rebase [branch]`: Reapply commits on top of another branch.
- `git reset --hard [commit]`: Reset the repository to a specific commit.

### Inspect & Compare

- `git log`: View commit history.
- `git diff`: View changes between commits or branches.
