@echo off
setlocal EnableDelayedExpansion

echo Generating gogo proto code

cd ..
cd proto
buf generate --template buf.gen.gogo.yaml

cd ..

:: move proto files to the right places
xcopy /s /i github.com\cosmos\ibc-go\v*\modules\* modules\
xcopy /s /i github.com\cosmos\ibc-go\modules\* modules\
rmdir /s /q github.com

go mod tidy
