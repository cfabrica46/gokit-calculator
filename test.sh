#!/bin/bash

echo "Add Operation"
curl -X POST http://localhost:8080/add -d '{"v1":"4","v2":"10"}'

echo "Subtract Operation"
curl -X POST http://localhost:8080/subtract -d '{"v1":"4","v2":"10"}'
