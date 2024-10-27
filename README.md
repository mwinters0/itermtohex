# itermtohex
Convert iTerm colors to hex. (In a single binary!)

## Install
```shell
go install github.com/mwinters0/itermtohex@latest
```

Or see the [releases](https://github.com/mwinters0/itermtohex/releases) page.

## Usage
```shell
itermtohex my-theme.itermcolors # convert to json and output on stdout
itermtohex print my-theme.itermcolors # print the converted color palette (assumes RGB support)
```
