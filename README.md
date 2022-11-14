# ShortLink
ShortLink generates a corresponding shortlink for a given Longlink. The shortlinks generated are stored in a caching service of redis.

### Go Mods
This template uses [Go Mods](https://github.com/golang/go/wiki/Modules) to manage dependencies. All 
external dependencies are in the `go.mod` file.
### Routing
The router we are using is the [Gorilla Mux](https://github.com/gorilla/mux) router to serve the http routes which is present inside the router.go


### Running Locally
    1. Make sure you have the RedisDB installed. For installation run the docker-compose file present.
        i. `docker-compose -f docker-compose.yml up` 

    2. If fails to run on Docker then cd to cmd/ and type the command go run main.go

    3. To call the service for generating short-urls, hit the endpoint with POST method given below and pass the body as provided.
        i. `http://localhost:8000/shorten-url`

        ii. {
                "longURL": "https://github.com/Innocect/InfraCloudAssignment/tree/master"   
            }
    
    4. Please pass a valid LongURL, in accordance with [RFC3987]. If the URL is not valid 500 response code will be returned

    5. After successful service request, a unique short-url, along with the long-url will be returned.
        i. { 
            "longURL":"https://github.com/Innocect/InfraCloudAssignment/tree/master",
            "shortURL":"https://innocect/XFFg"
            }