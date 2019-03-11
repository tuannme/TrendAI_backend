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
API v1 for development server: [https://trendai-bb810.appspot.com/v1/](https://trendai-bb810.appspot.com/v1/)

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
                "education": "University",
                "interest_categories": ["5c7bd24195d42205ad86b4f7", "5c7bd24195d42205ad86b503"]
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

+ Response 400 (application/json)

        {
            "code": "category_not_found",
            "message": "Categories doesn't exists"
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
                "id": "5c7bd24195d42205ad86b4f5",
                "name": "News",
                "slug": "news",
                "child": [
                    {
                        "id": "5c7bd24195d42205ad86b4f6",
                        "name": "Weather",
                        "slug": "news/weather"
                    },
                    {
                        "id": "5c7bd24195d42205ad86b4f7",
                        "name": "History",
                        "slug": "news/history"
                    },
                    {
                        "id": "5c7bd24195d42205ad86b4f8",
                        "name": "Politics",
                        "slug": "news/politics"
                    },
                    {
                        "id": "5c7bd24195d42205ad86b4f9",
                        "name": "Health",
                        "slug": "news/health"
                    }
                ]
            },
            {
                "id": "5c7bd24195d42205ad86b500",
                "name": "Lifestyle",
                "slug": "lifestyle",
                "child": [
                    {
                        "id": "5c7bd24195d42205ad86b501",
                        "name": "Parenting",
                        "slug": "lifestyle/parenting"
                    },
                    {
                        "id": "5c7bd24195d42205ad86b502",
                        "name": "DIY & Home",
                        "slug": "lifestyle/diy-and-home"
                    }
                ]
            }
        ]

+ Response 500 (application/json)

        {
            "code": "get_failed",
            "message": "Couldn't get list categories"
        }


### Get trending topics [GET /trends{?lat,lng}]

Get list of trending topics by location.

+ Parameters
    + lat (optional, float) - Location's latitude
    + lng (optional, float) - Location's longitude

+ Request (application/json)

    + Headers
        
            Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1NTA2ODAyNDYsIm5iZiI6MTU1MDY4MDI0NiwianRpIjoiNGU0OGRhOTUtZDEyMS00MmViLWE4MWQtNGE4YmUzNzQ1NDNkIiwiZXhwIjoxNTUwNjgxMTQ2LCJpZGVudGl0eSI6IjVjNmQ0NmYyMzlkYThlMDA0NTkzMzczYyIsImZyZXNoIjpmYWxzZSwidHlwZSI6ImFjY2VzcyJ9.qAuvOPCQVr3YAFuD6B4RmqLmEokv-05ttL8BzPRBAug


+ Response 200 (application/json)

        [
            {
                "id": "5c867826f4e15b63d1954e59",
                "name": "Seungri",
                "volume": 564238
            },
            {
                "id": "5c867826f4e15b63d1954e5a",
                "name": "#EpikHighxSUGA",
                "volume": 108667
            },
            {
                "id": "5c867826f4e15b63d1954e5b",
                "name": "#PerthFMTinVietNam",
                "volume": 46406
            },
            {
                "id": "5c867826f4e15b63d1954e5c",
                "name": "Jimin",
                "volume": 1253731
            },
            {
                "id": "5c867826f4e15b63d1954e5d",
                "name": "#2wish",
                "volume": 51056
            }
        ]

+ Response 400 (application/json)

        {
            "code": "get_failed",
            "message": "Couldn't get trending data."
        }

+ Response 500 (application/json)

        {
            "code": "get_failed",
            "message": "Couldn't get trending data."
        }
