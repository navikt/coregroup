dockerHostConfig="--add-host=dockerhost:$(/sbin/ifconfig docker0 | grep \"inet addr\" | sed -r \"s/.*inet addr:([0-9.]*).*$/\\1/\")"
portConfig="-p 8080:8080 -p 8443:8443"
imageName="docker.adeo.no:5000/coregroups:0.1"
sudo docker run --restart=always -d ${dockerHostConfig} ${portConfig} ${imageName}