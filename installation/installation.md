# Introduction

## Source Code

http://git.srin.com/spc/gage/golang-academy

## Installation

### Step 1 - Download Go Latest Stable Version

```
$ wget -c https://golang.org/dl/go1.16.5.linux-amd64.tar.gz
```

### Step 2 - Extract to the Installation Location

```
$ sudo tar -C /usr/local -xvzf go1.16.5.linux-amd64.tar.gz
```

Where, **-C** specifies the destination directory..

### Step 3 - Configure System Environment for Go

Add `/usr/local/go/bin` to the PATH environment variable

```
$ echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
$ source ~/.bashrc
```

### Step 4 - Verify the Installation

```
$ go version
```