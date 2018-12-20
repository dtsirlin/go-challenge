# Go Challenge
A web application in Golang that exposes a REST endpoint that validates user credentials and returns JWT Token.

Author: Daniel Tsirlin

## Setup
How to set up the environment to run the application.

### Pre-requisites
- Docker must be installed on the machine.
- There must be a working internet connection on the machine this is run.

## Running
How to build, run and stop the application.

### Build Application
- Execute ```$ make build``` in the root directory of the repo

### Run Application
- Once the application has been built as a docker container, execute ```$ make run``` in the root directory of the repo
- You should now be able to hit the routes, for instance in your browser or in Postman

### Stop the Application
- Once running, either exit out of the run execution or execute ```$ make stop``` in another shell while in the root directory of the repo

## Assumptions
- Username and password credential pairs created within server command
- Username and passwords are case sensitive
- To load protected endpoint, a valid JWT token is one that is a valid jwt token which is a non-expired access token associated to valid users

---

###### README Last updated: 21/12/2018