package main

import (
	_ "github.com/sub-rat/goclass_morning/example2_array"
	_ "github.com/sub-rat/goclass_morning/example_error_handling"
	_ "github.com/sub-rat/goclass_morning/example_time_defer_sort"
	_ "github.com/sub-rat/goclass_morning/examplefilehandling"
	_ "github.com/sub-rat/goclass_morning/goroutinesexample"

	_ "github.com/sub-rat/goclass_morning/ex_testing_and_benchmark"
	_ "github.com/sub-rat/goclass_morning/example_string_functions"
	httpclientserver "github.com/sub-rat/goclass_morning/http_client_server"
	_ "github.com/sub-rat/goclass_morning/json_example"
	_ "github.com/sub-rat/goclass_morning/number_parsing"
)

func main() {
	// initials.Init()
	// loopexample.InitLoopExample()
	// examplefunctions.InitSwap()
	// exapmle2array.Init()
	// exampledatastructure.InitDataStructure()
	// examplefunctions.InitVariadicFunction()
	// examplestructs.InitExampleStruct()
	// examplestructs.InitExampleStruct2()

	// examplefunctions.InitExampleClosures()

	// exampleerrorhandling.InitExampleErrorHandling()
	// goroutinesexample.InitGorountineExample()
	// goroutinesexample.InitChannelExample2()
	// exampletime.InitExampleTime()
	// example_time_defer_sort.InitExampleSort()
	// example_time_defer_sort.InitDeferExample()
	// example_string_functions.InitExampleStringFunction()
	// json_example.InitJsonExample()
	// number_parsing.InitNumberParsing()
	// examplefilehandling.InitFileHandling("example.txt")
	// examplefilehandling.InitFileWrite()

	// args := os.Args
	// argsWithoutName := args[1:]
	// fmt.Println(args)
	// fmt.Println(argsWithoutName)

	// extestingandbenchmark.InitTestingAndBenchmark()
	// httpclientserver.InitHttpClient()
	httpclientserver.InitHttpServer()

}
