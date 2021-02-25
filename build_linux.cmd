del /Q moeobject
del /Q moeobject_linux.tgz
set GOOS=linux
go build
tar czf moeobject_linux.tgz moeobject moeobject.toml.example
del /Q moeobject
