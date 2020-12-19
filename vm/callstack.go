package vm

import (
        "sync"
        "errors"
)

type CallStack struct {
        records         []*StackFrame
        lock            sync.Mutex
}

func NewCallStack() *CallStack {
        return &CallStack{
                records: make([]*StackFrame, 0),
                lock: sync.Mutex{}}
}

// https://stackoverflow.com/a/28542256
func (cs *CallStack) Push(sf *StackFrame) {
    cs.lock.Lock()
    defer cs.lock.Unlock()

    cs.records = append(cs.records, sf)
}

func (cs *CallStack) Pop() (*StackFrame, error) {
    cs.lock.Lock()
    defer cs.lock.Unlock()


    l := len(cs.records)
    if l == 0 {
        return nil, errors.New("Empty CallStack")
    }

    res := cs.records[l-1]
    cs.records = cs.records[:l-1]
    return res, nil
}

func (cs *CallStack) ResolveName(name string) (*HiruObject, error) {
        value, err := cs.GetTopMost().ResolveName(name)
        if err != nil {
                current_global := cs.GetTopMost().LinkedTo()
                for current_global != nil {
                        if current_global.LinkedTo() != nil {
                                current_global = current_global.LinkedTo()
                        } else {
                                return current_global.ResolveName(name)
                        }
                }
        }
        return value, nil
}

func (cs *CallStack) GetTopMost() *StackFrame {
        return cs.records[len(cs.records)-1]
}

func (cs *CallStack) Define(name string, obj *HiruObject) *HiruObject {
        cs.GetTopMost().Define(name, obj)
        return obj
}

type StackFrame struct {
        name   string
        parent *StackFrame
        enviroment map[string]*HiruObject
}

func NewStackFrame(name string) *StackFrame {
        sf := StackFrame{name: name}
        sf.enviroment = make(map[string]*HiruObject)
        return &sf
}

func (sf StackFrame) GetName() string {
        return sf.name
}

func (sf *StackFrame) MakeLinkTo(parent *StackFrame) {
        sf.parent = parent
}

func (sf StackFrame) LinkedTo() *StackFrame {
        return sf.parent
}

func (sf *StackFrame) Define(name string, obj *HiruObject) {
        sf.enviroment[name] = obj
}

func (sf *StackFrame) ResolveName(name string) (*HiruObject, error) {
        value, err := sf.GetLocalName(name)
        if err != nil {
                if sf.parent != nil {
                        value, err = sf.LinkedTo().GetLocalName(name)
                } else {
                        return nil, err
            }
        }
        return value, err
}

func (sf *StackFrame) GetLocalName(name string) (*HiruObject, error) {
        if obj, ok := sf.enviroment[name]; ok {
                return obj, nil
        }
        return nil, errors.New("Local value not found")
}
