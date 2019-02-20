# TrendAI_mobile_backend
Mobile backend development works

## Setup application
Follow bellow steps to setup your application:

- Clone this repository to your machine.
- Run `cd /your/project/directory` to go inside project directory.
- Run `cp .env.example .env` to create new `.env` file from example file.
- Edit environment variables necessary in `.env` file.
- Run `docker-compose build` to build your application.

## Run application

To run application, please go inside project directory and run `docker-compose up`.
Your application located at [http://localhost:5000](http://localhost:5000). 

## APIs
### Login [POST /auth/login]
Use this API for authentication.

+ Request (application/json)

        {
            "access_token": "6212189571-wzBJbsjkqASHSDm0MWcyAytGcsaJiK1XsKoeYcR",
            "access_token_secret": "NeSZj2xvwEJgUTOAmqNPln5jALE5hv12xvXF2KRhXb12"
        }

+ Response 200 (application/json)

        {
            "user": {
                "id": "5c6d46f239da8e004593373c",
                "name": "Bui Thang",
                "email": "zenthangplus@gmail.com",
                "createdAt": "Wed, 20 Feb 2019 12:24:17 GMT"
            },
            "token": {
                "access_token": "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NTA2ODAyNDYsIm5iZiI6MTU1MDY4MDI0NiwianRpIjoiNGU0OGRhOTUtZDEyMS00MmViLWE4MWQtNGE4YmUzNzQ1NDNkIiwiZXhwIjoxNTUwNjgxMTQ2LCJpZGVudGl0eSI6IjVjNmQ0NmYyMzlkYThlMDA0NTkzMzczYyIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.qAuvOPCQVr3YAFuD6B4RmqLmEokv-05ttL8BzPRBAug"
            }
        }

+ Response 401 (application/json)

        {
            "msg": "Unauthorized"
        }

+ Response 522 (application/json)

        {
            "msg": "Missing access_token"
        }

### Get user [GET /user]

In order to authenticate a user, please add this header ``Authorization: {access_token}`` to your request.

+ Request (application/json)

    + Headers
        
            Authorization: eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NTA2ODAyNDYsIm5iZiI6MTU1MDY4MDI0NiwianRpIjoiNGU0OGRhOTUtZDEyMS00MmViLWE4MWQtNGE4YmUzNzQ1NDNkIiwiZXhwIjoxNTUwNjgxMTQ2LCJpZGVudGl0eSI6IjVjNmQ0NmYyMzlkYThlMDA0NTkzMzczYyIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.qAuvOPCQVr3YAFuD6B4RmqLmEokv-05ttL8BzPRBAug
            

+ Response 200 (application/json)

        {
            "id": "5c6d46f239da8e004593373c",
            "name": "Bui Thang",
            "email": "zenthangplus@gmail.com",
            "createdAt": "Wed, 20 Feb 2019 12:24:17 GMT"
        }

+ Response 401 (application/json)

        {
            "msg": "Unauthorized"
        }
