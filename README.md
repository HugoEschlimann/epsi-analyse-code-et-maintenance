# Analyse Maintenance - Application Web avec Gin

Cette application est une API développée en Go avec le framework Gin. Elle utilise MySQL comme base de données et est déployée avec Docker.

## Prérequis

Avant de lancer l'application, assurez-vous d'avoir installé les outils suivants :
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Installation

1. Clonez le dépôt :
    ```bash
    git clone https://github.com/votre-utilisateur/analyse-maintenance.git
    cd analyse-maintenance
    ```
2. Construisez l'image Docker Compose :
    ```bash
    docker compose up --build
    ```

3. Accédez à l'application :
   Ouvrez votre navigateur et allez à `http://localhost:8080`.

4. Accédez à la database :
    ```bash
    docker exec -it analyse-maintenance-db-1 mysql -u user -ppassword
    ```