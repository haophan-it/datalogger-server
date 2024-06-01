# Datalogger - Backend

## Setup
- Install Go
- Install Taskfile
- Install Serverless
- Create .env file with
  IAM_ID=""

## Build Go applications

To build the Go applications, you run the following command:

```sh
task build-go
```

To build the only one function, you run the following command:
```sh
task build-function func="{function_name}"
```

## Deploy
To deploy the all function, you run the following commands:

```sh
serverless deploy
```

To deploy the only one function, you run the following command:
```sh
task deploy-function func="{function_name}"
```