#!/bin/bash

set -e

./scripts/generate_grpc.sh
docker image build -t localhost:5000/phybe-login .
docker push localhost:5000/phybe-login