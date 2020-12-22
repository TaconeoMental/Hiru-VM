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
    @5    "print@io"
    @6    "ffi"
    @7    "Is even"
    @8    "Is odd"
SEGMENT .names
    @0    double
    @1    final
    @2    math
    @3    add
    @4    ffi
    @5    loadPrimitive
    @6    print
    @7    isEven
SEGMENT .code
    lconst @6 ("ffi")
    import
    makemod
    sname  @4 (ffi)

    lname  @4 (ffi)
    lattr  @5 (loadPrimitive)
    lconst @5 ("print@io")
    callfn 1  (ffi.loadPrimitive.call)
    sname  @6 (print)

    lname  @6 (print)
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

    lname  @6 (print)
    lname  @1 (final)
    callfn 1  (print.call)

    lname  @2 (math)
    lattr  @7 (isEven)
    lname  @1 (final)
    callfn 1  (math.isEven.call)
    
    (Si es par, saltamos)
    pjmpt  even
    lname  @6 (print)
    lconst @8 ("Is odd")
    callfn 1  (print.call)
    jmpabs end

even:
    lname  @6 (print)
    lconst @7 ("Is even")
    callfn 1  (print.call)

end:
    exit
    
