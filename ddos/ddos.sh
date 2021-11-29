#!/bin/sh
# Docs: https://httpd.apache.org/docs/2.4/programs/ab.html
ab -c 100 -n 1000 http://host.docker.internal:1997/
