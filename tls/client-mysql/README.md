# Connect Securely to MySQL Using TLS

Create a go client that connect to MySQL using TLS.

## MySQL Docker Setup

```shell
$ docker-compose up --build  # Force a rebuild of the local Dockerfile
$ docker exec -it mysql /bin/bash
$ mysql -u root -p --ssl-mode=required
mysql> show variables like 'ssl_%';
+---------------------------+------------------------------------+
| ssl_ca                    | /var/lib/mysql-tls/ca.pem          |
| ssl_capath                |                                    |
| ssl_cert                  | /var/lib/mysql-tls/server-cert.pem |
| ssl_cipher                |                                    |
| ssl_crl                   |                                    |
| ssl_crlpath               |                                    |
| ssl_fips_mode             | OFF                                |
| ssl_key                   | /var/lib/mysql-tls/server-key.pem  |
| ssl_session_cache_mode    | ON                                 |
| ssl_session_cache_timeout | 300                                |
+---------------------------+------------------------------------+
mysql> # If you see the above, then the configuration is set up correctly
mysql> # Also, we can also run \s to show the current connection setup
mysql> \s
```

### Notes

* MySQL provides a program `mysql_ssl_rsa_setup` to generate the **self-signed** certs and **2048-bit RSA** keys needed by MySQL to establish a TLS connection. We pass `/var/lib/mysql-tls` to the `-d` flag as we want to save the tls certs and keys to the custom directory so that we can persist the tls certs and keys after we stop the mysql container. Use this for development only.
* We mount `docker/tls` local directory as `/var/lib/mysql-tls`. The generated tls certs and keys are saved in `docker/tls`.
* If we want to regenerate the keys, we remove the files in `docker/tls` and restart the mysql container.
* If we made changes to the local Dockerfile, we need to remove the local image, rebuild the image, and run the container.

## Go MySQL Client

The Go mysql client will use the client tls ca, cert and key, `ca.pen`, `client-cert.pem`, and `client-key.pem` respectively to establish a tls connection to the mysql server.

## Reference

* [MySQL: Creating SSL RSA files Using MYSQL](https://dev.mysql.com/doc/refman/5.7/en/creating-ssl-rsa-files-using-mysql.html)
* [MySQL: Using Encrypted Connections](https://dev.mysql.com/doc/refman/5.7/en/using-encrypted-connections.html)
* [Github: wixyvir/docker-mysql-tls](https://github.com/wixyvir/docker-mysql-tls)
