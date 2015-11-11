# image-https-proxy
A simple image https proxy to proxy an image from http.
If you have a https secure page and embed a non secure http image,
the internet browser will report a partial non encrypted link and you will lose the green https lock sign.
This image proxy will be a proxy between https to http image url. It is made by golang.



Build:
<pre>
cd image-https-proxy
export GOPATH=${PWD}
cd src/image-https-proxy
go install
cd ../../bin
export PROXY_PORT=8000; ./image-https-proxy
</pre>

Proxy will listen to port 8000, by default it will listen to port 8080



Try the proxy from browser

http://localhost:8000/?d=http://apartemencitraliving.com/img/header-bg.jpg


# Tentang Fojubel: 
#### Fojubel adalah website jual beli properti, elektronik, otomotif, fashion dan lainnya di Indonesia yang mudah dan aman.
Untuk keterangan lebih lanjut silahkan buka http://www.fojubel.com
