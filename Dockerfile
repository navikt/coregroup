FROM golang:1.9.4-alpine

COPY dist /opt/coregroups
EXPOSE 80
CMD ["/opt/coregroups/coregroups", "-file", "/opt/coregroups/coregroups.json"]
