$ go build command-line-subcommands.go 

# Nejdřív zavoláme podpříkaz foo.
$ ./command-line-subcommands foo -enable -name=joe a1 a2
subcommand 'foo'
  enable: true
  name: joe
  tail: [a1 a2]

# Teď zkusíme bar.
$ ./command-line-subcommands bar -level 8 a1
subcommand 'bar'
  level: 8
  tail: [a1]

# Ale bar neakceptuje volby foo.
$ ./command-line-subcommands bar -enable a1
flag provided but not defined: -enable
Usage of bar:
  -level int
    	level

# Dále se podíváme na proměnné prostředí, další běžný
# způsob jak parametrizovat programy.
