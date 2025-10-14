cmake_minimum_required(VERSION 3.14)
project(advent_of_code_2015)

if(POLICY CMP0135)
  cmake_policy(SET CMP0135 NEW)
endif()

# GoogleTest requires at least C++17
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include(FetchContent)

FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/52eb8108c5bdec04579160ae17225d66034bd723.zip
)

FetchContent_Declare(
  json
  URL https://github.com/nlohmann/json/releases/download/v3.12.0/json.tar.xz
)

# For Windows: Prevent overriding the parent project's compiler/linker settings
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

FetchContent_MakeAvailable(googletest)
FetchContent_MakeAvailable(json)

add_subdirectory(libcommon)

file(GLOB DAYDIRS RELATIVE ${CMAKE_CURRENT_LIST_DIR} day*)
foreach(subdir ${DAYDIRS})
  add_subdirectory(${subdir})
endforeach()

enable_testing()
