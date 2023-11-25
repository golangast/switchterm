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
  - [Special thanks](#special-thanks)



## General info
This project is a command line selection tool


## Why build this?
* Go never changes
* Wanted a easy quick way to run commands


## What does it do?
* lets you select multiple commands



## Technologies
Project is created with:
* [atomicgo.dev/keyboard](https://atomicgo.dev/keyboard) - For pressing keys


## Requirements
* go 1.21 for gonew



## Repository overview
```bash
├── switchtermer [folder for functions]
│   ├── colortermer [file for coloring selections]
│   ├── switchutility [file for selection functions]
│   └── switchtermer.go [where the service functions are]

```

## Things to remember
* using atomicgo.dev/keyboard there is no way to call itself after a key press



## Special thanks
* [Go Team because they are gods](https://github.com/golang/go/graphs/contributors)
* [Creators of atomicgo/keyboard - Marvin Wendt](https://github.com/MarvinJWendt)
