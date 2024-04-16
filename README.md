English | [简体中文](README-CN.md)

fastGit is a tool that helps you quickly submit code with a command line interface. It supports Linux, Mac, and Windows. The inspiration comes from [gum](https://github.com/charmbracelet/gum)

> This project is utilizing its own features to submit code.

![](assets/fast-git.gif)

### How to use

#### 1. Install Git

> Project dependencies on Git, please install Git first

```bash
# Debian/Ubuntu
sudo apt install git
# macOS
brew install git
```

#### 2. Install fastGit

```bash
# Linux/macOS
curl -sSL https://github.com/KevinYouu/fastGit/install.sh | bash
```

or

```bash
wget -qO- https://github.com/KevinYouu/fastGit/install.sh | bash
```

#### 3. Run

```bash
fastGit pa
```

### Features

- [x] `fastGit pa`, submit all changes in the working directory

- [x] `fastGit ps`, submit some changes in the working directory

- [x] `fastGit c`, clone a repository

- [x] `fastGit t`, create and push a tag

- [x] `fastGit s`, check the status of the repository

- [x] `fastGit rv`, get all remote repositories

- [x] `fastGit ra`, add a remote repository

More features will be added soon......

### Thanks to the following open source projects

[go](https://github.com/golang/go)

[bubbletea](github.com/charmbracelet/bubbletea)

[huh](github.com/charmbracelet/huh)
