
# Task Manager

Ein einfacher CLI-basierter Task Manager, entwickelt in Go. Dieses Tool ermöglicht es Ihnen, Aufgaben effizient zu verwalten, indem Sie sie hinzufügen, listen, vervollständigen und löschen.

## Features

- **Aufgaben hinzufügen:** Fügen Sie neue Aufgaben mit Beschreibung und Priorität hinzu.
- **Aufgaben listen:** Zeigen Sie alle Aufgaben oder filtern Sie nach Priorität.
- **Aufgaben vervollständigen:** Markieren Sie Aufgaben als abgeschlossen.
- **Aufgaben löschen:** Entfernen Sie unnötige Aufgaben aus der Liste.

## Voraussetzungen

Stellen Sie sicher, dass Sie Go auf Ihrem Computer installiert haben. Das Projekt wurde mit Go 1.15 und neuer entwickelt und getestet. Sie können Go von der offiziellen [Go-Website](https://golang.org/dl/) herunterladen.

## Installation

Um dieses Projekt zu verwenden, klonen Sie zunächst das Repository auf Ihren lokalen Computer:

```bash
git clone https://github.com/robertwade/my-task-manager.git
cd my-task-manager
```

## Verwendung

Um den Task Manager zu starten, navigieren Sie zum Projektverzeichnis und führen Sie das Programm aus:

```bash
go run cmd/main.go
```

Folgen Sie den Anweisungen im CLI, um Aufgaben zu verwalten. Sie können zwischen verschiedenen Optionen wählen, wie das Hinzufügen oder Anzeigen von Aufgaben.

## Struktur des Projekts

- `cmd/main.go`: Der Hauptzugangspunkt der Anwendung.
- `pkg/task/task.go`: Definiert die Task-Struktur und Konstruktorfunktion.
- `pkg/storage/storage.go`: Verantwortlich für die Verwaltung der Speicherung der Tasks in einer CSV-Datei.

## Lizenz

Dieses Projekt ist unter der MIT Lizenz veröffentlicht. Weitere Details finden Sie in der [LICENSE](LICENSE.md) Datei.

## Kontakt

Falls Sie Fragen haben oder Unterstützung benötigen, zögern Sie nicht, mich über [GitHub Issues](https://github.com/robertwade/my-task-manager/issues) zu kontaktieren.
