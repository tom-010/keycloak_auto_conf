#!/bin/bash

set -e 

./keycloak/realms/insert_env.py
docker-compose up --build --remove-orphans
