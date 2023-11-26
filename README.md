[![service.web-badge]][service.web-workflow]
[![service.battle-badge]][service.battle-workflow]
[![service.superhero-badge]][service.superhero-workflow]
[![test.integration-badge]][test.integration-workflow]

# Super Smash Heroes

Modern web development using Go, gRPC, Vue.js, PostgreSQL and Docker - a full-stack web application inspired by Super Smash Bros, mainly for learning purposes.

## Setup

To spin up the services:

- Run `docker compose up`
- Visit http://smash.localhost

## Managing Superheroes

Adding new superheroes from the [registry](https://superheroapi.com/ids.html):

```shell
curl -d '{"fullName": "Oliver Queen", "alterEgo": "Green Arrow"}' -H "Content-Type: application/json" -X POST http://localhost/api/superheroes
```

Deleting superheroes:

```shell
curl -X DELETE http://localhost/api/superheroes/{id}
```

[service.web-badge]: https://github.com/jace-ys/super-smash-heroes/workflows/service.web/badge.svg
[service.web-workflow]: https://github.com/jace-ys/super-smash-heroes/actions?query=workflow%3Aservice.web
[service.battle-badge]: https://github.com/jace-ys/super-smash-heroes/workflows/service.battle/badge.svg
[service.battle-workflow]: https://github.com/jace-ys/super-smash-heroes/actions?query=workflow%3Aservice.battle
[service.superhero-badge]: https://github.com/jace-ys/super-smash-heroes/workflows/service.superhero/badge.svg
[service.superhero-workflow]: https://github.com/jace-ys/super-smash-heroes/actions?query=workflow%3Aservice.superhero
[test.integration-badge]: https://github.com/jace-ys/super-smash-heroes/workflows/test.integration/badge.svg
[test.integration-workflow]: https://github.com/jace-ys/super-smash-heroes/actions?query=workflow%3Atest.integrations
