cmake_minimum_required(VERSION 3.14)
project(advent_of_code_2015)

cmake_policy(SET CMP0135 NEW)

include(FetchContent)

FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/f8d7d77c06936315286eb55f8de22cd23c188571.zip
)

FetchContent_Declare(
  json
  URL https://github.com/nlohmann/json/releases/download/v3.11.2/json.tar.xz
)

# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

FetchContent_MakeAvailable(googletest)
FetchContent_MakeAvailable(json)

# GoogleTest requires at least C++11
set(CMAKE_CXX_STANDARD 11)

add_subdirectory(libcommon)

file(GLOB DAYDIRS RELATIVE ${CMAKE_CURRENT_LIST_DIR} day*)
foreach(subdir ${DAYDIRS})
  add_subdirectory(${subdir})
endforeach()

enable_testing()
