# oh-my-update

>⚠️: This project is under active development

- [Description](#description)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)

## Description

oh-my-update is an [Oh My Zsh](https://ohmyz.sh/) plugin and theme updater written in Go

## Requirements
- [Go](https://go.dev/)
- [Oh My Zsh](https://ohmyz.sh/)

## Installation

### Arch Linux
[omu](https://aur.archlinux.org/packages/omu) is available as a package in the AUR. You can install it using your preferred AUR helper (e.g. `paru`):
```bash
$ paru -S omu
```

### Compiling from source
```bash
# Clone the repo
$ git clone git@github.com:utox39/oh-my-update.git

# cd to the repo
$ cd path/to/oh-my-update

# Build oh-my-update
$ go build omu.go

# Then move it somewhere in your $PATH. Here is an example:
$ mv omu ~/bin
```

## Usage

### Update plugins and themes
```bash
$ omu
```

## Contributing
If you would like to contribute to this project just create a pull request which I will try to review as soon as
possible.
