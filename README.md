# Getir Case Study


I use local config where under the .config folder.You can add your env configs but i suggest you to use the encryption tool like sops.
I want to write logs to file because this project does not contains any tools to store log. You can see logs/output.log
I deployed this app to heroku. Link => [https://getir-test-study-app.herokuapp.com/](https://getir-test-study-app.herokuapp.com/)

## Installation
Installation links used in this demo are given below:
* [Go](https://go.dev/doc/install)
* [Docker](https://docs.docker.com/get-docker)


## Running
```bash
#build & run it
 make build && ./getir-assignment
#Browse to http://localhost:3000
```
### With Docker
```bash
docker build -t <name> . -f Dockerfile   
```

## Available APIs
- `POST /record"`: The endpoint used to get records by filter.

#### Example Request
**Note** 

```bash
curl --location --request POST 'http://localhost:3000/record' \
--header 'Content-Type: application/json' \
--data-raw '{
    "startDate":"2016-12-27",
    "endDate":"2018-02-02",
    "minCount":2700,
    "maxCount":3000
}'
```
- `GET /in-memory?key=<key>`: The endpoint used to get key value pair.
#### Example Request
```bash
curl --location --request GET 'http://localhost:3000/in-memory?key=test'
```
- `POST /in-memory`: The endpoint used to add key value pair.
#### Example Request

```bash
curl --location --request POST 'http://localhost:3000/in-memory' \
--header 'Content-Type: application/json' \
--data-raw '{
    "key": "emre2",
    "value": "2"
}'
```
## Test
```bash
make unit-test
```

## Db Test
```bash
make unit-test
```