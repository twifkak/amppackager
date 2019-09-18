# Install script for directory: /usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3

# Set the install prefix
if(NOT DEFINED CMAKE_INSTALL_PREFIX)
  set(CMAKE_INSTALL_PREFIX "/usr/local")
endif()
string(REGEX REPLACE "/$" "" CMAKE_INSTALL_PREFIX "${CMAKE_INSTALL_PREFIX}")

# Set the install configuration name.
if(NOT DEFINED CMAKE_INSTALL_CONFIG_NAME)
  if(BUILD_TYPE)
    string(REGEX REPLACE "^[^A-Za-z0-9_]+" ""
           CMAKE_INSTALL_CONFIG_NAME "${BUILD_TYPE}")
  else()
    set(CMAKE_INSTALL_CONFIG_NAME "Release")
  endif()
  message(STATUS "Install configuration: \"${CMAKE_INSTALL_CONFIG_NAME}\"")
endif()

# Set the component getting installed.
if(NOT CMAKE_INSTALL_COMPONENT)
  if(COMPONENT)
    message(STATUS "Install component: \"${COMPONENT}\"")
    set(CMAKE_INSTALL_COMPONENT "${COMPONENT}")
  else()
    set(CMAKE_INSTALL_COMPONENT)
  endif()
endif()

# Install shared libraries without execute permission?
if(NOT DEFINED CMAKE_INSTALL_SO_NO_EXE)
  set(CMAKE_INSTALL_SO_NO_EXE "1")
endif()

# Is this installation the result of a crosscompile?
if(NOT DEFINED CMAKE_CROSSCOMPILING)
  set(CMAKE_CROSSCOMPILING "FALSE")
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  foreach(file
      "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so.1.0.26"
      "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so.1"
      )
    if(EXISTS "${file}" AND
       NOT IS_SYMLINK "${file}")
      file(RPATH_CHECK
           FILE "${file}"
           RPATH "")
    endif()
  endforeach()
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib" TYPE SHARED_LIBRARY FILES
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/liburiparser.so.1.0.26"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/liburiparser.so.1"
    )
  foreach(file
      "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so.1.0.26"
      "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so.1"
      )
    if(EXISTS "${file}" AND
       NOT IS_SYMLINK "${file}")
      if(CMAKE_INSTALL_DO_STRIP)
        execute_process(COMMAND "/usr/bin/strip" "${file}")
      endif()
    endif()
  endforeach()
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so")
    file(RPATH_CHECK
         FILE "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so"
         RPATH "")
  endif()
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib" TYPE SHARED_LIBRARY FILES "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/liburiparser.so")
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so")
    if(CMAKE_INSTALL_DO_STRIP)
      execute_process(COMMAND "/usr/bin/strip" "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/liburiparser.so")
    endif()
  endif()
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/include/uriparser" TYPE FILE FILES
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/include/uriparser/UriBase.h"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/include/uriparser/UriDefsAnsi.h"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/include/uriparser/UriDefsConfig.h"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/include/uriparser/UriDefsUnicode.h"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/include/uriparser/Uri.h"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/include/uriparser/UriIp4.h"
    )
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse")
    file(RPATH_CHECK
         FILE "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse"
         RPATH "")
  endif()
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/bin" TYPE EXECUTABLE FILES "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/uriparse")
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse" AND
     NOT IS_SYMLINK "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse")
    file(RPATH_CHANGE
         FILE "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse"
         OLD_RPATH "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build:"
         NEW_RPATH "")
    if(CMAKE_INSTALL_DO_STRIP)
      execute_process(COMMAND "/usr/bin/strip" "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/bin/uriparse")
    endif()
  endif()
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3" TYPE FILE FILES
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/cmake/uriparser-config.cmake"
    "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/cmake/uriparser-config-version.cmake"
    )
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  if(EXISTS "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3/uriparser.cmake")
    file(DIFFERENT EXPORT_FILE_CHANGED FILES
         "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3/uriparser.cmake"
         "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/CMakeFiles/Export/lib/cmake/uriparser-0.9.3/uriparser.cmake")
    if(EXPORT_FILE_CHANGED)
      file(GLOB OLD_CONFIG_FILES "$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3/uriparser-*.cmake")
      if(OLD_CONFIG_FILES)
        message(STATUS "Old export file \"$ENV{DESTDIR}${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3/uriparser.cmake\" will be replaced.  Removing files [${OLD_CONFIG_FILES}].")
        file(REMOVE ${OLD_CONFIG_FILES})
      endif()
    endif()
  endif()
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3" TYPE FILE FILES "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/CMakeFiles/Export/lib/cmake/uriparser-0.9.3/uriparser.cmake")
  if("${CMAKE_INSTALL_CONFIG_NAME}" MATCHES "^([Rr][Ee][Ll][Ee][Aa][Ss][Ee])$")
    file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib/cmake/uriparser-0.9.3" TYPE FILE FILES "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/CMakeFiles/Export/lib/cmake/uriparser-0.9.3/uriparser-release.cmake")
  endif()
endif()

if("x${CMAKE_INSTALL_COMPONENT}x" STREQUAL "xUnspecifiedx" OR NOT CMAKE_INSTALL_COMPONENT)
  file(INSTALL DESTINATION "${CMAKE_INSTALL_PREFIX}/lib/pkgconfig" TYPE FILE FILES "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/liburiparser.pc")
endif()

if(CMAKE_INSTALL_COMPONENT)
  set(CMAKE_INSTALL_MANIFEST "install_manifest_${CMAKE_INSTALL_COMPONENT}.txt")
else()
  set(CMAKE_INSTALL_MANIFEST "install_manifest.txt")
endif()

string(REPLACE ";" "\n" CMAKE_INSTALL_MANIFEST_CONTENT
       "${CMAKE_INSTALL_MANIFEST_FILES}")
file(WRITE "/usr/local/google/home/twifkak/.go/src/github.com/ampproject/amppackager/vendor/uriparser-0.9.3/build/${CMAKE_INSTALL_MANIFEST}"
     "${CMAKE_INSTALL_MANIFEST_CONTENT}")
