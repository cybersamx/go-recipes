#!/usr/bin/env bash

mongo -- "${MONGO_INITDB_DATABASE}" <<EOF
    var rootUser = '${MONGO_INITDB_ROOT_USERNAME}';
    var rootPassword = '${MONGO_INITDB_ROOT_PASSWORD}';
    var adminDB = db.getSiblingDB('admin');
    adminDB.auth(rootUser, rootPassword);

    var user = '${MONGO_INITDB_USERNAME}';
    var password = '${MONGO_INITDB_PASSWORD}';
    db.createUser({
        user: user,
        pwd: password,
        roles: [
            {
                role: 'readWrite',
                db: '${MONGO_INITDB_DATABASE}',
            }
        ]});
EOF
