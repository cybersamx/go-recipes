FROM mysql:8

# Set up the directory where tls certs and keys are saved.
RUN chown mysql:mysql /etc/mysql/conf.d/
RUN mkdir /var/lib/mysql-tls
RUN chown mysql:mysql /var/lib/mysql-tls

# Run the mysql_ssl_rsa_setup to generate self-signed certs and keys
# in the directory created above.
ADD tls-setup.sh /docker-entrypoint-initdb.d/
