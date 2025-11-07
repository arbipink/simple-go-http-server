# simple-go-http-server

## Run the file
### Run the go file
'''
go run main.go
'''
### Check output
you should see the output in your terminal
'''
server running on :8080
'''

## Test the code
You could use curl but in this example i will use pip httpie to make the command shorter
'''
pip install httpie
'''

### POST request
'''
http http://localhost:8080/Users name=Alice
'''

### GET request
'''
http http://localhost:8080/Users
'''

### DELETE request
'''
http DELETE http://localhost:8080/Users/1
'''
