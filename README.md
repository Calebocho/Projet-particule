# 🌌 Project Particles

Une simulation de système de particules en 2D développée en **Go** à l'aide de la bibliothèque de création de jeux 2D **Ebiten**. 

Ce projet a été réalisé dans le cadre de la formation BUT Informatique à l'IUT de Nantes (Nantes Université). Il démontre la gestion de moteurs de rendu, l'optimisation des performances avec un grand nombre d'entités, et l'application de règles physiques simples dans un environnement graphique.

## ✨ Fonctionnalités

*   **Rendu performant :** Affichage fluide de plusieurs milliers de particules simultanément grâce au moteur Ebiten.
*   **Génération dynamique :** Apparition continue ou ponctuelle de particules depuis des générateurs (emitters).
*   **Physique et Cycle de vie :** Gestion de la vélocité, de la trajectoire, de l'opacité et de la durée de vie des particules.
*   **Personnalisable :** Architecture permettant de modifier facilement les comportements (gravité, rebonds, couleurs).

## 🛠️ Technologies utilisées

*   **Langage :** [Go (Golang)](https://go.dev/)
*   **Moteur graphique :** [Ebiten](https://ebitengine.org/) (A dead simple 2D game library for Go)

## 📋 Prérequis

Pour compiler et exécuter ce projet localement, vous devez avoir installé :

1.  **Go** (version 1.18 ou supérieure recommandée).
2.  Les dépendances systèmes requises par Ebiten. Si vous êtes sous Linux (comme Fedora/Ubuntu), vous aurez potentiellement besoin des bibliothèques de développement audio et vidéo (ex: `alsa-lib-devel`, `libX11-devel`, `libXrandr-devel`).
