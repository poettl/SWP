# Übung 5 - Mediator & Observer

## Development server

Please install NodeJS and the Angular CLI before.

Run `ng serve` for a dev server. Navigate to `http://localhost:4200/`.

## Umsetzung

Bei dieser Übung habe ich mich für das Angular Framework entschieden da es sehr viele Vorteile im Bereich UI Kompenenten und State Management bietet.
Der Mediator wurde im prinzipe mit Hilfe der StorageMap (key value pair aus String und einem Storage Objekte). Hierbei habe ich beim setValue des Storage Objekt einen Notifer eingebaut damit damit mögliche Observers benachrichtig werden.

## Fragen oder Erkenntnisse

Beim Mediatorpattern sehe ich den Bezug zur Praxis nicht so ganz. Aufgrund der Componentbased Architechutre von modernen Frontend Frameworks ist das Mediatorpattern nicht wirklich notwendig. Hierbei werden die benötigten Daten direkt in der Kompenente gehalten. Somit werden die unterschiedlich benötigten Werte gut abstrahiert. Meiner Meinung nach hat das Mediatorpattern den aktuellen Code sogar komplizierter und anfänglich undurchschaubaer gemacht. Im Bezug auf das aktuellen Beispiel hätte ich einen List Komponente gebaut welche mit der listSelectionChanged Methode nur den Wert in eine Variable gespeichert hätte. Im HTML Template hätte ich das diabled Feld der Buttons einfach abhänig von Wert disabled bzw. enabled und der ChangeDetector von Angular hätte sich um den Rest gekümmert.

Ganz anders sieht es beim Observable Pattern aus. Das könnte man sich im Reative Programming gar nicht mehr wegdenken. Es war sehr interessant sowetwas selbst zu entwicklen aber bei einem echten Projekt würde ich mich da eher auf rxJS zurückgreifen.
