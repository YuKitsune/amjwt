#!/bin/bash

privFileName="private_key.pem"
pubFileName="public_key.pem"

rm -f $privFileName
rm -f $pubFileName

openssl ecparam -genkey -name prime256v1 -noout -out $privFileName
openssl ec -in $privFileName -pubout -out $pubFileName