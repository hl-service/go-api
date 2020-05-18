### Information
This repository contains web application (api).

### Requirements
- [Go](https://golang.org/) 1.14+
- [Docker](https://www.docker.com/) (if you want to run application in container)

### Installation
1. Install Go dependencies: ```go mod download```
2. Build app: ```go build -o app```
3. Run built application: ```./app``` on Linux/MacOS or execute ```app.exe``` on Windows.

### Run application with docker
1. Build an image: ```docker build -t hl-api .```
2. Run: ```docker run -d hl-api```

Application will available on: http://0.0.0.0:8080

