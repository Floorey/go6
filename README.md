Blockchain-Datenverarbeitung in Go

Dieses Projekt implementiert eine einfache Anwendung in Go, die Daten aus JSON- und CSV-Dateien einliest und in einer Blockchain speichert. Jeder Datensatz wird in einem separaten Block abgelegt, um die Integrität und Nachvollziehbarkeit der Daten sicherzustellen.
Installation

    Stellen Sie sicher, dass Go auf Ihrem System installiert ist. Anweisungen zur Installation finden Sie auf der Go-Website.

    Klonen Sie das Repository auf Ihren lokalen Computer:

    sh

git clone https://github.com/dein-benutzername/dein-repo.git

Navigieren Sie in das Verzeichnis des Projekts:

sh

cd dein-repo

Führen Sie das Programm aus:

sh

go run main.go -csv /pfad/zu/deiner/csv/datei.csv

oder

sh

    go run main.go -json /pfad/zu/deiner/json/datei.json

    Wenn kein Befehlszeilenargument angegeben wird, werden Sie aufgefordert, die Daten manuell einzugeben.

Verwendung

    Um Daten aus einer CSV-Datei zu lesen, verwenden Sie das Befehlszeilenargument -csv gefolgt vom Pfad zur CSV-Datei.
    Um Daten aus einer JSON-Datei zu lesen, verwenden Sie das Befehlszeilenargument -json gefolgt vom Pfad zur JSON-Datei.
    Wenn keine Dateipfade angegeben werden, können Sie die Daten manuell eingeben, wenn Sie dazu aufgefordert werden.

Beiträge

Beiträge sind willkommen! Wenn Sie Verbesserungen vornehmen möchten, erstellen Sie einfach eine Pull-Anfrage mit Ihren Änderungen.
