[HIRU]
SEGMENT .data
    @0    [Vector]
          SEGMENT .data
              @0    [abs]
                    SEGMENT .data
                        @0    2
                    SEGMENT .names
                        @0    x
                        @1    y
                        @2    math
                        @3    sqrt
                    SEGMENT .code
                        lself
                        lattr  @0    (x)
                        lconst @0    (2)
                        bpow

                        lself
                        lattr  @1    (y)
                        lconst @0    (2)
                        bpow
                        badd         (x^2 + y^2)
                        ret

              @1    [__new__]
                    SEGMENT .data
                        @0    2
                    SEGMENT .names
                        @0    x
                        @1    y
                        @2    nombre
                    SEGMENT .code
                        lconst @0
                        lname  @0
                        bmul
                        sname  @0

                        lname  @0
                        lself
                        sattr  @0

                        lname  @1
                        lself
                        sattr  @1

                        lname  @2
                        lself
                        sattr  @2

                        lself
                        ret
              @2    2
          SEGMENT .names
              @0    __new__
              @1    abs
          SEGMENT .code
              lconst @1    ([__new__])
              sname  @0    (__new__)

              lconst @0    ([abs])
              sname  @1    (abs)
              lvars
              ret
    @1    2
    @2    7
    @3    10
SEGMENT .names
    @0    Vector
    @1    instance
    @2    x
    @3    y
    @4    abs
    @5    math
SEGMENT .code
    lconst @0    ([Vector])
    callfn 0
    builds
    sname  @0    (Vector)

    lname  @0    (Vector)
    lconst @1    (2)
    lconst @2    (7)
    inits  2
    sname  @1    (instance)

    lname  @1    (instance)
    lattr  @2    (x)
    print

    lname  @1    (instance)
    lattr  @4    (abs)
    callfn 0
    print

    exit


