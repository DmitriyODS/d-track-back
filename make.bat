@echo off

echo Create folder bin ...
mkdir bin

echo Copy env to bin ...
copy .env.template bin\.env

echo Build app ...
go build -mod vendor -v -o ./bin/dtrack_back.exe ./cmd

echo All done! Your app is filder bin

pause