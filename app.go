package main

import (
  "os"
  "fmt"
  "net/http"
  "net"
  "io/ioutil"
  "encoding/json"
  "net/url"
)

func perror(err error) {
  if err != nil {
    panic(err)
  }
}

type Info struct {
  Ip net.IP
}

func main() {
  for _, param := range []string{"CLOUDFLARE_EMAIL", "CLOUDFLARE_TOKEN", "CLOUDFLARE_SUBDOMAIN", "CLOUDFLARE_ROOT_DOMAIN", "CLOUDFLARE_SUBDOMAIN"} {
    if os.Getenv(param) == "" {
      panic(fmt.Sprintf("missing configuration variable: " + param))
    }
  }

  var fulldomain = os.Getenv("CLOUDFLARE_SUBDOMAIN") + "." + os.Getenv("CLOUDFLARE_ROOT_DOMAIN")

  res, err := http.Get("http://jsonip.com")
  perror(err)
  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  perror(err)

  var thisIp Info

  err = json.Unmarshal(body, &thisIp)

  res, err = http.PostForm("https://www.cloudflare.com/api_json.html", url.Values{
      "a": {"rec_load_all"},
      "tkn": {os.Getenv("CLOUDFLARE_TOKEN")},
      "z": {os.Getenv("CLOUDFLARE_ROOT_DOMAIN")},
      "email": {os.Getenv("CLOUDFLARE_EMAIL")},
    })
  perror(err)
  defer res.Body.Close()
  body, err = ioutil.ReadAll(res.Body)
  perror(err)

  type Record struct {
    Rec_Id string
    Name string
  }

  type Record_List struct {
    Result string
    Response struct {
      Recs struct {
        Objs []Record
      }
    }
  }

  var list Record_List
  err = json.Unmarshal(body, &list)

  var targetRecord Record

  for _, value := range list.Response.Recs.Objs {
    if value.Name == fulldomain {
      targetRecord = value
    }
  }

  res, err = http.PostForm("https://www.cloudflare.com/api_json.html", url.Values{
      "a": {"rec_edit"},
      "tkn": {os.Getenv("CLOUDFLARE_TOKEN")},
      "id": {targetRecord.Rec_Id},
      "email": {os.Getenv("CLOUDFLARE_EMAIL")},
      "z": {os.Getenv("CLOUDFLARE_ROOT_DOMAIN")},
      "type": {"A"},
      "content": {thisIp.Ip.String()},
      "service_mode": {"0"},
      "ttl": {"120"},
      "name": {os.Getenv("CLOUDFLARE_SUBDOMAIN")},
    })
  perror(err)
  defer res.Body.Close()
  body, err = ioutil.ReadAll(res.Body)
  perror(err)
}
