# Cleanaux Backend

**WORK IN PROGRESS** - It's just my Go, Docker and Kubernetes playground for now. After some work on authorization, monitoring and acceptance tests it will become a cloud backend for the Cleanaux app. It's public because of specific database initialization which requires downloading files.

Cleanaux provides a list of things which should be cleaned regularly. If something wasn't cleaned for a long time it appears with a warning color. After cleaning whole list is updated.

Whole backend is deployed on Google Cloud Platform with Google Kubernetes Engine.

## Content Service

Go microservice exposing REST API via Ingress. It uses MySQL database to store data. Database runs on separate pod and is initialized with an init container.

## Structure

- `*-service/` directory contains specific service files
- `kube/` directory contains Kubernetes files for cloud deployment
- `scripts/` directory contains installation scripts
- `test/` directory contains some testing tips

## To do

- add CI jobs
- add authorization method in Auth Service, enable logging to the app and create auth middleware for Content Service
- delete unused Entry Service
- create Swagger API description for both services
- add monitoring for both services and for databases
- add acceptance tests scenarios
- create end to end test app