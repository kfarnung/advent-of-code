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

target_compile_options(
  ${TARGET_NAME}
  PRIVATE
    $<$<CXX_COMPILER_ID:MSVC>:/W4 /WX>
    $<$<NOT:$<CXX_COMPILER_ID:MSVC>>:-Wall -Wextra -Wpedantic -Werror>
)

configure_file(../../private/inputs/2021/${TARGET_NAME}.txt input.txt COPYONLY)

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

target_compile_options(
  ${TARGET_NAME}_test
  PRIVATE
    $<$<CXX_COMPILER_ID:MSVC>:/W4 /WX>
    $<$<NOT:$<CXX_COMPILER_ID:MSVC>>:-Wall -Wextra -Wpedantic -Werror>
)

include(GoogleTest)
gtest_discover_tests(${TARGET_NAME}_test)
