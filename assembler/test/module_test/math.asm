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
SEGMENT .names
    @0    add
SEGMENT .code
    lconst @0
    sname  @0
    ret
