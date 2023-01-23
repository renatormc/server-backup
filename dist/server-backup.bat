@echo off
SET "PATH=C:\Windows\System32\OpenSSH;C:\Program Files\PostgreSQL\14\bin;%PATH%" && %~dp0\server-backup.exe %*