FROM docker.adeo.no:5000/go-scratch
MAINTAINER Sten R�kke<sten.ivar.rokke@nav.no>

COPY dist /opt/coregroups
EXPOSE 80
CMD ["/opt/coregroups/coregroups", "-file", "/opt/coregroups/coregroups.json"]
