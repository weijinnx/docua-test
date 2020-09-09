# Doc.ua Test Task

Your task is to write a simple application server that prints a message at a given time in the future.

## Description

The server has only 1 API:

```
[GET] /printMeAt?time=2020-09-05T15:16:15Z&message=message1
Time should be in UTC
```

It receives two parameters, time and message, and writes that message to the server console at the given time.

Since we want the server to be able to withstand restarts it will use Redis to persist the messages and the time they should be sent at.

You should also assume that there might be more than one server running behind a load balancer (load balancing implementation itself does not need to be provided as part of the answer).

In case the server was down when a message should have been printed, it should print it out when going back online.

## The focus of the exercise is

- the efficient use of Redis and its data types
- messages should not be lost
- the same message should be printed only once
- message order should not be changed
- should be scalable
- seeing your code in action (SOLID would be a plus)

**Use only Redis**

## Requirements

- Docker

## Get started

Copy `.env.local` as `.env`:

```bash
cp .env.local .env
```

So, to start project you need to clone it and just run:

```bash
make run
```

If you want to rebuild it, save your changes and run:

```bash
make build
```

After starting server it will be available on `localhost:8080`

