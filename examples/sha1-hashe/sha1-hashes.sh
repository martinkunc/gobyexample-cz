# Spuštění programu vypočítá hash a vypíše ho ve
# lidsky čitelném hexa formátu.
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94


# S použitím podobného vzoru můžeš vypočítat jiné hashe
# podobné tomu výše. Například pro výpočet MD5 hashe
# importuj `crypto/md5` a použij `md5.New()`.

# Poznamenej si, že pokud budeš potřebovat kryptograficky
#  bezpečné hashe, měl bys opatrně prozkoumat 
# [sílu hash funkcí](http://en.wikipedia.org/wiki/Cryptographic_hash_function)!
