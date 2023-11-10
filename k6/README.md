# Benchmark

The benchmark uses [k6](https://k6.io/) so you will need to have that installed.

To run the benchmark first start nungwi. In `./..`:
```
make run 
```
Then to run the benchmark:
```
k6 run bench.js
```

## Example output

```
$ k6 run bench.js

          /\      |‾‾| /‾‾/   /‾‾/   
     /\  /  \     |  |/  /   /  /    
    /  \/    \    |     (   /   ‾‾\  
   /          \   |  |\  \ |  (‾)  | 
  / __________ \  |__| \__\ \_____/ .io

  execution: local
     script: bench.js
     output: -

  scenarios: (100.00%) 1 scenario, 300 max VUs, 5m30s max duration (incl. graceful stop):
           * constantArrivalRate: 600.00 iterations/s for 5m0s (maxVUs: 10-300, gracefulStop: 30s)

     ✓ check status was 200
     ✓ allowed was true

     █ setup

       ✓ write schema status was 200
       ✓ write tuples status was 200

     █ teardown

       ✓ delete tuples status was 200
       ✓ delete schema status was 200

     checks.........................: 100.00% ✓ 91910      ✗ 0    
     data_received..................: 8.4 MB  25 kB/s
     data_sent......................: 13 MB   37 kB/s
     dropped_iterations.............: 134058  391.104208/s
     http_req_blocked...............: avg=13.82µs min=0s      med=2µs   max=57.63ms p(95)=6µs      p(99)=27µs 
     http_req_connecting............: avg=8.54µs  min=0s      med=0s    max=14.19ms p(95)=0s       p(99)=0s   
     http_req_duration..............: avg=1.95s   min=154µs   med=1.65s max=13.85s  p(95)=5.06s    p(99)=6.98s
       { expected_response:true }...: avg=1.95s   min=154µs   med=1.65s max=13.85s  p(95)=5.06s    p(99)=6.98s
     http_req_failed................: 0.00%   ✓ 0          ✗ 45967
     http_req_receiving.............: avg=59.53µs min=5µs     med=34µs  max=33.26ms p(95)=120.69µs p(99)=479µs
     http_req_sending...............: avg=19.74µs min=2µs     med=13µs  max=9.28ms  p(95)=34µs     p(99)=102µs
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s    max=0s      p(95)=0s       p(99)=0s   
     http_req_waiting...............: avg=1.95s   min=138µs   med=1.65s max=13.85s  p(95)=5.06s    p(99)=6.98s
     http_reqs......................: 45967   134.105291/s
     iteration_duration.............: avg=1.95s   min=28.36ms med=1.65s max=35.96s  p(95)=5.06s    p(99)=6.98s
     iterations.....................: 45943   134.035273/s
     vus............................: 0       min=0        max=300
     vus_max........................: 300     min=10       max=300


running (5m42.8s), 000/300 VUs, 45943 complete and 0 interrupted iterations
constantArrivalRate ✓ [======================================] 000/300 VUs  5m0s  600.00 iters/s
```
