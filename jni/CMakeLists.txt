cmake_minimum_required(VERSION 3.15)
project(jni C)

set(CMAKE_C_STANDARD 11)

include_directories(.)
link_libraries("D:\\workspace\\golang\\project\\crypto\\jni\\crypto.so")
add_executable(jni
    decode/main.c
    crypto.h)

target_link_libraries(${PROJECT_NAME})

