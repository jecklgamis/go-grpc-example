#!/bin/bash
SERVER_PORT=${SERVER_PORT:-4000}
cd /app && bin/server --port $SERVER_PORT
