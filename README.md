# Scraper pracuj.pl
## Użytkowanie
1. Należy uruchomić `main.go` (albo skompilowaną wersję).
1. Stworzyć zapytanie przez przeglądarkę na `localhost:8080/pracuj` podając parametry takie jak:
    * `city` (jeden argument) -- miasto w którym szukamy ofert,
    * `keyword` (wiele argumentów) -- słowa, których szukamy w opisach ofert.

## Struktura projektu
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

## Kompilacja
`go build -o ScraperPracuj.exe -i -ldflags "-s -w"`

## ToDo
* [ ] Dodać odpowiednią licencję:
    * [ ] dodać do repozytorium,
    * [ ] dodać do pliku (wypisanie przed startem).
* [ ] Wstawiać na GitHub nowe wydania (ang. *release*).
* [ ] Dodać sprawdzenie, czy nie wyszła nowa wersja -- jeśli tak, to poinformować użytkownika:
    * [ ] Informacja na stronie.
    * [ ] Informacja na konsoli (sprawdzanie -> wynik: "obecna wersja jest najnowsza" / "wyszła nowa wersja").
* [ ] Stworzyć builder -- uruchomienie kompiluje plik .exe i dodaję numer wersji w nazwie.
* [ ] Dodać przycisk wstecz do strony z wynikami.
* [ ] Dodać możliwość dopasowania dokładnego słów:
    * [ ] checkbox *dopasuj dokładnie*,
    * [ ] funkcjonalność w kodzie.
* [ ] Dodać możliwość wyszukania wszystkich słów:
    * [ ] checkbox *szukaj wszystkich*,
    * [ ] funkcjonalność w kodzie.
* [ ] Zbadać dlaczego niektóre błędy wypisują się w wielu linijkach.
* [ ] Zmienić kod HTML zamieszczony w kodzie na wzorce, które są pakowane za pomocą jakichś pakietów.
* [ ] Zakutalizować `README.md`.
