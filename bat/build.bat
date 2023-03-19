set currentDir=%cd%

cd %~dp0

REM build batch scp program
cd ..\batch-scp
go build -o bscp.exe batch-scp.go
move .\bscp.exe ..\dist\

REM build batch ssh program
cd ..\batch-ssh
go build -o bssh.exe batch-ssh.go
move .\bssh.exe ..\dist\

cd %currentDir%
echo "build done"