[HIRU]
SEGMENT .data
    @0    [add]
          SEGMENT .data
          SEGMENT .names
              @0    x
              @1    y
          SEGMENT .code
              lname  @0
              lname  @1
              badd
              ret

    @1    [isEven]
          SEGMENT .data
              @0    2
              @1    0
          SEGMENT .names
              @0    num
          SEGMENT .code
              lname  @0
              lconst @0
              bmod
              lconst @0
              cmpeq
              ret
SEGMENT .names
    @0    add
    @1    isEven
SEGMENT .code
    lconst @0
    sname  @0
    lconst @1
    sname  @1
    ret
