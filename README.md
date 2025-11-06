# Salomon en Go

Ce projet est un benchmark en Go de la [fonction de Salomon](https://benchmarkfcns.info/doc/salomonfcn.html) sur un [algorithme de mutation génétique](https://fr.wikipedia.org/wiki/Algorithme_g%C3%A9n%C3%A9tique).

## Usage

Pour configurer le benchmark, il faut modifier le fichier `main.go` en fonction de vos besoins.

```go
game := Game{
    dimensions: 10, // Nombre de dimensions
}
game.Populate(200) // Nombre de cellules
for i := 0; i < 20000; i++ { // Nombre de générations
    game.Generate(35) // Pourcentage des meilleurs éléments à conserver
}
```

## Build

Pour compiler le projet, il faut utiliser la commande suivante :

```bash
go build
```

Puis éxécuter le fichier `salomon-go.exe`