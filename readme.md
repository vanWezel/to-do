# To Do Service
This service will help you organize your tasks.

## ENV variables
- PORT=80 -> http port the server listens to

## How to run (single instance)?
```
docker build -t to-do .

docker run --rm to-do --name --p 80:80 to-do
```

## How to run (clustered)?
### Build the container
`docker build -t to-do .`

### Init docker swarm
`docker swarm init`

### Run
`docker stack deploy --compose-file docker-compose.yml to-do`

## How to stop?
````
docker stack rm to-do
docker swarm leave
````

## Endpoints
- GET /health
- GET / -> show all the tasks
- POST / -> create a new task {"task": "New task", "priority": 5}
- PUT /:id -> update a task {"task": "New task"} -> only the name is changed
- DELETE /:id -> delete a task
- PUT /:id/complete -> mark a task as complete
- POST /:task_id/comments -> add a comment to a task
- DELETE /:task_id/comments/:id -> delete a comment

## Dependencies
- go: 1.16
- docker

## Packages
### [Testify](https://github.com/stretchr/testify)
A good package for unit tests.

### [Echo](https://echo.labstack.com/guide/)
Well documented, lightweight, easy to use and unit test.

### [Validator](https://github.com/go-playground/validator)
Echo doesn’t have built-in data validation capabilities, so I use this third-party library.
