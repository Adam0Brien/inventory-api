#!/bin/bash
set -e

source ./scripts/check_docker_podman.sh

kind create cluster --name inventory-cluster

kubectl create secret docker-registry redhat-registry-secret \
  --docker-server=$DOCKER_SERVER \
  --docker-username=$QUAY_USERNAME \
  --docker-password=$QUAY_PASSWORD \
  --docker-email=$QUAY_EMAIL


kubectl create secret generic spicedb-secrets \
  --from-literal=grpc_preshared_key="foobar" \
  --from-literal=postgres_user="postgres" \
  --from-literal=postgres_password="yPsw5e6ab4bvAGe5H" \
  -n default



kubectl create secret generic local-spicedb-secret --from-literal=token=foobar ##relations secret


kubectl create serviceaccount default
kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "redhat-registry-secret"}]}'

# build/tag image
${DOCKER} build -t localhost/inventory-api:latest -f Dockerfile .
${DOCKER} build -t localhost/inventory-e2e-tests:latest -f Dockerfile-e2e .
${DOCKER} pull quay.io/redhat-services-prod/project-kessel-tenant/kessel-relations/relations-api:latest

${DOCKER} tag localhost/inventory-api:latest localhost/inventory-api:e2e-test
${DOCKER} tag localhost/inventory-e2e-tests:latest localhost/inventory-e2e-tests:e2e-test

rm -rf inventory-api.tar
rm -rf inventory-e2e-tests.tar
rm -rf relations-api.tar

${DOCKER} save -o inventory-api.tar localhost/inventory-api:e2e-test
${DOCKER} save -o inventory-e2e-tests.tar localhost/inventory-e2e-tests:e2e-test
${DOCKER} save -o relations-api.tar quay.io/redhat-services-prod/project-kessel-tenant/kessel-relations/relations-api:latest

kind load image-archive inventory-api.tar --name inventory-cluster
kind load image-archive inventory-e2e-tests.tar --name inventory-cluster
${DOCKER} load -i relations-api.tar
kind load image-archive relations-api.tar --name inventory-cluster


kubectl create configmap inventory-api-psks --from-file=config/psks.yaml

kubectl apply -f https://strimzi.io/install/latest\?namespace\=default

kubectl apply -f deploy/kind/inventory/kessel-inventory.yaml
kubectl apply -f deploy/kind/inventory/invdatabase.yaml
kubectl apply -f deploy/kind/e2e/e2e-batch.yaml

kubectl apply -f deploy/kind/inventory/strimzi.yaml

## deploy relations-api for readyz check
#echo "Deploying relations-api service"
#kubectl apply -f deploy/kind/relations/spicedb.yaml
#kubectl apply -f deploy/kind/relations/kessel-relations.yaml
