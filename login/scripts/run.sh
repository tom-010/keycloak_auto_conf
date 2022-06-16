#!/bin/bash

export PORT=8000
export REALM=skytala
export CLIENT_ID=clientid-03
export CLIENT_SECRET=K7t2lxzlp4p0rENZSVFGZr99SFw11zar
export KEYCLOAK_URL=http://auth.example.com

go run server/server.go