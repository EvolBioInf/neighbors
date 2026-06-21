f=test.nwk
./pickle 7 $f   > r1.txt
./pickle 7,3 $f > r2.txt
./pickle 9 $f   > r3.txt
./pickle 4 $f   > r4.txt

./pickle -c 7 $f   > r5.txt
./pickle -c 7,3 $f > r6.txt
./pickle -c 9 $f   > r7.txt
./pickle -c 4 $f   > r8.txt

./pickle -t 7 $f   > r9.txt
./pickle -t 7,3 $f > r10.txt
./pickle -t 9 $f   > r11.txt
./pickle -t 4 $f   > r12.txt

./pickle -C 7 $f   > r13.txt
./pickle -C 7,3 $f > r14.txt
./pickle -C 9 $f   > r15.txt
./pickle -C 4 $f   > r16.txt

./pickle -c -t 7 $f   > r17.txt
./pickle -c -t 7,3 $f > r18.txt
./pickle -c -t 9 $f   > r19.txt
./pickle -c -t 4 $f   > r20.txt

