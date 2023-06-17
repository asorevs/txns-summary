#!/bin/bash  

# Build image
docker build -t txn-summary .

# Run Docker container
docker run -it --rm \
  -e SMTP_PORT=$SMTP_PORT \
  -e EMAIL_SENDER=$EMAIL_SENDER \
  -e EMAIL_PASSWORD=$EMAIL_PASSWORD \
  -e EMAIL_RECIPIENT=$EMAIL_RECIPIENT \
  -e SMTP_HOST=$SMTP_HOST \
  txn-summary
