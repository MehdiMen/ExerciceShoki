package main

import (
   "encoding/json"
   "log"
   "net/http"
   "strings"
   "strconv"

   "github.com/httprouter"
   "github.com/gocolly/colly"
)

type pageInfo struct {
   StatusCode int
   Links      map[string]int
}

func handler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  fmt.Fprint(w, "Welcome!\n\n")
  {/*Definition de l'URL comme argument pour les tests et vérification*/}
   URL := r.URL.Query().Get("url")
   if URL == "" {
      log.Println("Il manque l'URL comme argument")
      return
   }
   log.Println("visiting", URL)

   c := colly.NewCollector()

   p := &pageInfo{Links: make(map[string]int)}

   {/*Utilisation de Colly*/}
   {/*Recherche des données sur les commentaires sur la page avec le mot clé "avis" et récupération de l'entier*/}
   c.OnHTML("a[class]", func(e *colly.HTMLElement) {
      link := e.Text
      if strings.Contains(link, "avis") {
        s := strings.Split(link, " ")
        avis := s[0]
        _ , err := strconv.Atoi(avis)
        if err==nil{
          p.Links[avis]++
        }
      }
   })

   {/*Recherche des données sur la note moyenne en regardant dans la zone appropiée et récupération du flottant*/}
   c.OnHTML("span.restaurants-detail-overview-cards-RatingsOverviewCard__overallRating--nohTl", func(e *colly.HTMLElement) {
      link := e.Text
      if link != "" {
        note:=link[:3]
         p.Links[note]++
      }
   })

   {/*Affichage du statut*/}
   c.OnResponse(func(r *colly.Response) {
      log.Println("response received", r.StatusCode)
      p.StatusCode = r.StatusCode
   })
   c.OnError(func(r *colly.Response, err error) {
      log.Println("error:", r.StatusCode, err)
      p.StatusCode = r.StatusCode
   })

   c.Visit(URL)

   {/*Conversion du résultat en json*/}
   b, err := json.Marshal(p)
   if err != nil {
      log.Println("failed to serialize response:", err)
      return
   }
   w.Header().Add("Content-Type", "application/json")
   w.Write(b)
}

func main() {
   {/*Serveur à l'adresse 8083 et utilisation de httprouter*/}
   addr := ":8083"

   router := httprouter.New()
 	 router.GET("/", handler)

   log.Println("listening on", addr)
   log.Fatal(http.ListenAndServe(addr, router))
}
