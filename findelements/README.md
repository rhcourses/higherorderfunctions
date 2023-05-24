# Funktionen, die Elemente in Listen suchen

## Suche von Elementen in `int`-Listen

In der Datei [find_elements_int.go](find_elements_int.go) finden sich drei Funktionen,
die Elemente in `int`-Listen suchen und deren Position zurückgeben.

Alle drei Funktionen sind fast identisch: Sie verwenden alle eine Range-Schleife,
um die Liste zu durchlaufen und prüfen in jedem Schleifendurchlauf, ob das aktuelle
Element dem Suchkriterium entspricht.

Die Funktionen unterscheiden sich also lediglich im Suchkriterium, das in der Schleife
überprüft wird.

Wir wollen dieses Muster nun verallgemeinern, indem wir eine Funktion schreiben, die
die Suche nach einem Element in einer Liste durchführt und dabei als Parameter eine
Funktion übergeben bekommt, die das Suchkriterium prüft.

Solche Funktionen, die eine andere Funktion als Parameter erwarten, werden auch
*Funktionen höherer Ordnung* genannt.

Eine verallgemeinerte Implementierung einer solchen Suchfunktion findet sich in der
Datei [find_elements_int_ho.go](find_elements_int_ho.go) in der
Funktion `FindFirst`.

In der Datei [find_elements_int_ho.go](find_elements_int_ho.go)
finden sich auch Beispiele, wie man die Funktion `FindFirst` verwenden kann.

## Aufgabe: Schreiben Sie weitere Funktionen, die Elemente in Listen suchen

Schreiben Sie Funktionen, die Elemente in Listen suchen und dabei die Funktion
`FindFirst` verwenden, indem sie ihr eine geeignete Hilfsfunktion übergeben.

**Beispiele**:

* Das erste Element in einer Liste finden, das im Intervall [0, 10] liegt.
* Das erste Element in einer Liste finden, das in einem per Parameter übergebenen Intervall liegt.
* Das erste Element in einer Liste finden, das größer als ein per Parameter übergebener Wert ist.
* Das erste Element in einer Liste finden, das nicht in einem gegebenen Intervall liegt.

## Aufgabe: Verallgemeinern Sie das Konzept auf andere Typen

Verallgemeinern Sie das Konzept auf andere Typen, indem Sie eine Variante der Funktion
`FindFirst` schreiben, die mit anderen Typen arbeiten, z.B. `string`, `float64` oder
ein selbst definiertes Struct.

**Beispiele**:

* Das erste Element in einer Liste von `string` finden, das mit einem bestimmten Buchstaben anfängt.
* Das erste Element in einer Liste von `string` finden, das mit per Parameter übergebenen Buchstaben anfängt.
* Das erste Element in einer Liste von `float64` finden, das größer als ein gegebener Wert ist.
* Ein Struct für Personen (z.B. Name, Alter) definieren und die erste aus einer Liste von Personen finden,
  die ein bestimmtes Alter hat.

**Erweiterung dieser Aufgaben:**

* Was fällt Ihnen auf, gibt es Gemeinsamkeiten zwischen den verschiedenen Varianten?
* Können Sie eine Funktion schreiben, die für alle diese Varianten funktioniert?
  * Hinweis: Suchen Sie nach dem Schlagwort *Generics* im Zusammenhang mit Go.
  * Ein Beispiel für eine sehr einfache generische Suchfunktion finden Sie in der Datei
    [find_elements_generic.go](find_elements_generic.go).
