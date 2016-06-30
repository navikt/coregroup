FROM docker.adeo.no:5000/go-scratch
MAINTAINER Sten Røkke<sten.ivar.rokke@nav.no>

WORKDIR /src
ADD ./dist .
#ADD server.key /etc/pki/tls/private/
#ADD server.crt /etc/pki/tls/certs/
EXPOSE 8080
CMD ["./coregroups"]
