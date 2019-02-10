# nginx

## localhost의 어느 포트로 연결되어있는 서버로 port forwarding 실패시..

- `$ cat /var/log/nginx/error.log` 로 에러 로그를 확인한다.

  ```
$ cat /var/log/nginx/error.log
2019/02/10 13:41:08 [crit] 29414#0: *1 connect() to 127.0.0.1:8080 failed (13: Permission denied) while connecting to upstream, client: {실제-요청IP}, server: _, request: "GET / HTTP/1.1", upstream: "http://127.0.0.1:8080/", host: "실제-요청IP"
2019/02/10 13:41:08 [crit] 29414#0: *1 connect() to [::1]:8080 failed (13: Permission denied) while connecting to upstream, client: {실제-요청IP}, server: _, request: "GET / HTTP/1.1", upstream: "http://[::1]:8080/", host: "실제-요청IP"
  ```
- CentOS에서는 OS차원에서 앱이 네트워크를 사용하는 것을 막아놓았기 때문에 포트포워딩이 되지 않기에 옵션을 켜주어야한다.
  - `/usr/sbin/setsebool httpd_can_network_connect true`

### Reference
- https://stackoverflow.com/questions/25235453/nginx-proxy-server-localhost-permission-denied
