# image-https-proxy
A simple image https proxy to proxy an image from http.
If you have a https secure page and embed a non secure http image,
the internet browser will report a partial non encrypted link and you will lose the green https lock sign.
This image proxy will be a proxy between https to http image url. It is made by golang.



Build:
export GOPATH=image-https-proxy
cd src/image-https-proxy
go build
go install


Run the proxy

http://www.example/url?d=http://apartemencitraliving.com/img/header-bg.jpg


Tentang Fojubel: 
Fojubel adalah website jual beli properti, elektronik, otomotif, fashion dan lainnya di Indonesia yang mudah dan aman. 
Untuk keterangan lebih lanjut silahkan buka http://www.fojubel.com
