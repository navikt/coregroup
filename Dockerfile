FROM docker.adeo.no:5000/go-scratch
MAINTAINER Sten Røkke<sten.ivar.rokke@nav.no>

COPY dist /opt/coregroups

#ADD server.key /etc/pki/tls/private/
#ADD server.crt /etc/pki/tls/certs/
EXPOSE 8443
CMD ["/opt/coregroups/coregroups", "-file /opt/coregroups/coregroups.json", "-cert $TLS_CERT", "-key $TLS_PRIVATE_KEY"]
