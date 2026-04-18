set terminal postscript monochrome size 20cm,7cm
set output "bl.ps"
set object rectangle from 1.4575e-5,-0.04 to 7.38e-5,0.04 lc "red"
set arrow from -0.00016,-0.03 to -0.00016,0.03 nohead
set arrow from 0.00025,-0.03 to 0.00025,0.03 nohead
set arrow from -0.00016,0 to 0.00025,0 nohead
unset ytics
plot [-0.00018:*][-0.25:0.25] "bl.dat" u 1:(0.06*rand(0)-0.03) title "descendants" w p, \
"-" t "parent" w p pt 7
0.000385 0