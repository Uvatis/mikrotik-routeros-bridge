# MikroTik RouterOS Bridge

## English

**MikroTik RouterOS Bridge** is a lightweight Go service that allows you to remotely interact with MikroTik routers through the RouterOS API.
It exposes a REST endpoint so you can execute RouterOS commands and receive structured JSON responses, ideal for integrating MikroTik management into your applications, dashboards, or automation systems.

### 🛠 Features

* Connect to MikroTik routers via the RouterOS API (v6 or v7)
* Execute any RouterOS command
* Return results in JSON format
* Ready-to-use with Docker, just clone and run
* Easily scalable with multiple instances

### 🚀 Quick Start

Make sure Docker is installed, then:

```bash
git clone https://github.com/uvatis/mikrotik-routeros-bridge.git
cd mikrotik-routeros-bridge
docker compose up -d
```

That’s it — the API is available at:

```
http://localhost:8080
```

### Example Request

#### Execute a command:

```json
{
  "host": "192.168.88.1",
  "port": "8728",
  "user": "admin",
  "password": "password",
  "command": "/system/resource/print"
}
```

#### Response:

```json
[
  {
    "architecture-name": "arm",
    "board-name": "hAP ax lite",
    "cpu": "ARM",
    "cpu-count": "2",
    "cpu-frequency": "800",
    "version": "7.15.2 (stable)",
    "uptime": "15h46m8s"
  }
]
```

#### Create a resource (with payload):

```json
{
  "host": "192.168.88.1",
  "port": "8728",
  "user": "admin",
  "password": "password",
  "command": "/ip/hotspot/user/profile/add",
  "payload": {
    "name": "profileTest",
    "shared-users": "1",
    "rate-limit": "10M/10M"
  }
}
```

#### Response:

```json
{ "status": "ok" }
```

---

## 🇫🇷 Français

**MikroTik RouterOS Bridge** est un service léger écrit en Go permettant d’interagir à distance avec les routeurs MikroTik via l’API RouterOS.
Il expose une API REST pour exécuter des commandes RouterOS et récupérer les résultats au format JSON.
C’est l’outil idéal pour intégrer la gestion de routeurs MikroTik dans vos applications ou automatisations.

### ⚙️ Fonctionnalités

* Connexion à un routeur MikroTik via l’API RouterOS (v6 ou v7) directement en Rest
* Exécution de n’importe quelle commande RouterOS
* Résultats renvoyés en JSON
* Déploiement ultra simple avec Docker (`docker compose up -d`)
* Extensible et compatible multi-instances

### 🚀 Démarrage rapide

Assurez-vous d’avoir **Docker** installé, puis :

```bash
git clone https://github.com/uvatis/mikrotik-routeros-bridge.git
cd mikrotik-routeros-bridge
docker compose up -d
```

L’API sera disponible à l’adresse suivante :

```
http://localhost:8080
```

### Exemple de requête

#### Exécution d’une commande :

```json
{
  "host": "192.168.88.1",
  "port": "8728",
  "user": "admin",
  "password": "password",
  "command": "/system/resource/print"
}
```

#### Réponse :

```json
[
  {
    "architecture-name": "arm",
    "board-name": "hAP ax lite",
    "cpu": "ARM",
    "cpu-count": "2",
    "cpu-frequency": "800",
    "version": "7.15.2 (stable)",
    "uptime": "15h46m8s"
  }
]
```

#### Création d’une ressource :

```json
{
  "host": "192.168.88.1",
  "port": "8728",
  "user": "admin",
  "password": "password",
  "command": "/ip/hotspot/user/profile/add",
  "payload": {
    "name": "profileTest",
    "shared-users": "1",
    "rate-limit": "10M/10M"
  }
}
```

#### Réponse :

```json
{ "status": "ok" }
```
