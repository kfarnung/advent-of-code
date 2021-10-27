cmake_minimum_required(VERSION 3.14)
project(advent_of_code_2015)

include(FetchContent)

FetchContent_Declare(
  googletest
  URL https://github.com/google/googletest/archive/609281088cfefc76f9d0ce82e1ff6c30cc3591e5.zip
)

FetchContent_Declare(
  json
  GIT_REPOSITORY https://github.com/ArthurSonzogni/nlohmann_json_cmake_fetchcontent
  GIT_TAG v3.10.4
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
