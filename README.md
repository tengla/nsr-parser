# NSR XML Parser
Parse NSR records with Golang

```
make all
```

There's a zip'ed xml blob called nsr.current.xml.gz
Unzip it!

```
gunzip nsr.current.xml.gz
./dist/import -file ./nsr.current.xml
```
