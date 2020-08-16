<h1 align="center">
      <img alt="SuperAPI" title="SuperAPI" src=".github/logo.png" width="300px" />
</h1>

<h3 align="center">
  SuperAPI
</h3>

<p align="center">SuperAPI, create supers and manage them 🦸</p>
<p align="center">Made with Golang and PostgreSQL 🚀</p>
<p align="center">Using Docker, Gorm, UUID, Testify, Mux, and Gjson ✔️</p>
<p align="center">Integration with superheroapi ✔️</p>

<p align="center">
  <img alt="GitHub language count" src="https://img.shields.io/github/languages/count/Lgdev07/superapi?color=%2304D361">

  <img alt="Made by Lgdev07" src="https://img.shields.io/badge/made%20by-Lgdev07-%2304D361">

  <img alt="License" src="https://img.shields.io/badge/license-MIT-%2304D361">

  <a href="https://github.com/Lgdev07/superapi/stargazers">
    <img alt="Stargazers" src="https://img.shields.io/github/stars/Lgdev07/superapi?style=social">
  </a>
</p>

<p align="center">
  <a href="#-installation-and-execution">Installation and execution</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-available-routes">Available Routes</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#-how-to-contribute">How to contribute</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

## 🚀 Installation and execution

1. Clone this repository and go to the directory;
2. Rename sample .env;

<h4> 🔧 Development </h4>

1. Run `docker-compose up`;
2. Make the Requests;

<h4> 🧪 Tests </h4>

1. Run `docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit`;

## 🛣️ Available Routes

POST - Create Super:
  - /supers <br>
  Expected Json Body:<br>
  {"name": "Custom Super Name"}

GET - List Supers by Query Params:
- /supers <br>
  Optional Query Params<br>
  alignment: "good" or "bad"<br>
  name: Super Name<br>
  example: ?alignment=good&name=Super%20Name<br>

GET - Get Super by UUID:
- /supers/{uuid} <br>
  Requires Route Params<br>
  uuid: uuid_type<br>
  example: /supers/123e4567-e89b-12d3-a456-426614174000<br>

DELETE - Delete Super by UUID:
- /supers/{uuid} <br>
  Requires Route Params<br>
  uuid: uuid_type <br>
  example: /supers/123e4567-e89b-12d3-a456-426614174000<br>

## 🤔 How to contribute

- Fork this repository;
- Create a branch with your feature: `git checkout -b my-feature`;
- Commit your changes: `git commit -m 'feat: My new feature'`;
- Push to your branch: `git push origin my-feature`.

After the merge of your pull request is done, you can delete your branch.

---