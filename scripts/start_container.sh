#!/bin/bash
echo "Starting Docker container..."
aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com
docker run -d -p 80:8084 --name f-bot -v /var/param/.env:/app/.env $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/f-bot:latest