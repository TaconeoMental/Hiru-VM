[HIRU]
SEGMENT .data
    @0    "io"
    @1    [concat]
          SEGMENT .data
          SEGMENT .names
              @0    s1
              @1    s2
          SEGMENT .code
              lname   @0
              lname   @1
              badd
              ret
    @2    [main]
          SEGMENT .data
              @0    "¡Hola"
              @1    ", "
              @2    "Mundo!"
          SEGMENT .names
              @0    a
              @1    b
              @2    c
              @3    concat
              @4    io
              @5    println
          SEGMENT .code
              lconst    @0
              sname     @0
              lconst    @2
              sname     @1
              lname     @3
              lname     @3
              lname     @0
              lconst    @1
              callfn    2
              lname     @1
              callfn    2
              sname     @2
              lname     @4
              lattr     @5
              lname     @2
              callfn    1
              ret
SEGMENT .names
    @0    concat
    @1    main
SEGMENT .code
    lconst     @0
    makemod
    lconst     @1
    makefn
    sname      @0
    lconst     @2
    makefn
    sname      @1
    lname      @1
    callfn     0
