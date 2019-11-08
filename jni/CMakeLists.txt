cmake_minimum_required(VERSION 3.15)
project(jni C)

set(CMAKE_C_STANDARD 11)

include_directories(.)
link_directories(.)
link_libraries(./crypto.a)

add_executable(jni
        crypto.h
        main.c)


#add_library(crypto SHARED)
target_link_libraries(${PROJECT_NAME} crypto)