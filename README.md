# Switchterm

![GitHub repo file count](https://img.shields.io/github/directory-file-count/golangast/switchterm) 
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/golangast/switchterm)
![GitHub repo size](https://img.shields.io/github/repo-size/golangast/switchterm)
![GitHub](https://img.shields.io/github/license/golangast/switchterm)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/golangast/switchterm)
![Go 100%](https://img.shields.io/badge/Go-100%25-blue)
![status beta](https://img.shields.io/badge/Status-Beta-red)


  - [switchterm](#switchterm)
  - [General info](#general-info)
  - [Why build this?](#why-build-this)
  - [What does it do?](#what-does-it-do)
  - [Technologies](#technologies)
  - [Requirements](#requirements)
  - [Repository overview](#repository-overview)
  - [Things to remember](#things-to-remember)
  - [Reference Commands](#reference-commands)
  - [Special thanks](#special-thanks)



## General info
This project is a command line selection tool.
It also generates the config for storing the commands


## Why build this?
* Go never changes
* Wanted a easy quick way to run commands


## What does it do?
* lets you select multiple commands and run them



## Technologies
Project is created with:
* [atomicgo.dev/keyboard](https://atomicgo.dev/keyboard) - For pressing keys
* [spf13/viper](https://github.com/spf13/viper) - Basically for keeping state


## Requirements
* go 1.21 for gonew

## How to run as is?
*1. install the library
```bash
go get github.com/golangast/switchterm
```
*2. do mod tidy/vendor
```bash
go mod tidy && go mod vendor
```
*3. run the program
```bash
go run .
```


*1. install gonew to pull down project quickly
```bash
go install golang.org/x/tools/cmd/gonew@latest
```
*2. run gonew
```bash
gonew github.com/golangast/switchterm example.com/switchterm
```

## Repository overview
```bash
├── switchtermer [folder for functions]
├── configure [folder for config utility functions]
├── switchtermer [folder for functions]
│   ├── colortermer [file for coloring selections]
│   ├── switchutility [file for selection functions]
│   └── switchtermer.go [where the service functions are]

```

## Things to remember
* using atomicgo.dev/keyboard there is no way to call itself after a key press

## Reference Commands
* "enter" is to select
* "e" is to select many
* "r" is to remove
* "x" is to execute
* "q" is to quit



## Special thanks
* [Go Team because they are gods](https://github.com/golang/go/graphs/contributors)
* [Creators of atomicgo/keyboard - Marvin Wendt](https://github.com/MarvinJWendt)
