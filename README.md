## Gopost - A simple CLI tool for testing API endpoints

### Usage
```
go build main.go
./main -m <method> -j <key:value> -e <endpoint w/o preceeding slash> <url> <port>
```
<br><hr><br>
### Example
```
./main -m GET -e api/v2/user/avatar -j id:abcxyz localhost 42069
```
### Output
```
Server is healthy!

 GET
 localhost :  42069
 Body:   {"id":"abcxyz"}
 Response:   {"avatar":"burgers.webp"}
```
