[HIRU]
SEGMENT .data
    @0    [loadPrimitive]
          SEGMENT .data
          SEGMENT .names
              @0    libName
              @1    funcName
              @2    lib
              @3    loadFunc
          SEGMENT .code
              lname  @0
              gollib
              sname  @2
              lname  @2
              lattr  @3
              lname  @1
              callfn 1
              ret
SEGMENT .names
    @0    loadPrimitive
SEGMENT .code
    lconst @0
    sname  @0
    exit
    
