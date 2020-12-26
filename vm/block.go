package vm

import (
        "sync"
        "errors"
)

type Block struct {
        blockType   string // TODO: Reemplazar por un Enum
        handlerIp   int32
}

func NewBlock(btype string, handler int32) *Block {
        return &Block{blockType: btype, handlerIp: handler}
}

type BlockStack struct {
        blocks []*Block
        lock   sync.Mutex
}

func NewBlockStack() *BlockStack {
        return &BlockStack{
                blocks: make([]*Block, 0),
                lock: sync.Mutex{}}
}

func (bs *BlockStack) Push(b *Block) {
    bs.lock.Lock()
    defer bs.lock.Unlock()

    bs.blocks = append(bs.blocks, b)
}

func (bs *BlockStack) Pop() (*Block, error) {
    bs.lock.Lock()
    defer bs.lock.Unlock()


    l := len(bs.blocks)
    if l == 0 {
        return nil, errors.New("Empty BlockStack")
    }

    res := bs.blocks[l-1]
    bs.blocks = bs.blocks[:l-1]
    return res, nil
}

func (bs *BlockStack) GetTopMostType() string {
    l := len(bs.blocks)
    if l == 0 {
        return "Empty"
    }

    res := bs.blocks[l - 1]
    return res.blockType
}
