#!/bin/bash

# CREDIT: https://ryantravitz.com/2018-10-14-mongodb-change-streams/

until curl http://mongo1:27017/serverStatus\?text\=1 2>&1 | grep uptime | head -1; do
  printf '.'
  sleep 1
done

echo curl http://mongo1:27017/serverStatus\?text\=1 2>&1 | grep uptime | head -1
echo 'Started..'

sleep 8

echo time now: `date +"%T" `
mongo --host mongo1:27017 --username root --password password <<EOF
  var cfg = rs.conf();
  cfg.members = [
            {
                _id: 0,
                host: 'mongo1:27017',
                priority: 2
            },
            {
                _id: 1,
                host: 'mongo2:27017',
                priority: 0
            },
            {
                _id: 2,
                host: 'mongo3:27017',
                priority: 0
            },
        ];
    rs.initiate(cfg, { force: true });
    rs.reconfig(cfg, { force: true });
    db.getMongo().setReadPref('nearest');
EOF
