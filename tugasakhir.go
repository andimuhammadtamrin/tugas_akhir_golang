package main

import "encoding/json"
import "fmt"
import "net/http"
import "bytes"
import "net/url"
import "database/sql"
import _ "mysql-master"

var baseURL = "http://localhost:8080"

type daftar_buku struct{
  ID string
  Judul string
  Pengarang string
  Tahun int
}

func koneksi()(*sql.DB, error){
  db,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/db_daftarbuku")
  if err != nil{
    return nil,err
  }
  return db,nil
}



func ambil_api()([]daftar_buku,error){
  var err error
  var client = &http.Client{}
  var data []daftar_buku

  request,err := http.NewRequest("POST",baseURL + "/daftar",nil)

  if err != nil{
    return nil,err
  }

  response ,err  := client.Do(request)
  if err != nil{
    return nil,err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err !=nil{
    return nil,err
  }
  return data,nil
}


func cari_dataapi(daftar string)(daftar_buku,error){
  var err error
  var client = &http.Client{}
  var data daftar_buku

  var param = url.Values{}
  param.Set("Judul",daftar)
  var payload = bytes.NewBufferString(param.Encode())


  request,err := http.NewRequest("POST",baseURL + "/cari_buku",payload)

  if err != nil{
    return data,err
  }
  request.Header.Set("Content-Type","application/x-www-form-urlencoded")

  response ,err  := client.Do(request)
  if err != nil{
    return data,err
  }
  defer response.Body.Close()

  err = json.NewDecoder(response.Body).Decode(&data)
  if err !=nil{
    return data,err
  }
  return data,nil
}

func ubah_api(){
  db, err := koneksi()
  if err != nil{
    fmt.Println(err.Error())
    return
  }
  defer db.Close()
  _, err = db.Exec("update tbl_buku set Judul=? where Judul = ? ","Ilmu Ukur","Ilmu Ukur Wilayah")
  if err != nil{
    fmt.Println(err.Error())
    return
  }
  fmt.Println("Data berhasi di Ubah")

}

func hapus_api(){
  db, err := koneksi()
  if err != nil{
    fmt.Println(err.Error())
    return
  }
  defer db.Close()
  _, err = db.Exec("delete from tbl_buku  where ID = ? ","B04")
  if err != nil{
    fmt.Println(err.Error())
    return
  }
  fmt.Println("Hapus berhasil")
}



func tampilkan_daftar(){
  var daftar, err = ambil_api()
  if err != nil{
    fmt.Println("Buku tidak tersedia!",err.Error())
    return
  }
  fmt.Println("Data Berhasil Di Tampilkan")
  for _, each := range daftar{

      fmt.Println(" ID : ",each.ID," Judul : ",each.Judul," Pengarang : ",each.Pengarang," Tahun : ",each.Tahun)
  }
}

func cari_databuku(){
  var daftar, err = cari_dataapi("SMART GRAMMAR")
  if err != nil{
    fmt.Println("Buku tidak tersedia!",err.Error())
    return
  }
    fmt.Println("Data Berhasil Di Tampilkan")
    fmt.Println(" ID : ",daftar.ID," Judul : ",daftar.Judul," Pengarang : ",daftar.Pengarang," Tahun : ",daftar.Tahun)

}



func main(){
  hapus_api()
  ubah_api()
  cari_databuku()
  tampilkan_daftar()

}
