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
Your application located at [http://localhost:8080](http://localhost:8080). 

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
            "id": "fVwYk3hcSkpNv7brodYq",
            "name": "Bui Xuan Thang",
            "email": "zenthangplus@gmail.com",
            "gender": "male",
            "dob": "1994-05-11T00:00:00Z",
            "education": "",
            "created_at": "2019-02-25T15:51:42.060512Z"
        }

+ Response 401 (application/json)

        {
            "code": "unauthorized",
            "message": "Unauthorized"
        }


### Update user [PATCH /user]

Update information for current user.

+ Request (application/json)

    + Headers
        
            Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NTA2ODAyNDYsIm5iZiI6MTU1MDY4MDI0NiwianRpIjoiNGU0OGRhOTUtZDEyMS00MmViLWE4MWQtNGE4YmUzNzQ1NDNkIiwiZXhwIjoxNTUwNjgxMTQ2LCJpZGVudGl0eSI6IjVjNmQ0NmYyMzlkYThlMDA0NTkzMzczYyIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.qAuvOPCQVr3YAFuD6B4RmqLmEokv-05ttL8BzPRBAug
            
    + Body
    
            {
                "name": "Bui Thang",
                "gender": "male",
                "dob": "1994-05-11T00:00:00Z",
                "education": "University"
            }

+ Response 200 (application/json)

        {
            "id": "fVwYk3hcSkpNv7brodYq",
            "name": "Bui Thang",
            "email": "zenthangplus@gmail.com",
            "gender": "male",
            "dob": "1994-05-11T00:00:00Z",
            "education": "University",
            "created_at": "2019-02-25T15:51:42.060512Z"
        }

+ Response 400 (application/json)

        {
            "code": "request_invalid",
            "message": "Couldn't parse your request"
        }


+ Response 401 (application/json)

        {
            "code": "unauthorized",
            "message": "Unauthorized"
        }

+ Response 500 (application/json)

        {
            "code": "update_failed",
            "message": "Couldn't update user"
        }
        
### Get categories available [GET /user/categories]

Get list of categories available.

+ Request (application/json)

    + Headers
        
            Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NTA2ODAyNDYsIm5iZiI6MTU1MDY4MDI0NiwianRpIjoiNGU0OGRhOTUtZDEyMS00MmViLWE4MWQtNGE4YmUzNzQ1NDNkIiwiZXhwIjoxNTUwNjgxMTQ2LCJpZGVudGl0eSI6IjVjNmQ0NmYyMzlkYThlMDA0NTkzMzczYyIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.qAuvOPCQVr3YAFuD6B4RmqLmEokv-05ttL8BzPRBAug
     
+ Response 200 (application/json)

        [
            {
                "id": "5h3fO6lXipYjZGxvLDw9",
                "slug": "gaming",
                "name": "Gaming",
                "sub_categories": [
                    {
                        "id": "dZvHuNP6LRWMImNtM22T",
                        "slug": "gaming-news",
                        "name": "Gaming News"
                    },
                    {
                        "id": "lUhk0q2bM3ZoPbM6Vl2b",
                        "slug": "esport",
                        "name": "eSport"
                    },
                    {
                        "id": "o0rRy49gpeUesf9yD16F",
                        "slug": "games",
                        "name": "Games"
                    },
                    {
                        "id": "opyh5IB8B6BTbFc4HHJA",
                        "slug": "celebrity-gamer",
                        "name": "Celebrity Gamer"
                    }
                ]
            },
            {
                "id": "DocYIB8M57d4qj4vxdz7",
                "slug": "Nonprofits",
                "name": "Nonprofits",
                "sub_categories": [
                    {
                        "id": "3fmUK3L6Gvi9788xGwOK",
                        "slug": "humanitarian",
                        "name": "Humanitarian"
                    }
                ]
            },
            {
                "id": "OuBuXlKxvKBoRnSPgTWB",
                "slug": "government-and-polytics",
                "name": "Government & Polytics",
                "sub_categories": [
                    {
                        "id": "cDoXJnWQGR7rkpwD6Cbm",
                        "slug": "gov-officials-and-agencies",
                        "name": "Gov Officials & Agencies A"
                    }
                ]
            }
        ]

+ Response 500 (application/json)

        {
            "code": "get_failed",
            "message": "Couldn't get list categories"
        }
