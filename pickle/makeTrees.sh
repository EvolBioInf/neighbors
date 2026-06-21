d="3.75cm,3.75cm"
./pickle -t 9 test.nwk |
    plotTree -p tree1.ps -d $d
./pickle -t -c 8 test.nwk |
    plotTree -p tree2.ps -d $d
./pickle -C 8 test.nwk |
    plotTree -p tree3.ps -d $d
