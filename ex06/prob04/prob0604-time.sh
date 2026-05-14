echo "time GODEBUG=gctrace=1 ./prob0604"
time GODEBUG=gctrace=1 ./prob0604
echo ""
echo "\n\n================="
echo "$ time GOGC=50 GODEBUG=gctrace=1 ./prob0604"
time GOGC=50 GODEBUG=gctrace=1 ./prob0604
echo "\n\n================="
echo "$ time GOGC=200 GODEBUG=gctrace=1 ./prob0604"
time GOGC=200 GODEBUG=gctrace=1 ./prob0604
echo "\n\n================="
echo "$ time GOGC=1000 GODEBUG=gctrace=1 ./prob0604"
time GOGC=1000 GODEBUG=gctrace=1 ./prob0604
echo "\n\n================="
echo "$ time GOGC=off GODEBUG=gctrace=1 ./prob0604"
time GOGC=off GODEBUG=gctrace=1 ./prob0604
