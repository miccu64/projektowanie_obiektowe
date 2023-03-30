#!/bin/bash

cd .. && symfony -d -q server:start


test_endpoint() {
    local params="$1"
    local expectingResponse="$2"
    local expectedFinalResult="$3"

    echo "Testing method, params and URL: $params"

    result=0
    httpCode=$(curl --write-out '%{http_code}' --silent --output /dev/null -sb -H -X $params)
    if [ "$httpCode" -ge 200 ] && [ "$httpCode" -lt 300 ] ; then
        result=1
        if [ "$expectingResponse" -eq 1 ] ; then
            resultBody="$(curl -sb -H -X $params)"
            echo "Response body: $resultBody"
            if [[ $resultBody != {* ]] && [[ $resultBody != [* ]]; then
                result=0
            fi
        fi
    else
        result=0
    fi
    echo "Test result: $result"
    echo " "
    return "$result"
}


params="GET http://localhost:8000/products"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="GET http://localhost:8000/products"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="DELETE http://localhost:8000/products/2"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="POST -d {\"id\":2,\"name\":\"Produkt1\",\"price\":22222} http://localhost:8000/products"
result=$(test_endpoint "$params" 0 1)
echo "$result"

params="PATCH -d {\"id\":2,\"name\":\"Produkt1\",\"price\":22222} http://localhost:8000/products"
result=$(test_endpoint "$params" 1 1)
echo "$result"



params="GET http://localhost:8000/pizzas"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="GET http://localhost:8000/pizzas"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="DELETE http://localhost:8000/pizzas/2"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="POST -d {\"id\":2,\"type\":\"Pepperoni\",\"cost\":11} http://localhost:8000/pizzas"
result=$(test_endpoint "$params" 0 1)
echo "$result"

params="PATCH -d {\"id\":2,\"type\":\"Studencka\",\"cost\":22} http://localhost:8000/pizzas"
result=$(test_endpoint "$params" 1 1)
echo "$result"



params="GET http://localhost:8000/cars"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="GET http://localhost:8000/cars"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="DELETE http://localhost:8000/cars/2"
result=$(test_endpoint "$params" 1 1)
echo "$result"

params="POST -d {\"id\":2,\"type\":\"Pepperoni\",\"color\":\"white\",\"isSuv\":true} http://localhost:8000/cars"
result=$(test_endpoint "$params" 0 1)
echo "$result"

params="PATCH -d {\"id\":2,\"type\":\"Studencka\",\"color\":\"black\",\"isSuv\":true} http://localhost:8000/cars"
result=$(test_endpoint "$params" 1 1)
echo "$result"


symfony -q local:server:stop
