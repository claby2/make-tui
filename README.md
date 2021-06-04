# make-tui

![preview](./assets/preview.png)

`make-tui` is an application to display and run Makefile rules in the terminal.

## Installation

    $ go get github.com/claby2/make-tui

## Configuration

An optional configuration file is located at `${HOME}/.config/make-tui/config.yml`.

The following is a sample `config.yml` file:

```yaml
select_foreground: white
select_background: black
theme: vim
```

## Usage

    Usage of build/make-tui:
      -a, --all                Display all targets including special built-in targets
      -f, --file-name string   Parse given file as Makefile
      -h, --help               Print this message and exit
          --list-themes        List built-in syntax highlighting themes
