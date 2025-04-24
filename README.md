# Bibliothèque de Logging

Cette bibliothèque de logging est conçue pour centraliser et standardiser le processus de logging pour les microservices dans une architecture Kubernetes. Elle permet aux microservices d'envoyer des logs à un proxy de logging, qui peut ensuite les transmettre à divers fournisseurs de logging comme Datadog, Graylog, Atlas MongoDB, ou autres.

## Table des Matières

1. [Introduction](#introduction)
2. [Installation](#installation)
3. [Utilisation](#utilisation)
4. [Architecture Kubernetes](#architecture-kubernetes)
5. [Changement de Fournisseur de Logging](#changement-de-fournisseur-de-logging)
6. [Contribution](#contribution)

## Introduction

La bibliothèque de logging offre une interface simple pour créer et envoyer des logs depuis n'importe quel microservice. Elle utilise une connexion WebSocket pour envoyer les logs à un proxy de logging, qui peut ensuite les transmettre à différents fournisseurs de logging.

## Installation

Pour utiliser cette bibliothèque, ajoutez-la à votre projet Go en tant que dépendance :

```bash
go get github.com/crettien/logger
