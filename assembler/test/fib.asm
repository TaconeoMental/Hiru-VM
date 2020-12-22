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
              bsub
              callfn 1     (fib.call)
              lname  @1    (fib)
              lname  @0    (n)
              lconst @1    (2)
              bsub
              callfn 1
              badd
              ret
    @1    20
SEGMENT .names
    @0    fib
SEGMENT .code
    lconst @0    ([fib])
    sname  @0    (fib)
    lname  @0    (fib)
    lconst @1    (2)
    callfn 1     (fib.call)
    print
    exit
