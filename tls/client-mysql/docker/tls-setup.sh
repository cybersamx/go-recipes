#!/bin/bash

mysql_ssl_rsa_setup -d /var/lib/mysql-tls

cat << EOF > /etc/mysql/conf.d/ssl.cnf
[mysqld]
require_secure_transport = ON
ssl-ca=/var/lib/mysql-tls/ca.pem
ssl-cert=/var/lib/mysql-tls/server-cert.pem
ssl-key=/var/lib/mysql-tls/server-key.pem
EOF
