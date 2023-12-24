cmake_minimum_required(VERSION 3.14)
project(advent_of_code_2021)

if(POLICY CMP0135)
  cmake_policy(SET CMP0135 NEW)
endif()

include(FetchContent)

FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/f8d7d77c06936315286eb55f8de22cd23c188571.zip
)

# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

FetchContent_MakeAvailable(googletest)

# GoogleTest requires at least C++11
set(CMAKE_CXX_STANDARD 11)

add_subdirectory(libcommon)

file(GLOB DAYDIRS RELATIVE ${CMAKE_CURRENT_LIST_DIR} day*)
foreach(subdir ${DAYDIRS})
  add_subdirectory(${subdir})
endforeach()

enable_testing()
