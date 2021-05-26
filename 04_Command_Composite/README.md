# Übung 4 - Command & Composite

```
go run main.go
```

## Umsetzung

Bei den Statefunktionen habe ich statt der ausführung der Methoden direkt die Command funktionalität ausgetauscht.
Für die Composite Implementierung habe ich eine Funktion implementiert, bei welcher die Commandstrings als Array übergeben und danach die Commands in einer Schleife ausgeführt werden.
In der `command.go` Datei wurden alle Commands definiert, als auch die Funktionen für `exec` als auch `execList`

## Fragen oder Erkenntnisse

Aufgrund der zunehmenden komplexität mit weiteren Commands und States ist die State Architektur mitunter ein Grund dieses einfacher zu erweiterung und nicht den überblick zu verlieren.

Bei der Implementierung des Composite Pattern mit Maktros bin ich auf das Problem gestoßen immer mit den akutellen State den Command zuzuweißen. Darum habe ich mich für eine Implementierung mit einem Parameter als Liste von Command Strings und den aktuellen State entschieden. Diese Commands werden sofort innerhalb der `execList` ausgeführt.
