cmake_minimum_required(VERSION 3.10)
project(MyTests)

# Manually specify paths for gtest
include_directories(/opt/homebrew/include)
link_directories(/opt/homebrew/lib)

add_executable(runTests 1480.cpp)
target_link_libraries(runTests gtest gtest_main)
