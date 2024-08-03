#!/usr/bin/bash

curl -i -X UPDATE -H "Content-Type: application/json" http://localhost:8080/update-category/1 -d '{"name": "Test Update", "url": "test-update" }'
