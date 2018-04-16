## __A simple bank management system written on Go__

# **Install Instruction** #


First of all you need Go and node.js/npm installed

Please install all necessary dependencies

```bash
npm install --save-dev
```

When it will be done you must add config file there will be your secure settings( database connection, secret salt etc.) from example *example.json*

```bash
cp example.json config.json
```

When you can change all needed env variables(set passwords for example)

```bash
nano config.json
```

Also you need to install all dependencies using [Govendor](https://github.com/kardianos/govendor) tool

```bash
go get -u github.com/kardianos/govendor
govendor sync
```

Than you need to build binary file
```bash
go build -o output_binary .
```

And run binary file

```bash
./output_binary
```
