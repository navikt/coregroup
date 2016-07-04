dockerHostConfig="--add-host=dockerhost:$(/sbin/ifconfig docker0 | awk '{ if ( $1 == "inet" ) {print $2}}')"
portConfig="-p 8080:8080 -p 8443:8443"
imageName="docker.adeo.no:5000/coregroups:0.1"
sudo docker run --restart=always -d ${dockerHostConfig} ${portConfig} ${imageName}
