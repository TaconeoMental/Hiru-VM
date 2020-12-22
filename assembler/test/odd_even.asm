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
              lconst @0
              lname  @0
              bmod
              lconst @1
              cmpeq
              ret
    @2    4
    @3    2
SEGMENT .names
    @0    add
    @1    isEven
    @2    res
SEGMENT .code
    lconst @0    ([add])
    sname  @0    (add)
    lconst @1    ([isEven])
    sname  @1    (isEven)

    lname  @1    (isEven)
    lname  @0    (add)
    lconst @2    (4)
    lconst @3    (1)
    callfn 2     (add.call)
    callfn 1     (isEven.call)
    sname  @2    (res)
    lname  @2
    print
    exit
