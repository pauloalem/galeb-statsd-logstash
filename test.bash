#! /bin/bash

echo "statsite.router.appname.some.requestTime:1|ms|@1.0" | nc -w 1 -u 0.0.0.0 8125
