#!/bin/bash
docker run -d  --name openldap -e LDAP_ORGANISATION="Test" -e LDAP_DOMAIN=localhost.localdomain -e LDAP_ADMIN_PASSWORD=JonSn0w osixia/openldap 

docker run -d  --link openldap:ldap --name gocd-server -p 8153:8153 gocd/gocd-server

docker run -d  --link gocd-server:go-server  gocd/gocd-agent
docker run -d  --link gocd-server:go-server  gocd/gocd-agent


docker run -d  -p 443:443 --link openldap:ldap -e PHPLDAPADMIN_LDAP_HOSTS=ldap osixia/phpldapadmin
