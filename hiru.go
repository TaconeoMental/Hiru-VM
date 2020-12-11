package main

import (
        "fmt"
        "os"
        "flag"
        "path/filepath"
)

func main() {
        // Definimos nuestro propio usage del CLI
        flag.Usage = func() {
            fmt.Fprintf(os.Stderr, "Usage: %s [Options] [Filename]\n", os.Args[0])
            fmt.Fprintln(os.Stderr, "Options:")

            flag.VisitAll(func(f *flag.Flag) {
                fmt.Fprintf(os.Stderr, "    -%v,\t%v\n", f.Name, f.Usage) // f.Name, f.Value
            })
        }

        debug := flag.Bool("debug", false, "Run the VM in debug mode.")
        flag.Parse()

        // Si se entregó un path válido entonces proseguimos
        if filePath := flag.Arg(0); filePath != "" {

                // Expandimos el path para que sea absoluto
                cwdPath, err := os.Getwd()
                if err != nil {
                        fmt.Fprintln(os.Stderr, "Cannot get current working directory")
                        os.Exit(1)
                }
                expandedPath := filepath.Join(cwdPath, filePath)

                if *debug {
                        fmt.Printf("CLI OPTIONS file: '%s' debug: '%t'\n", expandedPath, *debug)
                }

                os.Exit(0)
        } else {
                flag.Usage()
                os.Exit(1)
        }
}
