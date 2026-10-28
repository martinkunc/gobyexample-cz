# Pro spuštění příkladu použijte tyto příkazy.
# (Pozn.: kvůli omezení go playground je možné
# tento příklad spustit jen na vašem počítači.)
$ mkdir -p folder
$ echo "hello go" > folder/single_file.txt
$ echo "123" > folder/file1.hash
$ echo "456" > folder/file2.hash

$ go run embed-directive.go
hello go
hello go
123
456
