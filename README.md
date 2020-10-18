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

## Use case 2

Implement a contacts directory with search and flat files.

We assume a directory structure where each file is named after a contact and each
line in the file is a phone number or some other contact detail:

```
~/Dropbox/phones $ cat Ada\ Lovelace
email: ada@example.com
phone: +30 1234567890
```

A script like the following will let you edit the contact using fuzzy search:

find.sh:

```bash
find * ! -name 'find.sh' -exec awk '{print FILENAME " - " $0}' {} \;  | fuzzygui |  sed -r 's/(.*) -.*/\1/' | xargs -d"\n" gvim
```
