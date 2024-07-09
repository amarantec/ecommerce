#!/bin/csh

cd cmd/api

go122 build -o app

./app
