#!/bin/bash

echo "Cleaning up old containers (if any)..."
docker compose down --remove-orphans --volumes

echo "Starting experiment..."
docker compose up --build
