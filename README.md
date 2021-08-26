# hju

Taking *git* to the next level!

You do this in an existing repository:
```bash
hju add git@github.com:vandmo/dependency-lock-maven-plugin.git
hju add git@github.com:vandmo/google-java-format.git
git commit -a -m"Initial commit"
git push
```

This will have created a file named _hju.json_ which contains a list of the added repositories.
.gitignore will also have been updated to exclude the folders dependency-lock-maven-plugin and google-java-format.

Someone else then clones the same repository and does:
`hju clone`

Run `hju status` to get a summary of the contained repositories.

Run `hju fetch` to run `git fetch` for each contained repository.
