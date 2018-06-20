Coregroup service
=================

Go service used to check which coregroup in Websphere a certain application is supposed to belong to. Used during setup of WAS and deploy of applications.

corgroups.json contains the mapping of applications to coregroups. If the application is not defined in the JSON-file, the DefaultCoreGroup will be used. Returns a string with the coregroup name.
