add_library(
  common
  src/input_parser.cc
  src/md5.cc
  src/point2d.cc
  src/string_convert.cc
  src/vector_convert.cc
)

target_include_directories(
  common
  PUBLIC
    $<INSTALL_INTERFACE:include>
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  PRIVATE
    ${CMAKE_CURRENT_SOURCE_DIR}/src
)
