# NSR XML Parser
Parse NSR records into Golang

```
make
# Then run
./dist/parser -help
```

There's a zip'ed xml blob called nsr.current.xml.gz
Unzip it!

```
gunzip nsr.current.xml.gz
./dist/parser -xmlfile ./nsr.current.xml
```
