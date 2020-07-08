# MediChat-UI

MediChat-Users is the Backend component of a Chat Provider application.

The application is written in **go** and support for these resources:
- User
- UserPref
- UserAuth
- Channel

Key Features:
- Ability to **sign in**/**sign out** with JWT and password Authentication

## How to install

- Clone the git repo

### with DOCKER

- Update or leave unchanged these files
- .docker/.env
- app.ini
- Dockerfile
- docker-compose.yaml

Once Happy, run this command

```
docker-compose up --build
```

I also like these commands:

```
docker-compose down --remove-orphans --volumes && docker-compose up --build
docker-compose down --remove-orphans --volumes && docker-compose build —no-cache &&  docker-compose up
```

### Locally

- Copy _app.ini.default_ to **app.ini** and make changes as needed.

#### Execute for development:

```
go run main.go
```

## Known BUGS

- Websocket is allowing channel to broadcast to clients.
- Swagger is not loading spec as expected.
