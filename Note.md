

准备工程：
    go mod init GoHotFix

plugin调用耗时：
    goos: linux
    goarch: amd64
    pkg: GoHotFix/Server
    cpu: AMD EPYC 7763 64-Core Processor                
    BenchmarkCallHotFix-2           396380382                2.845 ns/op
    BenchmarkCallLocal-2            1000000000               0.3389 ns/op
    BenchmarkCallHotFixNoParam-2    480988370                2.600 ns/op
    BenchmarkCallLocalNoParam-2     1000000000               0.3219 ns/op

    1. plugin调用性能差异在8倍多
    2. 带参数和不带参数差异不大

查看so导出
    objdump -tT hotfix.so
    nm -D hotfix.so
