# 🌟 Lux Lucet 
<img src="https://europe-west1-captains-badges.cloudfunctions.net/function-clone-badge-pc?project=team142/lux-lucet" />&nbsp;
<img src="https://travis-ci.org/team142/lux-lucet.svg?branch=master" />&nbsp;
<a href="https://goreportcard.com/report/github.com/team142/lux-lucet"><img src="https://goreportcard.com/badge/github.com/team142/lux-lucet" /></a>&nbsp; 
<a href="https://codeclimate.com/github/team142/lux-lucet/maintainability"><img src="https://api.codeclimate.com/v1/badges/ee3e04d0fac7419ccae9/maintainability" /></a>&nbsp; 
[![License](http://img.shields.io/:license-mit-blue.svg?style=flat)](http://badges.mit-license.org)

System health server written in Go for systems composed of and dependant on subsystems, requiring concurrently updating each subsystems health status. 

## Features
- Overall system health
- n subsystems
- Thread-safe updating of state
- Thread-safe reading of state
- Http server - listens on address and return state in json

## Examples
See [example.go](/example/example.go) and [example docs](/example/readme.md)
