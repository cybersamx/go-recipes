# Connect Securely to Postgres Using TLS

Create a go client that connect to Postgres using TLS.

## Postgres Docker Setup

1. Run a script to generate ca, server, and client certs and keys in the `postgresql` directory.

   ```shell
   cd docker
   ./postgresql-setup.sh
   ```

1. Run docker-compose to build the local Dockerfile, which uses the `postgres`
 base image and copy the generated certs/keys over to the container image.

   ```shell
   docker-compose up
   ```

1. To test, run the following:

   ```shell
   $ docker exec -it postgres /bin/bash
   $ psql -h localhost -U postgres
   postgres=# select pg_ssl.pid, pg_ssl.ssl, pg_ssl.version,
              pg_sa.backend_type, pg_sa.usename, pg_sa.client_addr
              from pg_stat_ssl pg_ssl
              join pg_stat_activity pg_sa
              on pg_ssl.pid = pg_sa.pid;
    pid | ssl | version |  backend_type  | usename  | client_addr
   -----+-----+---------+----------------+----------+-------------
    124 | t   | TLSv1.3 | client backend | postgres | 127.0.0.1
   postgres=# # Also, we can also run \s to show the current connection setup
   postgres=# \conninfo
   You are connected to database "postgres" as user "postgres" on host "localhost" (address "127.0.0.1") at port "5432".
   SSL connection (protocol: TLSv1.3, cipher: TLS_AES_256_GCM_SHA384, compression: off)
   ```

   If psql is installed locally, we test the certificate auth-based connection from another ip address by run the following:

   ```shell
   $ psql 'host=localhost port=5432 user=postgres sslmode=verify-full sslrootcert=docker/postgresql/ca-cert.pem sslcert=docker/postgresql/client-cert.pem sslkey=docker/postgresql/client-key.pem'
   psql (15.1 (Debian 15.1-1.pgdg110+1))
   SSL connection (protocol: TLSv1.3, cipher: TLS_AES_256_GCM_SHA384, compression: off)
   postgres=#
   ```

### Notes

* We generate the tls certs/keys using openssl and then copy the self-signed ca and server certs/keys to the postgres container.
* Instead of building of own custom postgres container and copying the certs/keys over, we could have opted for mounting the `postgresql` directory onto the postgres container. The problem is that `postgres` requires the server key to be read only and must be owned by the postgres root user. And docker volume mount uses the uid of the host system. So we have to build our own Dockerfile as a result.
* The script used to generate the tls certs/keys are used only for development. We need to harden the process of generating these certs/keys and configuration of postgres if we were to use this in production.

## Reference

* [pgx - Go PostgreSQL Driver and Toolkit](https://github.com/jackc/pgx)
* [PostgreSQL: Secure TCP/IP Connection with SSL](https://www.postgresql.org/docs/current/ssl-tcp.html)
* [PosgreSQL: SSL Support](https://www.postgresql.org/docs/current/libpq-ssl.html)
* [PostgreSQL Docs: Runtime Config File Locations](https://www.postgresql.org/docs/current/runtime-config-file-locations.html)
