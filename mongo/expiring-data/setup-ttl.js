// Turn on the TTLMonitor thread to remove expired data set by the TTL index.
db.adminCommand({
    setParameter: 1,
    ttlMonitorEnabled: true,
});

// Set the TTLMonitor log to be report the output of the TTLMonitor.
// Set the log level at 1 for the right amount of info.
// db.setLogLevel(1, "index");

// Run the TTLMonitor every minute. Default is 60 so this is set so that we may change it later.
db.adminCommand({
    setParameter: 1,
    ttlMonitorSleepSecs: 60,
});
