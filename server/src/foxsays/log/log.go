package log

import stdLog "log"

var (
	Fatal     = stdLog.Fatal
	Fatalf    = stdLog.Fatalf
	Fatalln   = stdLog.Fatalln
	Panic     = stdLog.Panic
	Panicf    = stdLog.Panicf
	Panicln   = stdLog.Panicln
	Prefix    = stdLog.Prefix
	Print     = stdLog.Print
	Printf    = stdLog.Printf
	Println   = stdLog.Println
	SetOutput = stdLog.SetOutput
)

func FatalIf(err error) {
	if err != nil {
		stdLog.Fatal(err)
	}
}
