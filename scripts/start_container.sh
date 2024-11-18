#!/bin/bash
echo "Starting Docker container..."

AWS_ACCOUNT_ID="438465148767"
AWS_REGION="us-east-1"

# Login to AWS ECR
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# Pull the latest image from AWS ECR
echo "Pulling the latest image from AWS ECR..."
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/f-bot:latest
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/f-bot-telegram:latest

# Stop and remove any existing container named 'f-bot' (if any)
echo "Stopping and removing any existing 'f-bot' container..."
docker stop f-bot || true
docker rm f-bot || true

docker stop f-bot-telegram || true
docker rm f-bot-telegram || true

# Run the new container
echo "Running the new container..."
docker run -d -p 80:8084 --name f-bot -v /var/param/.env:/app/.env $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/f-bot:latest
docker run -d --name f-bot-telegram -v /var/param/.env:/app/.env $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/f-bot-telegram:latest
