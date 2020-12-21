[HIRU]
SEGMENT .data
    @0    [concat]
          SEGMENT .data
          SEGMENT .names
            @0    s1 ("wena")
            @1    s2
          SEGMENT .code
            lname    @0
            lname    @1
            badd
            ret
    @1  "Hello"
    @2  "World"
SEGMENT .names
    @0  concat
SEGMENT .code
    lconst    @0
    makefn
    sname     @0
    lname     @0
    lconst    @1
    lconst    @2
    callfn    2
    ret
