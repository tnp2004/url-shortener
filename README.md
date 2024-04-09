# Url Shortener
Simple url shortener app powered by golang, mongodb, and fiber. It's easy to generate concise links for sharing purposes.

## Using
```
go run . <PATH_TO_ENV>
```

## API Path
> `GET` Health check
```bash
http://localhost:3000/v1/healthcheck
```
> `GET` Search short url id

*replace `:id` with short id* 
```bash
http://localhost:3000/v1/:id
```
> `POST` Convert Url
```bash
http://localhost:3000/v1/convert
```
*JSON body request Example*
```json
{ 
    "url": "https://www.youtube.com/watch?v=xvFZjo5PgG0"
}
```

## Example Environment
```bash
APP_STAGE=example
VERSION=v1
APP_NAME=urlshortener
APP_URL=localhost:3000
DB_URL=mongodb://<MONGO_DB_USERNAME>:<MONGO_DB_PASSWORD>@0.0.0.0:27017
```