package vm

import (
        "bytes"
        "os"
        "encoding/binary"
        "io/ioutil"
        "path/filepath"
        "log"
)

// Funciones auxiliares

func expandPath(path string) (string, error) {
        cdPath, err := os.Getwd()
        if err != nil {
                log.Fatal("Cannot get current working directory")
                return "", err
        }
        expandedPath := filepath.Join(cdPath, path)
        return expandedPath, nil
}

type HiruFile struct {
        path   string
        buffer *bytes.Reader
}

func NewHiruFile(path string) (*HiruFile, error) {
        hf := new(HiruFile)
        path, err := expandPath(path)
        if err != nil {
                return nil, err
        }
        hf.path = path

        content, err := ioutil.ReadFile(path)
        if err != nil {
                return nil, err
        }

        hf.buffer = bytes.NewReader(content)
        return hf, nil
}

// Métodos

func (hf HiruFile) FileName() string {
        return filepath.Base(hf.path)
}

func (hf HiruFile) FullPath() string {
        return hf.path
}

func (hf *HiruFile) ReadByte() (byte, error) {
        var b byte
        err := binary.Read(hf.buffer, binary.LittleEndian, &b)
        return b, err
}

func (hf *HiruFile) ReadWord() (uint16, error) {
        var w uint16
        err := binary.Read(hf.buffer, binary.LittleEndian, &w)
        return w, err
}
