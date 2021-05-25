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
