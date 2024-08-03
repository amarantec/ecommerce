#!/usr/bin/bash

curl -i -X POST -H "Content-Type: application/json" http://localhost:8080/new-category -d '{"name": "Books", "url": "books"}'
