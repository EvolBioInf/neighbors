set terminal postscript monochrome size 20cm,7cm
set output "bl.ps"
set object rectangle from 1.4575e-5,-0.04 to 7.38e-5,0.04
set label "q1" at 1.4575e-5,-0.06 center
set label "q2" at 7.38e-5,-0.06 center
set label "rr" at 4.41875e-05,0.065 center
set arrow from 3.41875e-5,0.065,1.4575e-5 to 1.4575e-5,0.065
set arrow from 5.41875e-05,0.065 to 7.38e-5,0.065
set arrow from -0.000074,-0.03 to -0.000074,0.03 nohead
set arrow from 0.00016,-0.03 to 0.00016,0.03 nohead
set arrow from -0.00016,-0.03 to -0.00016,0.03 nohead
set arrow from 0.00025,-0.03 to 0.00025,0.03 nohead
set arrow from -0.00016,0 to 0.00025,0 nohead
unset ytics
plot [-0.00018:*][-0.25:0.25] "bl.dat" u 1:(0.06*rand(0)-0.03) title "descendants" w p, \
"-" t "parent" w p pt 7
0.000385 0