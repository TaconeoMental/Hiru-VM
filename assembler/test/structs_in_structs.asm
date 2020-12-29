[HIRU]
SEGMENT .data
    @0    [Vector]
          SEGMENT .data
              @0    [__new__]
                    SEGMENT .data
                    SEGMENT .names
                        @0    pos
                        @1    name
                        @2    String
                    SEGMENT .code
                        lname  @0    (pos)
                        lself
                        sattr  @0    (self.pos)

                        lname  @2    (String)
                        lname  @1    (name)
                        inits  1
                        lself
                        sattr  @1    (self.name)

                        lself
                        ret
          SEGMENT .names
              @0    __new__
          SEGMENT .code
              lconst @0    ([__new__])
              sname  @0    (__new__)
              lvars
              ret

    @1    [Point]
          SEGMENT .data
              @0    [__new__]
                    SEGMENT .data
                    SEGMENT .names
                        @0    x
                        @1    y
                    SEGMENT .code
                        lname  @0    (x)
                        lself
                        sattr  @0    (self.x)

                        lname  @1    (y)
                        lself
                        sattr  @1    (self.y)

                        lself
                        ret
          SEGMENT .names
              @0    __new__
          SEGMENT .code
              lconst @0    ([__new__])
              sname  @0    (__new__)
              lvars
              ret
    @2    [String]
          SEGMENT .data
              @0    [__new__]
                    SEGMENT .data
                    SEGMENT .names
                        @0    Value
                    SEGMENT .code
                        lname  @0
                        lself
                        sattr  @0

                        lself
                        ret
          SEGMENT .names
              @0    __new__
          SEGMENT .code
              lconst @0    ([__new__])
              sname  @0    (__new__)
              lvars
              ret
    @3    4
    @4    10
    @5    "mario"
SEGMENT .names
    @0    Vector
    @1    Point
    @2    point
    @3    vec
    @4    pos
    @5    x
    @6    String
    @7    name
    @8    Value
SEGMENT .code
    lconst @0    ([Vector])
    callfn 0
    builds
    sname  @0    (Vector)

    lconst @1    ([Point])
    callfn 0
    builds
    sname  @1    (Point)

    lconst @2
    callfn 0
    builds
    sname  @6

    lname  @1    (Point)
    lconst @3    (4)
    lconst @4    (10)
    inits  2
    sname  @2    (point)

    lname  @0    (Vector)
    lname  @2    (point)
    lconst @5    ("mario")
    inits  2
    sname  @3    (vec)

    lconst @3
    lname  @3    (vec)
    lattr  @4    (pos)
    sattr  @5    (x)

    lname  @3
    lattr  @4
    lattr  @5
    print

    lname  @3
    lattr  @7
    lattr  @8
    print
 
    exit
