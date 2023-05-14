@echo off
setlocal enabledelayedexpansion

cd ./bots/bin


set CGO_ENABLED=1
set GO111MODULE=on
set GOARM=7
set ANDROID_API_LEVEL=29
set ANDROID_NDK_HOME=C:\Users\0\AppData\Local\Android\Sdk\ndk\25.2.9519653\toolchains\llvm\prebuilt\windows-x86_64\bin
set BUILD_COMMAND=garble -literals -tiny -seed=random build ../src

echo Building for Windows...
%BUILD_COMMAND%
del bots_win.exe
ren src.exe bots_win.exe


set GOOS=android

echo Building for Android (arm64-v8a)...
set GOARCH=arm64
set CC=%ANDROID_NDK_HOME%\aarch64-linux-android%ANDROID_API_LEVEL%-clang.cmd
set CXX=%ANDROID_NDK_HOME%\aarch64-linux-android%ANDROID_API_LEVEL%-clang++.cmd
%BUILD_COMMAND%
del bots_android_arm64-v8a
ren src bots_android_arm64-v8a

echo Building for Android (armeabi-v7a)...
set GOARCH=arm
set CC=%ANDROID_NDK_HOME%\armv7a-linux-androideabi%ANDROID_API_LEVEL%-clang.cmd
set CXX=%ANDROID_NDK_HOME%\armv7a-linux-androideabi%ANDROID_API_LEVEL%-clang++.cmd
%BUILD_COMMAND%
del bots_android_armeabi-v7a
ren src bots_android_armeabi-v7a

echo Building for Android (x86)...
set GOARCH=386
set CC=%ANDROID_NDK_HOME%\i686-linux-android%ANDROID_API_LEVEL%-clang.cmd
set CXX=%ANDROID_NDK_HOME%\i686-linux-android%ANDROID_API_LEVEL%-clang++.cmd
%BUILD_COMMAND%
del bots_android_x86
ren src bots_android_x86

echo Building for Android (x86_64)...
set GOARCH=amd64
set CC=%ANDROID_NDK_HOME%\x86_64-linux-android%ANDROID_API_LEVEL%-clang.cmd
set CXX=%ANDROID_NDK_HOME%\x86_64-linux-android%ANDROID_API_LEVEL%-clang++.cmd
%BUILD_COMMAND%
del bots_android_x86_64
ren src bots_android_x86_64


rem zip it
set rar_path="C:\Program Files\WinRAR\WinRAR.exe"
set name=bots
cd ..
if not exist %name% mkdir %name%
cd bots
xcopy /s /e /y /i /q ..\bin\* .\ >nul 2>nul
if exist %name%.zip del %name%.zip
cd ..
if exist bin\%name%.zip del bin\%name%.zip >nul 2>nul
%rar_path% a %name%.zip %name%\* >nul 2>nul

rem clean up
rmdir /s /q bots >nul 2>nul
move /y %name%.zip bin\%name%.zip >nul 2>nul
