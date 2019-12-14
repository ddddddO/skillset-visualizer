IMAGE=""

run:
	cd app && API_HOST=localhost npm run dev

open:
	cmd.exe /c start http://localhost:8080/#/bar-chart

connpg:
	psql -U postgres -h localhost -d skillsets -p 5432

# gcsへのイメージのpush/pull
# https://cloud.google.com/container-registry/docs/pushing-and-pulling?hl=ja
# example: make buildpush IMAGE=app
buildpush:
	gcloud auth configure-docker
	docker build -t us.gcr.io/work1111/sv-${IMAGE} -f deployments/dockerfiles/${IMAGE}/Dockerfile .
	docker push us.gcr.io/work1111/sv-${IMAGE}
