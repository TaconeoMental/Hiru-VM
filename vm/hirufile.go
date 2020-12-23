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

        // Esto es para ObjectBasedMode
        buffer *bytes.Reader

        indexPointer int

        // Este para IndexBasedMode
        //buffer []byte
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
        hf.indexPointer = 0
        return hf, nil
}

// Métodos

func (hf HiruFile) Buffer() *bytes.Reader {
        return hf.buffer
}

func (hf HiruFile) FileName() string {
        return filepath.Base(hf.path)
}

func (hf HiruFile) FullPath() string {
        return hf.path
}

func (hf *HiruFile) ReadBytes(length int) []byte {
        b := make([]byte, length)

        binary.Read(hf.buffer, binary.BigEndian, &b)

        hf.indexPointer += length
        return b
}

func (hf *HiruFile) ReadByte() byte {
        b := hf.ReadBytes(1)

        byteBuffer := bytes.NewReader(b)
        var byteVar byte
        binary.Read(byteBuffer, binary.BigEndian, &byteVar)
        return byteVar

        /*
        var b byte
        binary.Read(hf.buffer, binary.BigEndian, &b)

        hf.indexPointer += 1
        return b
        */
}

func (hf *HiruFile) Read4Bytes() int32 {
        return int32(binary.BigEndian.Uint32(hf.ReadBytes(4)))
        /*
        var w uint32
        binary.Read(hf.buffer, binary.BigEndian, &w)

        hf.indexPointer += 4
        return w
        */
}

func (hf *HiruFile) Seek(offset int) int {
        hf.buffer.Seek(0, offset)

        hf.indexPointer = offset
        return offset
}
