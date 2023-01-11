FROM postgres:15.1

# Can't mount the TLS server/ca key/cert as postgres expect the files
# to be downed by the database root user. On the mac, mounted volume
# assumes the uid of the host mac. Hence building this Dockerfile by
# copying the tls files makes more sense.

COPY --chown=postgres:postgres ./postgresql/ /etc/postgresql

RUN chmod 600 /etc/postgresql/server-key.pem
