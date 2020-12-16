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

        // Si se entregó un path válido entonces proseguimos
        if filePath := flag.Arg(0); filePath != "" {

                // reemplazar por una función log nueva que tome como argumento
                // a *debug
                if *debug {
                        fmt.Printf("CLI OPTIONS file: '%s' debug: '%t'\n", filePath, *debug)
                }

                vm, err := vm.NewVm(filePath, *debug)
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
