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
Your application located at [http://localhost:5000](http://localhost:8080). 

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
                "id": "5c7285d05d40e0d4bfb7ede5",
                "name": "Bui Thang",
                "email": "zenthangplus@gmail.com",
                "created_at": "2019-02-24T11:53:52.155Z"
            },
            "token": {
                "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NTE2MzE0MjUsInVpZCI6IjVjNzI4NWQwNWQ0MGUwZDRiZmI3ZWRlNSJ9.n3ALesfOg7gzBiYjETdsPrLP7K9fwaNja8teFAkXqVg",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJyZWYiOiJyZWZyZXNoIiwidWlkIjoiNWM3Mjg1ZDA1ZDQwZTBkNGJmYjdlZGU1In0.EPbAA5QF-dy8nEItbUshzcnRd368WLGw7-d1lIzQ3j8"
            }
        }

+ Response 401 (application/json)

        {
            "code": "unauthorized",
            "message": "Unauthorized"
        }

+ Response 400 (application/json)

        {
            "code": "unauthorized",
            "message": "Couldn't parse your request"
        }

### Get user [GET /user]

In order to authenticate a user, please add this header ``Authorization: Bearer {access_token}`` to your request.

+ Request (application/json)

    + Headers
        
            Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NTA2ODAyNDYsIm5iZiI6MTU1MDY4MDI0NiwianRpIjoiNGU0OGRhOTUtZDEyMS00MmViLWE4MWQtNGE4YmUzNzQ1NDNkIiwiZXhwIjoxNTUwNjgxMTQ2LCJpZGVudGl0eSI6IjVjNmQ0NmYyMzlkYThlMDA0NTkzMzczYyIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.qAuvOPCQVr3YAFuD6B4RmqLmEokv-05ttL8BzPRBAug
            

+ Response 200 (application/json)

        {
            "id": "5c7285d05d40e0d4bfb7ede5",
            "name": "Bui Xuan Thang",
            "email": "zenthangplus@gmail.com",
            "created_at": "2019-02-24T11:53:52.155Z"
        }

+ Response 401 (application/json)

        {
            "code": "unauthorized",
            "message": "Unauthorized"
        }
