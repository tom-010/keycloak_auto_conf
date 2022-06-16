#!/bin/bash

set -e

export $(cat .env | xargs)
dockerComposeName=${PWD##*/}
#docker exec ${dockerComposeName}_auth_1 /opt/jboss/keycloak/bin/add-user-keycloak.sh -u $KEYCLOAK_USER -p $KEYCLOAK_PASSWORD
docker exec ${dockerComposeName}_auth_1 cat /tmp/out
#docker-compose restart auth