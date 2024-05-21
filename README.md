# AZO - zadanie projektowe nr 2
# Badanie efektywności algorytmów grafowych w zależności od rozmiaru instancji oraz sposobu reprezentacji grafu w pamięci komputera.

Autor: [Artur Kręgiel](https://github.com/arkregiel)

Prowadzący: [dr inż. Zbigniew Buchalski](https://wit.pwr.edu.pl/wydzial/struktura-organizacyjna/pracownicy/zbigniew-buchalski)

## Opis projektu

Pełny opis projektu znajduje się na [tej stronie](http://dariusz.banasiak.staff.iiar.pwr.wroc.pl/azo/AZO_lista2.pdf).

Należało zaimplementować oraz dokonać pomiaru czasu działania wybranych algorytmów grafowych rozwiązujących następujące problemy:
- wyznaczanie minimalnego drzewa rozpinającego (MST) - algorytm Prima oraz algorytm Kruskala,
- wyznaczanie najkrótszej ścieżki w grafie – algorytm Dijkstry oraz algorytm Forda-Bellmana,

Algorytmy te należy zaimplementować dla obu poniższych reprezentacji grafu w pamięci komputera:
- reprezentacja macierzowa (macierz incydencji),
- reprezentacja listowa (lista następników/poprzedników). 

Projekt, za zgodą prowadzącego, został zaimplementowany w języku programowania [Go](https://go.dev/).

## Kompilacja i uruchamianie

```
$ go build -o aizo2
$ ./aizo2
```

Program można również uruchomić bez wyraźnego budowania projektu:

```
$ go run .
```

