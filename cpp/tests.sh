#!/bin/bash

# Check if argument is passed
if [ -z "$1" ]; then
    echo "Specify a C++ file containing test cases."
    exit 1
fi

CPP_FILE=$1

# Use sed to replace the template.cpp in
# CMakeLists.txt with the given cpp file.

# Linux
# sed -i "" "s/\(add_executable(runTests\)[^)]*\)/\1$CPP_FILE)/" CMakeLists.txt

# MacOS
# sed -i "" "s/add_executable(runTests([^)]*)/add_executable(runTests($CPP_FILE)/" CMakeLists.txt
sed -i '' "s/add_executable(runTests [^)]*)/add_executable(runTests $CPP_FILE)/" CMakeLists.txt

# Configure Project
cmake -S . -B build
# This effectively command sets up the environment
# so that you can later build the project.

# Build the Project
cmake --build build
# This step compiles the source files,
# links them, and generates your executable
# (e.g., runTests).

# Run the Test Executable
./build/runTests
