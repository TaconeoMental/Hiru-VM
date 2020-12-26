[HIRU]
SEGMENT .data
    @0    [Car]
          SEGMENT .data
              @0    [__new__]
                    SEGMENT .data
                    SEGMENT .names
                        @0    start
                        @1    max
                        @2    velocity
                    SEGMENT .code
                        lname  @0    (start)
                        lself
                        sattr  @2    (velocity)

                        lname  @1    (max)
                        lself
                        sattr  @1    (max)

                        lself
                        ret

              @1    [accelerate]
                    SEGMENT .data
                        @0    "Too fast!"
                        @1    -1
                    SEGMENT .names
                        @0    amount
                        @1    velocity
                        @2    max
                    SEGMENT .code
                        lname  @0    (amount)
                        lself        (self)
                        lattr  @1    (self.velocity)
                        badd         (amount + self.velocity)
                        lself        (self)
                        lattr  @2    (self.max)
                        cmplt        (amount + self.velocity < self.max)
                        pjmpf  toofast

                        lname  @0    (amount)
                        lself        (self)
                        lattr  @1    (self.velocity)
                        badd         (amount + self.velocity)
                        lself        (self)
                        sattr  @1    (self.velocity)

                        lname  @0    (amount)
                        ret

                    toofast:
                        lconst @0    ("Too fast!")
                        print
                        lconst @1    (-1)
                        ret
          SEGMENT .names
              @0    __new__
              @1    accelerate
          SEGMENT .code
              lconst @0    ([__new__])
              sname  @0    (__new__)

              lconst @1    ([accelerate])
              sname  @1    (accelerate)
              lvars
              ret
    @1    "Current speed: "
    @2    10
    @3    30
    @4    "Accelerating"
    @5    "You crashed :("
    @6    -1
SEGMENT .names
    @0    Car
    @1    new_car
    @2    velocity
    @3    accelerate
SEGMENT .code
    lconst @0    ([Car])
    callfn 0
    builds
    sname  @0    (new_car)

    lname  @0    (Car)
    lconst @2    (10)
    lconst @3    (30)
    inits  2     (Car{10, 30})
    sname  @1    (new_car)

    lconst @1    ("Current speed: ")
    print

    lname  @1    (new_car)
    lattr  @2    (new_car.velocity)
    print

    lconst @4    ("Accelerating")
    print

    lname  @1    (new_car)
    lattr  @3    (new_car.accelerate)
    lconst @2    (10)
    callfn 1     (new_car.accelerate{10})

    lconst @1    ("Current speed: ")
    print

    lname  @1    (new_car)
    lattr  @2    (new_car.velocity)
    print

    lconst @4    ("Accelerating")
    print

    lname  @1    (new_car)
    lattr  @3    (new_car.accelerate)
    lconst @2    (10)
    callfn 1     (new_car.accelerate{10})
    lconst @6    (-1)
    cmpeq        (new_car.accelerate{10} == -1)
    pjmpt  crash

    lconst @1    ("Current speed: ")
    print

    lname  @1    (new_car)
    lattr  @2    (new_car.velocity)
    print
    exit

crash:
    lconst @5    ("You crashed :(")
    print
    exit


