# LMWN Covid Summary API

## Prerequisite

- Install Go

## How to run

- Run server
  ```
  go run .
  ```
- Request API to get covid summary data at path GET /covid/summary. For example cURL, Run this command in terminal
  ```
  curl http://localhost:8080/covid/summary
  ```

## How to test

```
cd server
go test
```
