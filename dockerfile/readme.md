Boinc Client container.
========================

build via:

cd {to Dockerfile directory}
docker build -t {dockerhub_accout/boinc} .


run via docker

docker run -ti -d --name {containerName} -e "boincurl=URL" -e "boinckey=#########" {dockerhub_account/boinc}

check logs

docker logs {containerName}

login and look around.
docker exec -ti {containerName} /bin/bash


** updating to include a golang based container lauch controller that takes a json file for inputs
