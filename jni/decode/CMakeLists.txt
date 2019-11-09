cmake_minimum_required(VERSION 3.15)
project(jni C)

set(CMAKE_C_STANDARD 11)

include_directories(.)
#link_directories("D:\\workspace\\golang\\project\\dhcrypto\\jni")
link_libraries("D:\\workspace\\golang\\project\\dhcrypto\\jni\\crypto.so")

add_executable(jni
        crypto.h
        main.c)

target_link_libraries(${PROJECT_NAME})