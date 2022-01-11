# assessment
## To run the program follow the below steps
### 1. git clone https://github.com/sderohan/assessment.git
### 2. cd assessment
### 3. go run main.go
### ENDPOINT : http://localhost:8000/v1/get-words
### REQUEST BODY
```
{
    "text": "Hello! Welcome to the world of the programming. This the example string to test the api "
}
```

### EXAMPLE RESPONSE
```
{
    "data": [
        {
            "word": "the",
            "count": 4
        },
        {
            "word": "to",
            "count": 2
        },
        {
            "word": "programming.",
            "count": 1
        },
        {
            "word": "This",
            "count": 1
        },
        {
            "word": "example",
            "count": 1
        },
        {
            "word": "api",
            "count": 1
        },
        {
            "word": "Hello!",
            "count": 1
        },
        {
            "word": "Welcome",
            "count": 1
        },
        {
            "word": "world",
            "count": 1
        },
        {
            "word": "of",
            "count": 1
        }
    ],
    "code": 200
}
```
