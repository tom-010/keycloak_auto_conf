#!/bin/bash

curl -X POST -d "login=username&password=wrong_password" "http://login.example.com"

