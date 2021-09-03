# hju

Taking *git* to the next level!

_hju_ is a tool to manage a set of repositories in a parent repository.
The biggest difference with git submodules is that submodules points to specific commits in the child repositories.
The managed repositories are in a file named _hju.json_ and the folders are ignored using .gitignore when using hju to add repositories.

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

## fix
Normalizes hju.json and makes sure all entries are in .gitignore

## fast-forward
Runs `git pull --ff-only` in all child repositories

## status
Prints a summary of the status of all child repositories

