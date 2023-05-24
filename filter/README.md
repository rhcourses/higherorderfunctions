# Filtern von Listen mittels Prädikaten und generischen Funktionen

In der Datei [filter.go](filter.go) findet sich eine Funktion `Filter`, die eine Liste
beliebigen Typs sowie ein Prädikat erwartet.
Ein Prädikat ist eine Funktion, die einen Wert eines bestimmten Typs erwartet und
einen `bool` zurückliefert. Es entscheidet also, ob ein Wert einem bestimmten Kriterium
entspricht oder nicht.

Die Funktion `Filter` liefert eine neue Liste zurück, die nur die Elemente der
ursprünglichen Liste enthält, die dem Prädikat entsprechen.

In der selben Datei gibt es ein Beispiel für die Verwendung der Funktion `Filter`.

## Aufgaben: Filterfunktionen für selbst definierte Typen

Implementieren Sie die Funktionen, die in der Datei [filter_person.go](filter_person.go) beschrieben sind.

In dieser Datei ist ein Typ `Person` definiert, der aus einem Namen und einem Alter besteht.
Außerdem sind eine Reihe von Funktionen vorgegeben, die Sie implementieren können, indem
Sie geeignete Prädikate definieren und diese an die Funktion `Filter` übergeben.
