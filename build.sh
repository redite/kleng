#!/usr/bin/env bash

echo -e "Compiling resounden.jsx"
cd ./public/js
babel --presets "react" -o resounden.js resounden.jsx

echo -e "Running server..."
cd ../../
go run main.go