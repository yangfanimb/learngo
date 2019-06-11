rd /s/q release
md release
go build -o blog.exe
COPY blog.exe release\
REM COPY favicon.ico release\favicon.ico
REM XCOPY config\*.* release\config\  /s /e
REM XCOPY mnt\*.* release\mnt\  /s /e
REM XCOPY asset\*.* release\asset\  /s /e
REM XCOPY view\*.* release\view\  /s /e
echo "package success, find app in release directory"
@pause
