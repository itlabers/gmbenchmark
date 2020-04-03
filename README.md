
### sm2   
 
#### mixbee与 tjfoc  签名与验证默认userId =1234567812345678   

#### flyinox 签名与验证消息是经过hash 的结果，未提供杂凑值Za||M计算 

#### 提供sm2 与golang原生ecdsa算法性能比较
```
goos: darwin
goarch: amd64
pkg: benchmark/sm2
BenchmarkECDSAGenKeyPair/ecdsa-4         	   40318	     27249 ns/op	     608 B/op	      12 allocs/op
BenchmarkECDSAGenKeyPair/mixbee-sm2-4    	   35164	     32346 ns/op	    1072 B/op	      20 allocs/op
BenchmarkECDSAGenKeyPair/tjfoc-sm2-4     	    2240	    508623 ns/op	    2194 B/op	      43 allocs/op
BenchmarkECDSAGenKeyPair/flyinox-sm2-4   	     322	   3621020 ns/op	  867267 B/op	    9240 allocs/op
BenchmarkSign/ecdsa-4                    	   25707	     48930 ns/op	    2689 B/op	      32 allocs/op
BenchmarkSign/mixbee-sm2-4               	   19486	     59654 ns/op	    4500 B/op	      67 allocs/op
BenchmarkSign/tjfoc-sm2-4                	    2377	    622073 ns/op	    4932 B/op	      77 allocs/op
BenchmarkSign/flyinox-sm2-4              	     313	   3675831 ns/op	  867601 B/op	    9255 allocs/op
BenchmarkVerify/ecdsa-4                  	    7528	    136934 ns/op	     928 B/op	      16 allocs/op
BenchmarkVerify/mixbee-sm2-4             	    7420	    164178 ns/op	    3474 B/op	      34 allocs/op
BenchmarkVerify/tjfoc-sm2-4              	     356	   3238721 ns/op	   76068 B/op	    1614 allocs/op
BenchmarkVerify/flyinox-sm2-4            	     156	   7498268 ns/op	 1765864 B/op	   18821 allocs/op
```



### sm3 
#### 提供sm3 与openssl(v1.1.1d) sm3(cgo)算法性能比较

```
goos: darwin

​goarch: amd64

​pkg: benchmark/sm3

​BenchmarkSM3/openssl-sm3-4                458504              2789 ns/op              92 B/op          5 allocs/op

​BenchmarkSM3/sha256-4                    1844538               592 ns/op             160 B/op          2 allocs/op

​BenchmarkSM3/mixbee-sm3-4                 806316              1468 ns/op             144 B/op          2 allocs/op

​BenchmarkSM3/tjfoc-sm3-4                  789799              1679 ns/op             216 B/op          6 allocs/op

​BenchmarkSM3/flyinox-sm3-4                973172              1198 ns/op             144 B/op          2 allocs/op 
```
 

sm4 

```
goos: darwin
goarch: amd64
pkg: benchmark/sm4
BenchmarkDecrypt/mixbee-sm4-4         	 2495433	       463 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrypt/tjfoc-sm4-4          	 5165262	       242 ns/op	       0 B/op	       0 allocs/op
BenchmarkDecrypt/flyinox-sm4-4        	  735567	      1687 ns/op	      48 B/op	       3 allocs/op
BenchmarkEncrypt/mixbee-sm4-4         	 2510671	       452 ns/op	       0 B/op	       0 allocs/op
BenchmarkEncrypt/tjfoc-sm4-4          	 5030077	       232 ns/op	       0 B/op	       0 allocs/op
BenchmarkEncrypt/flyinox-sm4-4        	  444978	      2619 ns/op	     112 B/op	       5 allocs/op
```