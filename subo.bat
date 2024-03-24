git add .
git commit -m "Ultimo Commit"
git push
go build -o bootstrap main.go
del main.zip
tar.exe -a -cf main.zip bootstrap