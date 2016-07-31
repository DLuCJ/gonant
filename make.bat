@echo off
REM TODO: version checking

python --version 1>NUL 2>NUL
IF errorlevel 1 goto INSTALL 
python sdgen.py > gonant/songdata.go

:INSTALL
go install 

IF errorlevel 1 (
   %GOPATH%\bin\gonant -load
   goto END
)

%GOPATH%\bin\gonant

:END


