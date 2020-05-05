# Scraper pracuj.pl

This is a software written in Go language which allows quick searching for given keywords in all offers found at Polish job board `pracuj.pl`. The rest of this README is written in Polish.

Program umożliwiający szybkie przeszukanie wszystkich ofert w serwisie `pracuj.pl` w danym mieście za pomocą wyszukiwania .


## Użytkowanie
1. Należy uruchomić `main.go` (albo skompilowaną wersję).
1. Wpisać nazwę miasta oraz słowa klucze.
1. Uruchomić wyszukiwanie (**uwaga:** wyszukiwanie może zająć do kilku minut - w zależności od prędkości łącza internetowego).
1. Po zakończeniu pracy z programem zamknąć jego okno (konsolę).

**Uwaga:** Wyszukiwane słowa kluczowe są wyszukiwane jako podciągi znaków, ponadto wielkość liter nie ma znaczenia. Na przykład *inżynier* zostanie znaleziony w słowie w ofercie zawierającej słowo *Inżynierska*.


### Informacje techniczne
#### Struktura projektu
* **scrapers**
    * `Collector.go` -- funkcje do inicjalizacji kolektora z modułu `colly`.
    * `ScraperPracuj.go` -- funkcje dodające obsługę poszczególnych elementów HTML do kolektora.
* **server**
    * `Client.go` -- zawiera definicję klienta serwera.
    * `MainHandler.go` -- zawiera stronę główną programu.
    * `PracujHandler.go` -- funkcja obsługująca endpoint `/pracuj?`.
    * `Run.go` -- powiązanie endpointów z funkcjami obsługującymi je.
* `main.go` -- uruchomienie serwera.
* `README.md` -- ten plik.

#### Kompilacja
`go build -o ScraperPracuj.exe -i -ldflags "-s -w"`
