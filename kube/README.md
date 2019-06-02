# Apply Kubernetes resources

To install Cleanaux Backend run scripts from `/scripts`

## Minikube

To install Cleanaux Backend on minikube run:

```bash
bash ../scripts/install-minikube.sh
```

**Remember** to clean up the minikube after work by running:

```bash
bash ../scripts/cleanup-minikube.sh
```

## Google Cloud Platform

To login into `gcloud` run:

```bash
gcloud auth login
```

To connect to the cluster run the command provided by GCP:

```bash
gcloud container clusters get-credentials {CLUSTER_NAME} --zone {ZONE} --project {PROJECT}
```

To install Cleanaux Backend on GCP run:

```bash
bash ../scripts/install-gcp.sh
```

## Both of them

To open MySQL database client run:

```bash
kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql-database-internal -ppassword
```

To open MongoDB database client run: 

```bash
#TODO
```
