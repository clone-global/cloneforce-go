# Changelog

## 0.3.0 (2026-07-18)

Full Changelog: [v0.2.0...v0.3.0](https://github.com/clone-global/cloneforce-go/compare/v0.2.0...v0.3.0)

### Features

* **client:** optimize json encoder for internal types ([7732587](https://github.com/clone-global/cloneforce-go/commit/7732587833c2c8b504a82546e501de601fa91a54))
* **go:** add default http client with timeout ([d69b81a](https://github.com/clone-global/cloneforce-go/commit/d69b81a43f671eb06ed02b15a2bd74023c1266cf))
* **stlc:** configurable CI runner and private-production-repo support in workflow templates ([277e5d0](https://github.com/clone-global/cloneforce-go/commit/277e5d06ef675c470e3a3c9512d5e0aa5eca654a))
* support setting headers via env ([15c5181](https://github.com/clone-global/cloneforce-go/commit/15c51814cbcd5a8ffee39ef2f32b1ac96522c38f))


### Bug Fixes

* fix issue with unmarshaling in some cases ([dec9119](https://github.com/clone-global/cloneforce-go/commit/dec911915d54fcb727b3d32d373f74896a5f001e))
* **go:** avoid panic when http.DefaultTransport is wrapped ([4e57aaf](https://github.com/clone-global/cloneforce-go/commit/4e57aafd58c9c045c646cc69c8cf10f1e41bb18d))


### Chores

* avoid embedding reflect.Type for dead code elimination ([a5f7b19](https://github.com/clone-global/cloneforce-go/commit/a5f7b1921fda5c29367ad19de6ad730d69f3b1fd))
* **internal:** more robust bootstrap script ([c207fd3](https://github.com/clone-global/cloneforce-go/commit/c207fd31031d46e35ade5fdb03b6de7786be858a))
* redact api-key headers in debug logs ([29d93da](https://github.com/clone-global/cloneforce-go/commit/29d93daecd3fb297a0d7c7efdd9ebcd5a62820a8))

## 0.2.0 (2026-03-29)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/clone-global/cloneforce-go/compare/v0.1.0...v0.2.0)

### Features

* **internal:** support comma format in multipart form encoding ([07f273c](https://github.com/clone-global/cloneforce-go/commit/07f273c793d7e29f2ea44def8a4656a0fee1305d))


### Bug Fixes

* prevent duplicate ? in query params ([8d7524a](https://github.com/clone-global/cloneforce-go/commit/8d7524a97939b189117b3fd31bdd60746e39221b))


### Chores

* **ci:** skip lint on metadata-only changes ([d7b94b2](https://github.com/clone-global/cloneforce-go/commit/d7b94b26c06929cdf5263e7c4c0766631818d17b))
* **ci:** support opting out of skipping builds on metadata-only commits ([987cb34](https://github.com/clone-global/cloneforce-go/commit/987cb34482b846f6fa9ae832bd0bbb02daccfb17))
* **client:** fix multipart serialisation of Default() fields ([f748914](https://github.com/clone-global/cloneforce-go/commit/f748914ea56bf1c8882466850527834e78e50673))
* **internal:** support default value struct tag ([21ebc37](https://github.com/clone-global/cloneforce-go/commit/21ebc379b7115d9f2f383d1c0a1cd830f42cb196))
* **internal:** update gitignore ([d82bce5](https://github.com/clone-global/cloneforce-go/commit/d82bce534ba34fd99a9585f215c2909ecbd4ed24))
* remove unnecessary error check for url parsing ([7b77550](https://github.com/clone-global/cloneforce-go/commit/7b77550c03dcbf22d67d6b7aab6a3de570480e84))
* update docs for api:"required" ([6106f5b](https://github.com/clone-global/cloneforce-go/commit/6106f5bddda87d9ef30241bac8bea15e69d3fe8b))

## 0.1.0 (2026-03-22)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/clone-global/cloneforce-go/compare/v0.0.2...v0.1.0)

### Features

* update SDK versions ([757e651](https://github.com/clone-global/cloneforce-go/commit/757e6515eec47031a99be1d809e150a6b8535e8e))

## 0.0.2 (2026-03-21)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/clone-global/cloneforce-go/compare/v0.0.1...v0.0.2)

### Chores

* update SDK settings ([58220ee](https://github.com/clone-global/cloneforce-go/commit/58220eecabef441f409d67bb1ba52dc4c1d2c58b))
