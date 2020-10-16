This is a fuzzy finder which can be used in a bash pipe command chain to achieve
anything you can imagine.

It is inspired by the awesome [fzf](https://github.com/junegunn/fzf) cli but uses
a GTK dialog box instead of the console.

## Usage

```
$ find * | fuzzygui | gvim
```

## Use case 1

Fuzzy finder for pass manager (regex removes extensions)

```
function fuzzypass {
  pushd ~/.password-store
  find * | fuzzygui | sed 's/\(.*\)\..*/\1/' | xargs pass -c1
  popd
}
```