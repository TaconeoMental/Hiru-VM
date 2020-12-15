[HIRU]
SEGMENT .data
    @0    "abcd"
    @1    [F1]
          SEGMENT .data
              @0    42
              @1    [F2]
                    SEGMENT .data
                        @0    [F3]
                              SEGMENT .data
                              SEGMENT .names
                              SEGMENT .code
                    SEGMENT .names
                    SEGMENT .code
          SEGMENT .names
          SEGMENT .code
SEGMENT .names
SEGMENT .code
