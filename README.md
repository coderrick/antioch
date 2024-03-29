# Antioch
![alt text](antioch1.png)

## Overview

Antioch is a progressive web application that uses Microservices architecture along with the InterSystems IRIS dataplatform to make both the code and the system infrastructure easy to scale.


## Installation
`git clone [this repo]`

BLOCKCHAIN SETUP:\
`cd blockchain/; ./setup.sh`\
`testrpc`\
`run truffle commands`\

## Installation IRIS Container
`git clone https://github.com/intersystems-community/objectscript-rest-docker-template.git:latest`

'cd' to the project folder

DOCKER SETUP:
`docker build --rm -f "Docker" -t iris-docker-disrupt:`

*Wait for build to complete sucessfully*

`docker run ...`

DATABASE SETUP:\
Must have cockroachdb installed!
`cd to db`\
`cockroach start --insecure` *must be done in separate terminal\
`cockroach sql --insecure --database=drill < schema.sql` *Initializing the db's and tables\

# Try it!

Take a look at our Beta web app [HERE!](http://www..com/)

[![Build Status](https://travis-ci.org/coderrick/drill.svg?branch=master)](https://travis-ci.org/coderrick/drill)
