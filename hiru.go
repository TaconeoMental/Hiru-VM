package main

import (
        "fmt"
        "os"
        "flag"
        "log"

        "github.com/TaconeoMental/Hiru-VM/vm"
)

func main() {
        debug := flag.Bool("debug", false, "Run the VM in debug mode.")
        objectMode := flag.Bool("obi", false, "Object based interpretation. Parses the bytecode as objects and then interprets them (Default).")
        indexMode := flag.Bool("ibi", false, "Index based interpretation. Directly interprets the bytecode without any intermediate representation.")

        // Definimos nuestro propio usage del CLI
        flag.Usage = func() {
            fmt.Fprintf(os.Stderr, "Usage: %s [Options] [Filename]\n", os.Args[0])
            fmt.Fprintln(os.Stderr, "Options:")

            flag.VisitAll(func(f *flag.Flag) {
                fmt.Fprintf(os.Stderr, "    -%v,\t%v\n", f.Name, f.Usage) // f.Name, f.Value
            })
        }

        // Parseamos los argumentos
        flag.Parse()

        var mode vm.VmMode

        // Son mutuamente excluyentes
        if *objectMode && *indexMode {
                fmt.Fprintln(os.Stderr, "Error: '-obi' and '-ibi' can not be selected mutually.")
                flag.Usage()
                return
        } else if *objectMode {
                mode = vm.ObjectBasedMode
        } else if *indexMode {
                mode = vm.IndexBasedMode
        } else {

                mode = vm.ObjectBasedMode
        }

        // Si se entregó un path válido entonces proseguimos
        if filePath := flag.Arg(0); filePath != "" {

                options := vm.NewVmOptions(*debug, mode)

                vm, err := vm.NewVm(filePath, *options)
                if err != nil {
                        fmt.Println(err.Error())
                        return
                }

                if err := vm.Run(); err != nil {
                        log.Fatal("The VM has encountered an error: ", err.Error())
                        return
                }

        } else {
                flag.Usage()
                return
        }
}
