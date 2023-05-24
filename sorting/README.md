# Sortieren mittels Komparatoren

In der Datei [sort.go](sort.go) finden Sie eine Funktion `Sort`, die eine Liste beliebigen Typs sowie einen Komparator erwartet.
Diese Funktion demonstriert, wie die Sortierfunktiones aus der Standardbibliothek von Go funktionieren.

Ein Komparator ist in diesem Fall eine Funktion, die zwei Positionen aus der Liste erwartet und entscheidet, ob das Element an der ersten Position kleiner als das Element an der zweiten Position ist.
D.h. ein Komparator bestimmt für ein generisches Sortierverfahren, wonach sortiert werden soll.

## Aufgabe: Sortieren von Listen beliebigen Typs

Es ist eine Beispielfunktion vorgegeben, die mittels Komparator eine Liste von Zahlen
sortiert.
Darüber hinaus gibt es eine Reihe von Sortierfunktionen für komplexere Typen,
die Sie implementieren können, indem Sie geeignete Komparatoren definieren und diese an die Funktion `Sort` übergeben.
