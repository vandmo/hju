# hju

Taking *git* to the next level!

_hju_ is a tool to manage a set of repositories in a parent repository.
The biggest difference with git submodules is that submodules points to specific commits in the child repositories.
The managed repositories are in a file named _hju.json_ and the folders are ignored using .gitignore when using hju to add repositories.

# Installation

## Linux, Debian based
```bash
echo "deb [trusted=yes] https://apt.fury.io/bztk/ /" | sudo tee /etc/apt/sources.list.d/bztk.list > /dev/null
sudo apt update
sudo apt install hju
```

## Linux, Red Hat based
Add the following YUM Repository: https://yum.fury.io/bztk/
`sudo yum install hju`

## Other
Download the suitable release and unpack it.

# Commands

## add
Creates or updates hju.json and .gitignore
You do this in an existing repository:

Example:
```bash
hju add git@github.com:vandmo/dependency-lock-maven-plugin.git
hju add git@github.com:vandmo/google-java-format.git
git commit -a -m"Initial commit"
git push
```

This will have created a file named _hju.json_ which contains a list of the added repositories.
.gitignore will also have been updated to exclude the folders dependency-lock-maven-plugin and google-java-format.

## clone
Clones all child repositories.

## folders
Lists all managed folders. Can be used as pipe input.

## switch [-c|--create]
Switches to a branch in all managed repositories. Optionally creates it if missing.

## fix
Normalizes hju.json and makes sure all entries are in .gitignore

## fast-forward
Runs `git pull --ff-only` in all child repositories.

## remove
Removes a repository based on its folder name.

## status
Prints a summary of the status of all child repositories.
