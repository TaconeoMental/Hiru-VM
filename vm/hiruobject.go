package vm

import (
        "strconv"
)

// Representa un obeto en Hiru
type HiruObject interface {
        Type() string
        Inspect() string
}


// Un número. Fui estúpido y solo implementé enteros, pero ya se vienen los
// flotantes.
type HiruNumber struct {
        Value int64
}

func (n *HiruNumber) Type() string {
        return "Integer"
}

func (n *HiruNumber) Inspect() string {
        return strconv.FormatInt(n.Value, 10)
}

// Valor booleano. También olvidé implementarlos...
type HiruBoolean struct {
	Value bool
}

func (b *HiruBoolean) Type() string {
	return "Boolean"
}

func (b *HiruBoolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

// Valor nulo
type HiruNull struct{}

func (n *HiruNull) Type() string {
	return "Null"
}

func (n *HiruNull) Inspect() string {
	return "Null"
}

// String en hiru
type HiruString struct {
        Value string
}

func (s *HiruString) Type() string {
        return "String"
}

func (s *HiruString) Inspect() string {
        return s.Value
}

// Función en Hiru
type HiruFunction struct {
        CodeObject *CodeObject
}

func (f *HiruFunction) Type() string {
        return "Function"
}

func (f *HiruFunction) Inspect() string {
        return "<Function>"
}

func (f *HiruFunction) RawObject() *CodeObject {
        return f.CodeObject
}

// Para el ffi
type HiruForeignFunction struct {
        // TODO
}

// Módulo en Hiru
type HiruModule struct {
        CodeObject *CodeObject
        StackFrame *StackFrame
}

func (f *HiruModule) Type() string {
        return "Module"
}

func (f *HiruModule) Inspect() string {
        return "<Module>"
}

func (f *HiruModule) RawObject() *CodeObject {
        return f.CodeObject
}

// Estructura en Hiru
type HiruStructure struct {
        CodeObject *CodeObject
        StackFrame *HiruStructureVars
}

func (s *HiruStructure) Type() string {
        return "Structure"
}

func (s HiruStructure) Inspect() string {
        return "<Struct>"
}

func (s *HiruStructure) RawObject() *CodeObject {
        return s.CodeObject
}

func (s *HiruStructure) Vars() *HiruStructureVars {
        return s.StackFrame
}

// Instancia de una estructura en Hiru
type HiruInstance struct {
        CodeObject *CodeObject
        StackFrame *HiruStructureVars
}

func (i *HiruInstance) Type() string {
        return "Instance"
}

func (i HiruInstance) Inspect() string {
        return "<Instance>"
}

func (i *HiruInstance) RawObject() *CodeObject {
        return i.CodeObject
}

// Vars representa el stackframe interno de una estructura
type HiruStructureVars struct {
        sf *StackFrame
}

// En teoría nunca se van a usar, pero es para que cumpla la interfaz de
// HiruOBject
func (v *HiruStructureVars) Type() string {
        return "Vars"
}

func (v HiruStructureVars) Inspect() string {
        return "<Vars>"
}

func (v *HiruStructureVars) StackFrame() *StackFrame {
        return v.sf
}
