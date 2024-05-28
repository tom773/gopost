## Gopost - A simple CLI tool for testing API endpoints

### Usage
```
go build main.go
./main -m <method> -j <key:value> -e <endpoint w/o preceeding slash> <url> <port>
```
<br><hr><br>
### Example
```
./main -m GET -e api/v2/users -j id:abcxyz localhost 42069
```
### Output
```
$\color{green}{\textsf{Server is healthy!}}$

 POST
 localhost :  8090
 Body:   {"userid":"lho2ni111klqs1i"}
 Response:   {"avatar":"burgers_S9pUSnd1iP.webp"}
```
