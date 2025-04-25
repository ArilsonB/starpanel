#!/bin/bash
redis-server --daemonize no
# Other startup commands
exec "$@"