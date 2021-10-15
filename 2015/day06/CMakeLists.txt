# Use the current directory name to name the target.
get_filename_component(TARGET_NAME ${CMAKE_CURRENT_LIST_DIR} NAME)

add_library(
  ${TARGET_NAME}lib
  lib.cc
)

target_link_libraries(
  ${TARGET_NAME}lib
  common
)

target_link_libraries(
  ${TARGET_NAME}lib
  common
)

add_executable(
  ${TARGET_NAME}
  main.cc
)

target_link_libraries(
  ${TARGET_NAME}
  PRIVATE
  ${TARGET_NAME}lib
  common
)

configure_file(input.txt input.txt COPYONLY)

enable_testing()

add_executable(
  ${TARGET_NAME}_test
  test.cc
)

target_link_libraries(
  ${TARGET_NAME}_test
  PRIVATE
  ${TARGET_NAME}lib
  common
  gtest_main
)

include(GoogleTest)
gtest_discover_tests(${TARGET_NAME}_test)
