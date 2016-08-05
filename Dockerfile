FROM docker.adeo.no:5000/scratch
MAINTAINER Sten Røkke<sten.ivar.rokke@nav.no>

COPY dist /opt/coregroups
EXPOSE 8443
CMD ["/opt/coregroups/coregroups", "-file /opt/coregroups/coregroups.json", "-cert $TLS_CERT", "-key $TLS_PRIVATE_KEY"]
