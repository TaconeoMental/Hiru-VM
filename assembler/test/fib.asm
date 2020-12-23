[HIRU]
SEGMENT .data
    @0    [fib]
          SEGMENT .data
              @0    1
              @1    2
          SEGMENT .names
              @0    num
              @1    fib
          SEGMENT .code
              lname  @0    (num)
              lconst @0    (1)
              cmple
              pjmpf  recursion

              lname  @0    (num)
              ret

          recursion:
              lname  @1    (fib)
              lname  @0    (num)
              lconst @0    (1)
              bsub         (num - 1)
              callfn 1     (fib.call)

              lname  @1    (fib)
              lname  @0    (num)
              lconst @1    (2)
              bsub         (num - 2)
              callfn 1
              badd
              ret
    @1    [iterfib]
          SEGMENT .data
              @0    -1
              @1    0
              @2    1
          SEGMENT .names
              @0    num    (Parameter)
              @1    prev
              @2    curr
              @3    elem
              @4    count
          SEGMENT .code
              lconst @0    (-1)
              sname  @1    (prev)

              lconst @2    (1)
              sname  @2    (curr)

              lconst @1    (0)
              sname  @4    (count)

              lconst @1    (0)
              sname  @3    (elem)
              sloop  endfib
          startfib:
              lname  @4    (count)
              lname  @0    (num)
              cmple        (count < num)
              pjmpf  endfib

              lname  @1    (prev)
              lname  @2    (curr)
              badd         (prev + curr)
              sname  @3    (elem)

              lname  @2    (curr)
              sname  @1    (prev)

              lname  @3    (elem)
              sname  @2    (curr)

              lconst @2
              lname  @4    (count)
              badd         (1 + count)
              sname  @4    (count)
              jmpabs startfib

          endfib:
              ploop

          endloop:
              lname  @3    (elem)
              ret
            
    @2    25
    @3    "Iterative fibonacci:"
    @4    "Recursive fibonacci:"
SEGMENT .names
    @0    fib
    @1    iterfib
SEGMENT .code
    lconst @0    ([fib])
    sname  @0    (fib)

    lconst @1    ([iterfib])
    sname  @1    (iterfib)

    lconst @3
    print

    lname  @1    (iterfib)
    lconst @2    (25)
    callfn 1     (iterfib.call)
    print

    lconst @4
    print

    lname  @0    (fib)
    lconst @2    (25)
    callfn 1     (fib.call)
    print
    exit
