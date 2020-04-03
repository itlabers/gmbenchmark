
### sm2   
 
#### mixbee与 tjfoc  签名与验证默认userId =1234567812345678   

#### flyinox 签名与验证消息是经过hash 的结果，未提供杂凑值Za||M计算 

#### 提供sm2 与golang原生ecdsa算法性能比较
```
goos: darwin
goarch: amd64
pkg: benchmark/sm2
BenchmarkECDSAGenKeyPair/ecdsa-4               38439	     28067 ns/op	     608 B/op	      12 allocs/op
BenchmarkECDSAGenKeyPair/mixbee-sm2-4    	   32629	     34485 ns/op	    1072 B/op	      20 allocs/op
BenchmarkECDSAGenKeyPair/tjfoc-sm2-4     	    2072	    554709 ns/op	    2194 B/op	      43 allocs/op
BenchmarkECDSAGenKeyPair/flyinox-sm2-4   	     270	   4419475 ns/op	  866885 B/op	    9236 allocs/op
BenchmarkSign/ecdsa-4                    	   17396	     67241 ns/op	    2689 B/op	      32 allocs/op
BenchmarkSign/mixbee-sm2-4               	   18848	     55259 ns/op	    4472 B/op	      67 allocs/op
BenchmarkSign/tjfoc-sm2-4                	    2276	    579714 ns/op	    4933 B/op	      77 allocs/op
BenchmarkSign/flyinox-sm2-4              	     309	   3935704 ns/op	  868950 B/op	    9269 allocs/op
BenchmarkVerify/ecdsa-4                  	   10000	    143345 ns/op	     928 B/op	      16 allocs/op
BenchmarkVerify/mixbee-sm2-4             	    6630	    158595 ns/op	    3522 B/op	      35 allocs/op
BenchmarkVerify/tjfoc-sm2-4              	     391	   3669915 ns/op	   77524 B/op	    1649 allocs/op
BenchmarkVerify/flyinox-sm2-4            	     159	   7748317 ns/op	 1734323 B/op	   18513 allocs/op
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