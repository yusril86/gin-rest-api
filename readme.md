Gin REST API
Deskripsi
Proyek ini adalah implementasi RESTful API menggunakan framework Gin di Go. API ini dirancang untuk memberikan performa tinggi dan kemudahan penggunaan.

Fitur
CRUD (Create, Read, Update, Delete) untuk sumber daya
Middleware untuk otentikasi
Validasi input
Dokumentasi API menggunakan Swagger

Untuk menjalankan Compile daemon 

´´´bash
1. go env GOBIN                                                  
   go env GOPATH

2. export GOPATH="$HOME/go" 

3. export PATH="$PATH:$GOPATH/bin"                                   
4. source ~/.zshrc        
5.  ls "$(go env GOPATH)/bin/CompileDaemon"          
6. CompileDaemon -command="./gin-rest-api" 


´´´
