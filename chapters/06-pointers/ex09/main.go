package main

import "runtime"

func main() {
	// Ejecutar, por ejemplo, con GOGC=25 o GOGC=200 y gctrace=1.
	for etapa := 0; etapa < 10; etapa++ {
		data := make([][]byte, 0, 10000)
		for i := 0; i < 10000; i++ {
			data = append(data, make([]byte, 1024))
		}
		runtime.KeepAlive(data)
	}
}

// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09% GODEBUG=gctrace=1 GOGC=25 go run .
// gc 1 @0.001s 11%: 0.11+0.29+0.20 ms clock, 1.7+0.057/0.40/0+3.3 ms cpu, 1->1->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 2 @0.007s 3%: 0.010+0.19+0.003 ms clock, 0.17+0.039/0.46/0.046+0.054 ms cpu, 1->1->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 3 @0.009s 4%: 0.10+0.26+0.020 ms clock, 1.6+0.23/0.53/0+0.32 ms cpu, 1->2->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 4 @0.009s 4%: 0.079+0.27+0.068 ms clock, 1.2+0.33/0.45/0.006+1.0 ms cpu, 1->2->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 5 @0.010s 5%: 0.068+0.33+0.044 ms clock, 1.0+0.19/0.53/0+0.71 ms cpu, 1->2->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 6 @0.011s 6%: 0.053+0.18+0.014 ms clock, 0.86+0.050/0.50/0.011+0.22 ms cpu, 1->1->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 7 @0.012s 6%: 0.047+0.21+0.003 ms clock, 0.76+0.33/0.59/0.22+0.060 ms cpu, 1->2->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 8 @0.017s 5%: 0.042+0.28+0.012 ms clock, 0.68+0.20/0.60/0.001+0.20 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 9 @0.017s 5%: 0.13+0.32+0.003 ms clock, 2.0+0.19/0.62/0.031+0.054 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 10 @0.018s 6%: 0.062+0.23+0.003 ms clock, 0.99+0.42/0.42/0.004+0.052 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 11 @0.019s 6%: 0.015+0.23+0.012 ms clock, 0.24+0.027/0.67/0.15+0.19 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 12 @0.020s 6%: 0.035+0.20+0.003 ms clock, 0.57+0.20/0.60/0.001+0.053 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 13 @0.020s 6%: 0.035+0.19+0.003 ms clock, 0.56+0.18/0.56/0.043+0.058 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 14 @0.021s 6%: 0.019+0.18+0.018 ms clock, 0.30+0.13/0.57/0.023+0.29 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 15 @0.021s 6%: 0.030+0.25+0.003 ms clock, 0.49+0.032/0.58/0.16+0.062 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 16 @0.022s 6%: 0.011+0.23+0.004 ms clock, 0.19+0.025/0.55/0.25+0.065 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 17 @0.023s 6%: 0.010+0.20+0.003 ms clock, 0.16+0.018/0.54/0.14+0.048 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 18 @0.024s 6%: 0.011+0.15+0.002 ms clock, 0.17+0.059/0.43/0.13+0.041 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 19 @0.025s 6%: 0.014+0.22+0.003 ms clock, 0.22+0/0.52/0.050+0.056 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 20 @0.026s 6%: 0.009+0.19+0.002 ms clock, 0.15+0.10/0.50/0.090+0.039 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 21 @0.027s 6%: 0.011+0.18+0.003 ms clock, 0.18+0.074/0.46/0.25+0.048 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// # learning-go/chapters/06-pointers/ex09
// gc 1 @0.001s 6%: 0.093+0.45+0.017 ms clock, 1.4+0.21/1.1/0.12+0.28 ms cpu, 2->3->2 MB, 2 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 2 @0.002s 5%: 0.010+0.23+0.002 ms clock, 0.16+0.043/0.68/0.35+0.046 ms cpu, 3->3->2 MB, 3 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 3 @0.004s 4%: 0.011+0.48+0.021 ms clock, 0.18+0.071/1.0/0.55+0.34 ms cpu, 3->4->3 MB, 3 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 4 @0.005s 6%: 0.011+0.71+0.011 ms clock, 0.18+0.36/1.7/1.4+0.18 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 22 @0.040s 4%: 0.026+0.22+0.004 ms clock, 0.42+0.094/0.59/0.20+0.072 ms cpu, 2->2->1 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// # learning-go/chapters/06-pointers/ex09
// gc 1 @0.000s 6%: 0.089+0.23+0.031 ms clock, 1.4+0/0.47/0.072+0.49 ms cpu, 1->1->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 2 @0.002s 5%: 0.013+0.15+0.012 ms clock, 0.21+0.020/0.36/0.041+0.20 ms cpu, 1->1->1 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 3 @0.002s 5%: 0.012+0.16+0.016 ms clock, 0.19+0/0.41/0.12+0.26 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 4 @0.003s 5%: 0.010+0.41+0.022 ms clock, 0.17+0.021/0.92/0.29+0.36 ms cpu, 3->4->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 5 @0.004s 5%: 0.010+0.24+0.012 ms clock, 0.16+0/0.56/0.25+0.20 ms cpu, 4->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 6 @0.005s 5%: 0.010+0.19+0.011 ms clock, 0.17+0.016/0.51/0.33+0.18 ms cpu, 6->6->5 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 7 @0.007s 5%: 0.010+0.27+0.014 ms clock, 0.16+0.10/0.70/0.46+0.23 ms cpu, 7->7->6 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 8 @0.007s 5%: 0.011+0.25+0.019 ms clock, 0.18+0/0.67/0.39+0.31 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 9 @0.013s 3%: 0.014+0.40+0.013 ms clock, 0.22+0/1.0/0.64+0.21 ms cpu, 10->10->8 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 10 @0.017s 3%: 0.018+0.33+0.012 ms clock, 0.28+0/0.85/0.58+0.20 ms cpu, 9->9->7 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 11 @0.024s 2%: 0.015+0.37+0.012 ms clock, 0.25+0.019/1.1/0.73+0.20 ms cpu, 9->9->8 MB, 9 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 1 @0.000s 9%: 0.073+0.18+0.023 ms clock, 1.1+0/0.34/0.006+0.38 ms cpu, 0->1->1 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 2 @0.000s 8%: 0.012+0.12+0.013 ms clock, 0.19+0.012/0.16/0+0.21 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 3 @0.001s 7%: 0.016+0.14+0.002 ms clock, 0.27+0.033/0.21/0.014+0.044 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 4 @0.001s 7%: 0.013+0.18+0.003 ms clock, 0.21+0.028/0.25/0.017+0.053 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 5 @0.001s 7%: 0.008+0.17+0.003 ms clock, 0.14+0.013/0.24/0.010+0.051 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 6 @0.002s 7%: 0.010+0.27+0.003 ms clock, 0.16+0.075/0.45/0.052+0.054 ms cpu, 6->6->6 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 7 @0.002s 7%: 0.009+0.23+0.003 ms clock, 0.14+0.022/0.28/0.026+0.054 ms cpu, 7->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 8 @0.002s 7%: 0.008+0.24+0.002 ms clock, 0.13+0.015/0.33/0.009+0.040 ms cpu, 10->10->10 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 9 @0.003s 7%: 0.010+0.16+0.003 ms clock, 0.16+0.015/0.22/0.011+0.054 ms cpu, 12->12->2 MB, 12 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 10 @0.004s 6%: 0.010+0.16+0.003 ms clock, 0.16+0.016/0.22/0.007+0.052 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 11 @0.004s 6%: 0.010+0.17+0.003 ms clock, 0.16+0.029/0.26/0.020+0.049 ms cpu, 4->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 12 @0.005s 6%: 0.010+0.25+0.004 ms clock, 0.16+0.013/0.29/0.006+0.075 ms cpu, 6->6->6 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 13 @0.005s 6%: 0.008+0.43+0.012 ms clock, 0.13+0.031/0.39/0.031+0.20 ms cpu, 7->9->9 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 14 @0.006s 6%: 0.012+0.14+0.003 ms clock, 0.20+0.031/0.22/0.020+0.057 ms cpu, 10->11->1 MB, 11 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 15 @0.006s 6%: 0.011+0.15+0.003 ms clock, 0.18+0.009/0.19/0.007+0.048 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 16 @0.007s 6%: 0.012+0.15+0.016 ms clock, 0.20+0.011/0.18/0+0.27 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 17 @0.007s 5%: 0.008+0.17+0.002 ms clock, 0.14+0.014/0.24/0.010+0.045 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 18 @0.008s 5%: 0.008+0.20+0.003 ms clock, 0.13+0.024/0.25/0.015+0.057 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 19 @0.008s 5%: 0.010+0.17+0.002 ms clock, 0.17+0.017/0.22/0.010+0.042 ms cpu, 6->7->7 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 20 @0.008s 5%: 0.009+0.20+0.002 ms clock, 0.15+0.018/0.25/0.022+0.040 ms cpu, 8->9->9 MB, 9 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 21 @0.009s 5%: 0.009+0.14+0.003 ms clock, 0.15+0.023/0.19/0.011+0.057 ms cpu, 10->11->1 MB, 11 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 22 @0.009s 5%: 0.010+0.10+0.003 ms clock, 0.16+0.009/0.14/0.005+0.048 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 23 @0.010s 5%: 0.009+0.16+0.011 ms clock, 0.15+0.016/0.20/0+0.19 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 24 @0.010s 6%: 0.010+0.15+0.009 ms clock, 0.16+0.010/0.20/0.004+0.15 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 25 @0.010s 6%: 0.009+0.19+0.010 ms clock, 0.14+0.018/0.27/0.020+0.16 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 26 @0.011s 6%: 0.015+0.20+0.003 ms clock, 0.24+0.022/0.23/0.008+0.050 ms cpu, 6->7->7 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 27 @0.011s 6%: 0.008+0.23+0.002 ms clock, 0.14+0.034/0.31/0.046+0.046 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 28 @0.011s 6%: 0.010+0.13+0.003 ms clock, 0.16+0.018/0.18/0.013+0.048 ms cpu, 10->11->1 MB, 11 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 29 @0.012s 6%: 0.010+0.13+0.003 ms clock, 0.16+0.011/0.15/0+0.049 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 30 @0.012s 6%: 0.008+0.16+0.011 ms clock, 0.13+0.010/0.21/0+0.18 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 31 @0.012s 6%: 0.008+0.11+0.002 ms clock, 0.14+0.009/0.16/0.006+0.036 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 32 @0.013s 6%: 0.009+0.14+0.002 ms clock, 0.15+0.032/0.20/0.016+0.043 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 33 @0.013s 6%: 0.009+0.16+0.013 ms clock, 0.14+0.017/0.20/0.010+0.21 ms cpu, 6->6->6 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 34 @0.013s 6%: 0.024+0.16+0.002 ms clock, 0.38+0.012/0.23/0.013+0.040 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 35 @0.013s 6%: 0.009+0.25+0.003 ms clock, 0.14+0.022/0.31/0.007+0.052 ms cpu, 10->10->10 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 36 @0.014s 6%: 0.016+0.20+0.010 ms clock, 0.25+0.037/0.24/0.009+0.16 ms cpu, 12->13->3 MB, 13 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 37 @0.015s 6%: 0.011+0.13+0.002 ms clock, 0.17+0.010/0.17/0.003+0.045 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 38 @0.015s 6%: 0.009+0.22+0.011 ms clock, 0.14+0.013/0.29/0.007+0.18 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 39 @0.015s 6%: 0.010+0.23+0.010 ms clock, 0.17+0.017/0.25/0.010+0.17 ms cpu, 6->7->7 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 40 @0.016s 6%: 0.011+0.21+0.003 ms clock, 0.18+0.025/0.27/0.021+0.051 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 41 @0.017s 6%: 0.010+0.15+0.003 ms clock, 0.17+0.034/0.22/0.025+0.063 ms cpu, 10->11->1 MB, 11 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 42 @0.017s 6%: 0.010+0.14+0.004 ms clock, 0.16+0.011/0.19/0+0.078 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 43 @0.017s 6%: 0.009+0.15+0.017 ms clock, 0.15+0.010/0.19/0+0.27 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 44 @0.018s 6%: 0.009+0.16+0.003 ms clock, 0.14+0.010/0.21/0+0.049 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 45 @0.018s 6%: 0.009+0.19+0.002 ms clock, 0.15+0.029/0.29/0.042+0.045 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 46 @0.018s 6%: 0.009+0.24+0.003 ms clock, 0.15+0.025/0.29/0.011+0.052 ms cpu, 6->7->7 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 47 @0.019s 6%: 0.011+0.27+0.004 ms clock, 0.17+0.015/0.30/0.005+0.064 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 48 @0.020s 6%: 0.011+0.13+0.010 ms clock, 0.17+0.015/0.17/0.006+0.17 ms cpu, 10->10->0 MB, 11 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 49 @0.020s 6%: 0.010+0.11+0.003 ms clock, 0.16+0.014/0.13/0+0.052 ms cpu, 1->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 50 @0.020s 6%: 0.008+0.12+0.011 ms clock, 0.13+0.007/0.15/0+0.18 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 51 @0.021s 6%: 0.022+0.13+0.002 ms clock, 0.36+0.010/0.21/0.006+0.041 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 52 @0.021s 6%: 0.009+0.18+0.004 ms clock, 0.15+0.020/0.23/0.014+0.064 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 53 @0.022s 6%: 0.008+0.22+0.022 ms clock, 0.14+0.036/0.27/0.016+0.36 ms cpu, 6->6->6 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 54 @0.022s 6%: 0.012+0.18+0.014 ms clock, 0.19+0.022/0.22/0.004+0.23 ms cpu, 7->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 55 @0.022s 6%: 0.009+0.34+0.003 ms clock, 0.15+0.019/0.39/0.014+0.062 ms cpu, 9->10->10 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 56 @0.023s 6%: 0.010+0.14+0.003 ms clock, 0.16+0.013/0.19/0.009+0.054 ms cpu, 12->13->3 MB, 13 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 57 @0.023s 6%: 0.009+0.13+0.002 ms clock, 0.15+0.11/0.069/0+0.035 ms cpu, 4->4->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 58 @0.024s 6%: 0.009+0.25+0.011 ms clock, 0.15+0.013/0.29/0.004+0.17 ms cpu, 5->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 59 @0.024s 6%: 0.008+0.17+0.009 ms clock, 0.13+0.013/0.22/0.006+0.15 ms cpu, 6->6->6 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 60 @0.025s 6%: 0.009+0.21+0.002 ms clock, 0.15+0.024/0.27/0.019+0.039 ms cpu, 7->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 61 @0.025s 6%: 0.010+0.28+0.003 ms clock, 0.17+0.022/0.36/0.006+0.050 ms cpu, 9->10->10 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 62 @0.025s 6%: 0.010+0.19+0.003 ms clock, 0.17+0.022/0.23/0.005+0.056 ms cpu, 12->12->2 MB, 12 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 63 @0.026s 6%: 0.009+0.10+0.002 ms clock, 0.15+0.007/0.13/0.005+0.042 ms cpu, 3->3->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 64 @0.026s 6%: 0.012+0.15+0.011 ms clock, 0.19+0.010/0.17/0+0.17 ms cpu, 4->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 65 @0.027s 6%: 0.012+0.20+0.013 ms clock, 0.19+0.009/0.23/0.002+0.21 ms cpu, 6->6->6 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 66 @0.027s 6%: 0.010+0.20+0.002 ms clock, 0.16+0.025/0.28/0.035+0.041 ms cpu, 7->7->7 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 67 @0.027s 6%: 0.011+0.22+0.002 ms clock, 0.18+0.016/0.29/0.014+0.046 ms cpu, 9->9->9 MB, 9 MB goal, 0 MB stacks, 0 MB globals, 16 P
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09% GODEBUG=gctrace=1 GOGC=200 go run .
// gc 1 @0.015s 0%: 0.072+0.37+0.017 ms clock, 1.1+0.097/1.0/0.023+0.27 ms cpu, 7->8->1 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 2 @0.019s 1%: 0.012+0.28+0.003 ms clock, 0.20+0.068/0.77/0.45+0.062 ms cpu, 6->7->1 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// # learning-go/chapters/06-pointers/ex09
// gc 1 @0.001s 6%: 0.093+0.45+0.017 ms clock, 1.4+0.21/1.1/0.12+0.28 ms cpu, 2->3->2 MB, 2 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 2 @0.002s 5%: 0.010+0.23+0.002 ms clock, 0.16+0.043/0.68/0.35+0.046 ms cpu, 3->3->2 MB, 3 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 3 @0.004s 4%: 0.011+0.48+0.021 ms clock, 0.18+0.071/1.0/0.55+0.34 ms cpu, 3->4->3 MB, 3 MB goal, 0 MB stacks, 1 MB globals, 16 P
// gc 4 @0.005s 6%: 0.011+0.71+0.011 ms clock, 0.18+0.36/1.7/1.4+0.18 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 1 MB globals, 16 P
// # learning-go/chapters/06-pointers/ex09
// gc 1 @0.000s 6%: 0.089+0.23+0.031 ms clock, 1.4+0/0.47/0.072+0.49 ms cpu, 1->1->0 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 2 @0.002s 5%: 0.013+0.15+0.012 ms clock, 0.21+0.020/0.36/0.041+0.20 ms cpu, 1->1->1 MB, 1 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 3 @0.002s 5%: 0.012+0.16+0.016 ms clock, 0.19+0/0.41/0.12+0.26 ms cpu, 2->2->2 MB, 2 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 4 @0.003s 5%: 0.010+0.41+0.022 ms clock, 0.17+0.021/0.92/0.29+0.36 ms cpu, 3->4->3 MB, 3 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 5 @0.004s 5%: 0.010+0.24+0.012 ms clock, 0.16+0/0.56/0.25+0.20 ms cpu, 4->5->5 MB, 5 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 6 @0.005s 5%: 0.010+0.19+0.011 ms clock, 0.17+0.016/0.51/0.33+0.18 ms cpu, 6->6->5 MB, 6 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 7 @0.007s 5%: 0.010+0.27+0.014 ms clock, 0.16+0.10/0.70/0.46+0.23 ms cpu, 7->7->6 MB, 7 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 8 @0.007s 5%: 0.011+0.25+0.019 ms clock, 0.18+0/0.67/0.39+0.31 ms cpu, 8->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 9 @0.013s 3%: 0.014+0.40+0.013 ms clock, 0.22+0/1.0/0.64+0.21 ms cpu, 10->10->8 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 10 @0.017s 3%: 0.018+0.33+0.012 ms clock, 0.28+0/0.85/0.58+0.20 ms cpu, 9->9->7 MB, 10 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 11 @0.024s 2%: 0.015+0.37+0.012 ms clock, 0.25+0.019/1.1/0.73+0.20 ms cpu, 9->9->8 MB, 9 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 1 @0.000s 5%: 0.077+0.24+0.014 ms clock, 1.2+0/0.45/0.020+0.23 ms cpu, 7->8->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 2 @0.002s 3%: 0.009+0.18+0.003 ms clock, 0.15+0.027/0.25/0+0.054 ms cpu, 24->26->6 MB, 26 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 3 @0.008s 1%: 0.015+0.22+0.014 ms clock, 0.24+0.031/0.28/0.034+0.23 ms cpu, 17->17->7 MB, 18 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 4 @0.013s 1%: 0.017+0.12+0.013 ms clock, 0.27+0.021/0.18/0.032+0.20 ms cpu, 21->21->1 MB, 22 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 5 @0.014s 1%: 0.011+0.25+0.014 ms clock, 0.17+0.012/0.29/0+0.22 ms cpu, 6->7->7 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 6 @0.017s 1%: 0.013+0.17+0.004 ms clock, 0.21+0.019/0.24/0.013+0.076 ms cpu, 22->24->4 MB, 24 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 7 @0.019s 1%: 0.012+0.10+0.011 ms clock, 0.20+0.013/0.13/0.003+0.17 ms cpu, 11->11->1 MB, 12 MB goal, 0 MB stacks, 0 MB globals, 16 P
// gc 8 @0.019s 1%: 0.009+0.17+0.014 ms clock, 0.15+0.011/0.18/0+0.23 ms cpu, 6->7->7 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 16 P
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09% GODEBUG=gctrace=1 GOGC=off go run .
// [aabuezo@fedora]~/code/go/learning-go/chapters/06-pointers/ex09%
