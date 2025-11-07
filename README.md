# simple-go-http-server

## Run the file
### Run the go file
```sh
go run main.go
```
### Check output
you should see the output in your terminal
```sh
server running on :8080
```

## Test the code
You could use curl but in this example i will use pip httpie to make the command shorter
```sh
pip install httpie
```

### POST request
```sh
http http://localhost:8080/Users name=Alice
```

### GET request
```sh
http http://localhost:8080/Users
```

### DELETE request
```sh
http DELETE http://localhost:8080/Users/1
```
