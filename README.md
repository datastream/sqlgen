sqlgen
======

sql generator

sqlgen will not check sql syntax.

    q := sqlgen.Select("a","n")
    q.From("test")
    q.Where("a = ?", 1) // ok
    q.Where("a = ?") // ok, but it may fail
