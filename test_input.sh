#!/bin/bash
input=$1
while IFS= read -r line
do
  res=`curl --location --request POST 'localhost:8080/api/fund' \
--header 'Content-Type: application/json' \
--data-raw "$line" 2> /dev/null`
  echo $res
done < "$input"