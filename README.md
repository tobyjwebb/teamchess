# Toby's Team Chess

A pet project for the 1st period of 2022.

## Available commands

```shell
modd # Build and run the project, reloading if changes are detected
bazel run //src/cmd/web_frontend # Run the web frontend
```

## Env variables

The following environment variables can be used:

| Name                 | Description                                          | Default value |
|----------------------|------------------------------------------------------|---------------|
| TC_FRONTEND_ADDR     | The address for the frontend web server to listen on | :8081         |
| TC_REDIS_ADDR        | The address where the Redis server is listening      | :6379         |
