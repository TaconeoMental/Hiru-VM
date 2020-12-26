[HIRU]
SEGMENT .data
    @0    1
    @1    1343
    @2   "TRUE"
    @3   "FALSE"
    @4   [lessthan]
         SEGMENT .data
             @0    0
             @1    1
         SEGMENT .names
             @0    x
             @1    y
         SEGMENT .code
             lname  @0 (x)
             lname  @1 (y)
             cmplt
             pjmpt  truelab
             jmpabs falselab

         truelab:
             lconst @1
             ret

         falselab:
             lconst @0
             ret
    @5    0
    @6    1
SEGMENT .names
    @0    final
    @1    lessthan
SEGMENT .code
    lconst @4
    sname  @1
    lname  @1
    lconst @0
    lconst @1
    callfn 2
    lconst @5
    cmpeq
    pjmpt true
    jmpabs end

true:
    lconst @2
    sname  @0
    exit

end:
    lconst @3
    sname  @0
    exit


