[HIRU]
SEGMENT .data
    @0    [double]
          SEGMENT .data
              @0    2 
          SEGMENT .names
              @0    num
          SEGMENT .code
              lname  @0 (num)
              lconst @0 (2)
              bmul
              ret
    @1    100
    @2    5
    @3    200
    @4    "math"
SEGMENT .names
    @0    double
    @1    final
    @2    math
    @3    add
SEGMENT .code
    lconst @4 ("math")
    import
    makemod   (Necessary for IBI)
    sname  @2 (math)

    lname  @2 (math)
    lattr  @3 (math.add1)

    lconst @0 ([double])
    makefn    (Necessary for IBI)
    sname  @0 (double)
    lname  @0 (double)
    lconst @1 (3)
    callfn 1  (double.call)

    lconst @3 (2)
    callfn 2  (math.add.call)

    sname  @1 (final)
    exit
    
