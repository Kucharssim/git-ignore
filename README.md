# git-ignore

This is a simple extension for git which allows you to add statements to `.gitignore` files without having to 

1. Open the .gitignore file
2. Write to the .gitignore file
3. Close the file
4. Commit your .gitignore file

## Examples

### Adding a single folder to `.gitignore`

Adding a `bin/` folder to `.gitignore` is easy as:

```
git ignore bin/
```

Alternatively, using the `--commit` flag will also automatically commit the changes to the `.gitignore`:

```
git ignore bin/ --commit
```

### Adding multiple files or folders

```
git ignore bin/ secrets.txt --commit
```

This command adds `bin/` and `secrets.txt` to `.gitignore` and automatically commts the changes.

### Wildcards

Using wildcards is possible, but you need to be carefull whether you want to match existing files or whether you want to add the wildcard itself to the `.gitignore`:

```
git ignore "**.so"
```

would add `**.so` to `.gitignore`, which leads to ignoring any `.so` files in the future. On the other hand

```
touch hello.txt
touch world.txt
git ignore *.txt
```

would only add `hello.txt` and `world.txt` to `.gitignore`, thereby allowing to add other `.txt` files in the future.

### Local .gitignore

By default the ignored terms are added to the `.gitignore` file in the root of the git work tree. The flag `--here` allows you to add to `.gitignore` file in the current directory instead.

```
mkdir output
cd output
git ignore "**.html" --here
```

would write to a file `output/.gitignore`, thereby `.html` files only in the `output` directory are ignored.
