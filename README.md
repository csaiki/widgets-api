## Docker
```
docker-compose build
docker-compose up
```

## Local
``` bash
cd $GOPATH/src/github.com/csaiki/widgets-api # go to the project folder
mysql -uchris -p1234 < ./sql/00_schema.sql # load mysql schema
mysql -uchris -p1234 < ./sql/01_data.sql # load mysql data
go get # get dependencies
go build # build the binary
./widgets-api # run the binary
```

## Auth
```
username: chris
password: 1234
```

## Routes
GET http://localhost:8080/widgets
```
curl -X GET \
  http://localhost:8080/widgets \
  -H 'authorization: Basic Y2hyaXM6MTIzNA=='
```  
GET http://localhost:8080/widgets/99
```
curl -X GET \
  http://localhost:8080/widgets/99 \
  -H 'authorization: Basic Y2hyaXM6MTIzNA=='
```  

POST http://localhost:8080/widgets
```
curl -X POST \
  http://localhost:8080/widgets \
  -H 'authorization: Basic Y2hyaXM6MTIzNA==' \
  -H 'content-type: application/json' \
  -d '[{
	"name": "Super Ugly Shoes",
    "color": "black",
    "price": "101.00",
    "inventory": 25,
    "melts": false
},
{
	"name": "Nice Yummy Chocolate",
    "color": "white",
    "price": "1.00",
    "inventory": 2600,
    "melts": true
}]'
```
PUT http://localhost:8080/widgets/99
```
curl -X PUT \
  http://localhost:8080/widgets/99 \
  -H 'authorization: Basic Y2hyaXM6MTIzNA==' \
  -H 'content-type: application/json' \
  -d '{
    "name": "Practical Frozen Chair",
    "color": "azure",
    "price": "93.00",
    "melts": false
}'
```
GET http://localhost:8080/users
```
curl -X GET \
  http://localhost:8080/users \
  -H 'authorization: Basic Y2hyaXM6MTIzNA=='
```

GET http://localhost:8080/users/99
```
curl -X GET \
  http://localhost:8080/users/99 \
  -H 'authorization: Basic Y2hyaXM6MTIzNA=='
```
