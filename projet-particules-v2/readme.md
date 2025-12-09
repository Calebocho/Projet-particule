Fonctionalités
===============
TODO

Usage 
========
modif ....

lancer ....




Dev
=========

Taille du sprite 
------------------
Ajout de Dx, Dy dans src/config/type.go pour avoir la taille du sprite 

Dx                       float64 //JFB largeur du sprite
Dy                       float64 //JFB hauteur du sprite

dans /src/main.go
config.General.Dx = float64(assets.ParticleImage.Bounds().Dx())
config.General.Dy = float64(assets.ParticleImage.Bounds().Dy())
....
