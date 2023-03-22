# Doku Creator zur automatisierten Dokumentation

Dieses Go-Programm ermöglicht es Benutzern, den Inhalt von Dateien in einem Ordner rekursiv zu durchsuchen und automatisch eine Markdown-Dokumentation für jede Datei zu erstellen. Die Dokumentation wird mithilfe der OpenAI API generiert und enthält eine Zusammenfassung des Codes sowie mögliche nächste Schritte.


## Verwendung

- Laden Sie das Programm herunter und speichern Sie es auf Ihrem Computer.
- Expose das ENV: OPEN_AI_API_TOKEN mit deinem Open.ai API Token
- Öffnen Sie eine Befehlszeile oder ein Terminal und navigieren Sie zum Speicherort der heruntergeladenen Datei.
- Geben Sie den Befehl "go run main.go <Dateiendung> <Repository-URL>" ein und drücken Sie die Eingabetaste.
- Der Befehl führt das Programm aus und generiert automatisch eine Markdown-Datei für jede Go-Datei im angegebenen Repository.


### Enviorment Variablen

Expose das ENV: OPEN_AI_API_TOKEN mit deinem Open.ai API Token

### Parameter

`<Dateiendung>` (erforderlich): Die Dateiendung der zu durchsuchenden Dateien. Zum Beispiel "go" für Go-Dateien.
`<Repository-URL>` (erforderlich): Die URL des Repositories, das durchsucht werden soll.


### Beispiel

```go
go run main.go go https://github.com/mein-benutzername/mein-repo.git
```

Dieser Befehl durchsucht das Repository "mein-repo" nach allen Go-Dateien und generiert automatisch eine Markdown-Datei für jede gefundene Datei.

## Funktionen

- Durchsucht rekursiv alle Dateien im angegebenen Repository.
- Generiert automatisch eine Markdown-Datei für jede gefundene Go-Datei.
- Verwendet die OpenAI API, um eine Zusammenfassung des Codes und mögliche nächste Schritte zu generieren.
- Erstellt automatisch ein Verzeichnis für die Dokumentation, falls nicht vorhanden.

## Anforderungen

Um dieses Programm ausführen zu können, benötigen Sie eine funktionierende Go-Installation sowie einen API-Schlüssel für die OpenAI API. Der API-Schlüssel muss als ENV-Variable OPEN_AI_API_TOKEN vorliegen. Außerdem muss das Git-Tool installiert sein, um das Repository zu klonen.

## Autor

Dieses Programm wurde von [Manuel Engelhardt](https://about.me/manuelmueller) entwickelt, einem erfahrenen Softwareentwickler und Gründer der Firma ITDevOps.de.

[ITDevOps.de](https://ITDevOps.de) ist ein Unternehmen, das sich auf die Entwicklung von Softwarelösungen und die Bereitstellung von IT-Beratungsdiensten spezialisiert hat. Mit einem Team von erfahrenen Entwicklern und Beratern bietet [ITDevOps.de](https://ITDevOps.de) seinen Kunden eine breite Palette von Dienstleistungen an, darunter Softwareentwicklung, Systemadministration, DevOps-Implementierung, Cloud-Beratung und mehr.

[Manuel Engelhardt](https://about.me/manuelmueller) ist ein erfahrener Entwickler, der über umfangreiche Kenntnisse in den Bereichen Softwareentwicklung, Systemadministration, Cloud-Computing und DevOps verfügt. In seiner Freizeit trägt Manuel Engelhardt zur Open-Source-Community bei und teilt sein Wissen durch Konferenzvorträge und Schulungen.