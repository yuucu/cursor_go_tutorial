#!/bin/bash

echo "=== Hello API Test ==="
echo

echo "Test 1: English hello"
curl -X POST http://localhost:1323/hello \
  -H "Content-Type: application/json" \
  -d '{"message": "hello"}' \
  -w "\n"

echo
echo "Test 2: Japanese greeting"
curl -X POST http://localhost:1323/hello \
  -H "Content-Type: application/json" \
  -d '{"message": "こんにちは世界！"}' \
  -w "\n"

echo
echo "Test 3: Custom message"
curl -X POST http://localhost:1323/hello \
  -H "Content-Type: application/json" \
  -d '{"message": "Hello from Go API!"}' \
  -w "\n"

echo
echo "=== Test Complete === 