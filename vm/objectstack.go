package vm

import (
        "sync"
        "errors"
)

type ObjectStack struct {
        stack []HiruObject
        lock  sync.Mutex
}

func NewObjectStack() *ObjectStack {
        return &ObjectStack{
                stack: make([]HiruObject, 0),
                lock: sync.Mutex{}}
}

func (objs *ObjectStack) Push(obj HiruObject) {
    objs.lock.Lock()
    defer objs.lock.Unlock()

    objs.stack = append(objs.stack, obj)
}

func (objs *ObjectStack) Pop() (HiruObject, error) {
    objs.lock.Lock()
    defer objs.lock.Unlock()


    l := len(objs.stack)
    if l == 0 {
        return nil, errors.New("Empty ObjectStack")
    }

    res := objs.stack[l-1]
    objs.stack = objs.stack[:l-1]
    return res, nil
}


