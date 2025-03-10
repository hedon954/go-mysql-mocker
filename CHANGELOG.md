# Changelog

---
## [1.0.7](https://github.com/hedon-go-road/template-web/compare/v1.0.6..v1.0.7) - 2025-03-04

feat(go): upgrade go version to 1.24.0

### ⚙️ Miscellaneous Chores

- upgrade go version to 1.24.0 - ([c4ca277](https://github.com/hedon-go-road/template-web/commit/c4ca277196bf703cf5f6d1a5b4300141adcddbb7)) - hedon954

### 🐛 Bug Fixes

- **(ci)** replace exportloopref with copyloopvar - ([260a84d](https://github.com/hedon-go-road/template-web/commit/260a84d0fc665d1df8a06322bde6adfad71a8aa1)) - hedon954

### 📚 Documentation

- **(readme)** upgrade - ([8b4b22b](https://github.com/hedon-go-road/template-web/commit/8b4b22bbf5a15456a8bd30ebdfb775e27afddc9f)) - hedon954

<!-- generated by git-cliff -->

---
## [1.0.6](https://github.com/hedon-go-road/template-web/compare/v1.0.5..v1.0.6) - 2025-03-04

provide `GetPort` method for `Builder`

### ⛰️ Features

- **(builder)** provide `GetPort` method for `Builder` - ([7617246](https://github.com/hedon-go-road/template-web/commit/761724610a12bc0014d60b409c57a47c3bad4817)) - hedon954

<!-- generated by git-cliff -->

---
## [1.0.5](https://github.com/hedon-go-road/template-web/compare/v1.0.4..v1.0.5) - 2024-12-30

set gorm client silent log mod default

### ⛰️ Features

- **(builder)** set gorm silent log mod default - ([1f1a85c](https://github.com/hedon-go-road/template-web/commit/1f1a85c183ef68bf35b0a1c1c81dc79baf1e1320)) - hedon954

<!-- generated by git-cliff -->

---
## [1.0.4](https://github.com/hedon-go-road/template-web/compare/v1.0.3..v1.0.4) - 2024-11-26

add atomic bool to check if has been built

### 🐛 Bug Fixes

- **(gmm)** add atomic bool to check if has been built - ([f6e062a](https://github.com/hedon-go-road/template-web/commit/f6e062a723c548fbfa5c6531af8d13b6e3844141)) - hedon954

### 📚 Documentation

- **(readme)** fix it - ([8a8d92c](https://github.com/hedon-go-road/template-web/commit/8a8d92c1bc4f26f69fafb703e8e21163b540a887)) - hedon954
- **(readme)** update go reference link - ([eab6b1e](https://github.com/hedon-go-road/template-web/commit/eab6b1e50aa9af52666ad3bc0ed7a7f9cb9745af)) - hedon954
- **(readme)** add dependencies section - ([95bed7f](https://github.com/hedon-go-road/template-web/commit/95bed7f64550ba2d5775d7926c09eac4a90d34ec)) - hedon954

<!-- generated by git-cliff -->

---
## [1.0.3](https://github.com/hedon-go-road/template-web/compare/v0.0.1..v1.0.3) - 2024-08-27

publish stable version

### 📚 Documentation

- **(changelog)** delete useless records - ([2a2af4c](https://github.com/hedon-go-road/template-web/commit/2a2af4c9c061f9237e0828c83cfbe41f4b0e4494)) - hedon954
- **(readme)** add go reference badge - ([1c05472](https://github.com/hedon-go-road/template-web/commit/1c05472ad6bf701f9efdc5be0216749e36051775)) - hedon954
- **(readme)** update installation - ([9ab61b9](https://github.com/hedon-go-road/template-web/commit/9ab61b9e765491eb543234832ba5969fc2d03caf)) - Hedon
- **(readme)** update installation - ([2551f4b](https://github.com/hedon-go-road/template-web/commit/2551f4b19cfb5d5e0cb90d2a00457923eedec303)) - hedon954

<!-- generated by git-cliff -->

---
## [0.0.1] - 2024-08-23

Finish go-mysql-mocker basic features and support for go1.20 projects

### ⛰️ Features

- finish gmm - ([ea8f4a0](https://github.com/hedon-go-road/template-web/commit/ea8f4a0fe2cd0320ea3cc78a637338de72b8c4e9)) - wangjiahan
- remove some useless error check and improve unit test for `InitData` - ([280eb23](https://github.com/hedon-go-road/template-web/commit/280eb2396addc2d13fd7acb4164298e507788e08)) - hedon954

### 🐛 Bug Fixes

- **(port)** fix get unused port bug - ([a004a9b](https://github.com/hedon-go-road/template-web/commit/a004a9b6aebe11bf1ad08a03ae69cee5ef1e8bb4)) - hedon954

### 📚 Documentation

- **(gmm)** add some comments for get free port - ([572fb4c](https://github.com/hedon-go-road/template-web/commit/572fb4c5b728c7057ec07b8a11ae2b7e584a5153)) - hedon954
- **(readme)** fix wrong - ([ac3e0fb](https://github.com/hedon-go-road/template-web/commit/ac3e0fb95ff9dc2d348b29e28223ded4f04f42ae)) - hedon954
- **(readme)** add requirements note - ([abfe473](https://github.com/hedon-go-road/template-web/commit/abfe4738e2f8219cd055dc2ea9f9022914ca0d91)) - wangjiahan
- add readme - ([b9cd37a](https://github.com/hedon-go-road/template-web/commit/b9cd37a411ce52151540c6d328908c0d63e07a71)) - hedon954

### 🚜 Refactor

- rename repository to go-mysql-mocker - ([599dbd0](https://github.com/hedon-go-road/template-web/commit/599dbd0eb47cadf0a485076ee78ae1a8e71a1b9c)) - hedon954
- do not use slog to support lower go version user - ([fc235aa](https://github.com/hedon-go-road/template-web/commit/fc235aa48aa10d3f7451c910acabc011b58e7370)) - wangjiahan
- downgrade go-mysql-server to 0.18.0 to support go1.20 projects - ([f337dc8](https://github.com/hedon-go-road/template-web/commit/f337dc8301df31924d1a920f04c8ec2afcf92a8c)) - wangjiahan

<!-- generated by git-cliff -->
