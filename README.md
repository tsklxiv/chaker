<h1 align="center" class="header">Chaker</h1>
<p align="center" class="desc">
  The Hacker News 'client' in the terminal. Written in Golang.
</p>

## Table of Content
- [Introduction](#introduction)
- [Features](#features)
- [Contributions and Issues](#contributions-and-issues)
- [Installation](#installation)
- [Usage](#usage)
- [Note](#note)
- [Credits](#credits)

## Introduction
[![CodeFactor](https://www.codefactor.io/repository/github/hoangtuan110/chaker/badge)](https://www.codefactor.io/repository/github/hoangtuan110/chaker)

Chaker (formerly Hecker) is a Hacker News 'client' for the terminal written in Golang.

(The client is in quote because technially this is more of a web scraper with a UI rather than an actual client)

## Features

* Easy to use - You just need to learn a few keybinds to use Chaker.
* Move between submissions easily with Up and Down arrow key.
* Open the submission's URL using `Enter`.
* View the comment section (or more correctly the submission itself) using `c`.
* Move between different pages of submissions.
* Shows up submission's data (like upvote and so on) only when they are being pointed and fainted.
* Shows up the time on the top, and the page number and help at the bottom!

## Contributions and Issues

Pull requrests and contributions are welcome. If you found any bugs or want to requrest a feature, please open an issue.

## Installation

* First, install the Go compiler at [here](https://golang.org/dl/) or if you are using a Linux distro, install it using your distro's package manager.
* Next, if you have `$GOPATH` set up:
```sh
go get github.com/HoangTuan110/chaker
```
  And if you don't:
```sh
git clone https://github.com/HoangTuan110/chaker.git
cd chaker
go build
```
* Finally, you will see the `chaker` binary in the cloned directory.

## Usage

Just run:

```sh
./chaker
```

## Note
* When start the program, you will need to wait for a few seconds for the program to scrape data.

## Credits

Thanks [Charm](https://charm.sh) for their amazing CLI library `bubbletea` and `lipgloss`.

Thanks [README Templates](https://www.readme-templates.com) and [GitPoint's README](https://github.com/gitpoint/git-point#readme) for the template. This project uses GitPoint's README template.

Thanks [this TOC generator](https://ecotrust-canada.github.io/markdown-toc/) for the TOC (Table Of Content).
