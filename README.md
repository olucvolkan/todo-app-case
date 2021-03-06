# Todos Restful Api


Todo-app is backend and frontend project.I developed it using backend golang and frontend it using reactjs.

You can reach ui github repo [link.](https://github.com/olucvolkan/todoApp-ui)


# Demo-Preview

![Random GIF](./images/swagger-ui.png)

# Table of contents

- [Todos Restful Api](#todos-restful-api)
- [Demo-Preview](#demo-preview)
- [Table of contents](#table-of-contents)
- [Installation](#installation)
- [Usage](#usage)
    - [`build`](#build)
    - [`unit test`](#unit-test)
    - [`Docker for local build`](#docker-for-local-build)
    - [`Deployment`](#deployment)

# Installation
[(Back to top)](#table-of-contents)

To use this project, first clone the repo on your device using the command below:

```git init```

```git https://github.com/olucvolkan/todo-app-case```


# Usage
[(Back to top)](#table-of-contents)

### `build`

```sh
$ cd todo-app-case
$ go build main.go
$ go run main.go
```
Runs the app in the development mode.<br />
You must edit the .env file according to your own information!
Open [http://localhost:8080/swagger/index.html](http://localhost:3000) to view it in the browser.

The page will reload if you make edits.<br />
You will also see any lint errors in the console.

### `unit test`

```sh
$ cd test && go test
```

### `Docker for local build`

```sh 
$ docker build --tag full_todo_app .
$ docker-compose up --build
```


### `Deployment`

Deployment has been done to [heroku](https://aqueous-citadel-50556.herokuapp.com/swagger/index.html).

