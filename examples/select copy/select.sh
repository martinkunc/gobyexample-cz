# Přijmeme hodnoty `"jedna"` a pak `"dvě"` jak jsme
# očekávali.
$ time go run select.go 
přijato jedna
přijato dvě

# Všimněte si, že celkový čas běhu je jenom ~2 sekundy
# protože obě 1 i 2 sekundové `Sleep` se provede
# souběžně.
real	0m2.245s
