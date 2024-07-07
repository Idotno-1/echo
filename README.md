# Echo

Simple Chat made in Go.

## Run

```sh
cd src/
go run .
```

## Build

```sh
cd src/
go build .
```

## Config

- `ECHO_ENABLE_FRONTEND`: Enable frontend (default to "true")
- `ECHO_POSTGRES_USER`: Database user (default to "postgres")
- `ECHO_POSTGRES_PASS`: Database password (default to "postgres")
- `ECHO_POSTGRES_NAME`: Database name (default to "echo")
- `ECHO_POSTGRES_HOST`: Database host (default to "localhost")
- `ECHO_POSTGRES_PORT`: Database port (default to "5432")

## What's next ?

- Rooms (create or connect to)
- Display user currently seeing the room
- Notifications on new message
- Enhanced UI with CSS
- Keep the last X messages for a room to display on join
- Rate limiting
- Docker Compose
- CI/CD
- Profile Pictures
- Mutualise common front features in templates
- Redirect to login when invalid token
- Chat commands
- Markdown rendering
