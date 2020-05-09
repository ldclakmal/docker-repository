# How to Start LDAP Server with User Data?

First, if you have already started the `osixia/openldap:1.3.0` image, you have to delete the container. If not, ignore this step.
```bash
$ docker container ls -a
$ docker container rm <container-id>
```

Now, you need to create `.ldif` with user data. Refer to [bootstrap.ldif](./bootstrap.ldif) for a sample.
> Reference: https://www.digitalocean.com/community/tutorials/how-to-use-ldif-files-to-make-changes-to-an-openldap-system

Now, create `docker-compose.yml` file with required env variables, ports, volumes etc. Refer to [docker-compose.yml](./docker-compose.yml) for a sample.

Now, execute following command to start the Docker container.
```bash
$ docker-compose up
```

Execute following commands with `ldapsearch` to query the DIT and lookup entries.
```bash
$ docker exec openldap-server ldapsearch -x -H ldap://localhost -D "cn=admin,dc=avix,dc=lk" -w avix123 -b "dc=avix,dc=lk"
$ docker exec openldap-server ldapsearch -x -H ldap://localhost -D "cn=admin,dc=avix,dc=lk" -w avix123 -b "dc=avix,dc=lk" -s one -LLL dn
$ docker exec openldap-server ldapsearch -x -H ldap://localhost -D "cn=admin,dc=avix,dc=lk" -w avix123 -b "dc=avix,dc=lk" -LLL "(uid=ldclakmal)"
```

Use this command to verify the `bootstrap.ldif` is properly installed.
```bash
$ docker exec openldap-server ldapadd -x -H ldap://localhost -D "cn=admin,dc=avix,dc=lk" -w avix123 -f /container/service/slapd/assets/config/bootstrap/ldif/custom/bootstrap.ldif
```

# References
https://github.com/osixia/docker-openldap
https://ldap.com/basic-ldap-concepts/
https://www.digitalocean.com/community/tutorials/how-to-use-ldif-files-to-make-changes-to-an-openldap-system
https://www.digitalocean.com/community/tutorials/how-to-manage-and-use-ldap-servers-with-openldap-utilities
