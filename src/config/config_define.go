package config

import (
	"flag"
)

var FLAG_CONFIG_FILE = flag.String(
	"config_file", "/root/workspace/mornsail/src/config/flag.config", "Config file to read flags from.")
var FLAG_ADDR = flag.String(
	"ip_eth0", ":6666", "Config file to read flags from.")
var FLAG_REDIS_ADDR = flag.String(
	"redis_addr", ":6379", "Redis address")
var FLAG_REDIS_PASSWD = flag.String(
	"redis_passwd", "", "Redis password")
var FLAG_POSTGRES_DRIVER = flag.String(
        "postgres_driver", "", "postgres driver name")
var FLAG_POSTGRES_SOURCE = flag.String(
        "postgres_source", "", "postgres source")
var FLAG_MONGO_ADDR = flag.String(
        "mongo_addr", "", "Mongo address")
/*
var FLAG_POSTGRES_DRIVER = flag.String(
	"postgres_driver", "postgres", "postgres driver name")
var FLAG_POSTGRES_SOURCE = flag.String(
	"postgres_source", "host:db port:5432 dbname:lyingdragon2 user:postgres sslmode:disable", "postgres source")
var FLAG_MONGO_ADDR = flag.String(
	"mongo_addr", "10.0.3.235:9989", "Mongo address")
*/
