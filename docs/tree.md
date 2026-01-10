## Repository Tree

 ```console
.
├── .gitignore
├── .goreleaser.yaml
├── LICENSE
├── Makefile
├── NOTICE
├── README.md
├── TODO.md
├── cmd
│   └── spike
│       └── main.go
├── docs
│   ├── ROADMAP.md
│   └── tree.md
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   └── spike
│   │       └── cmd
│   │           ├── banner.go
│   │           ├── root.go
│   │           ├── run.go
│   │           ├── scan.go
│   │           ├── scan_db.go
│   │           ├── shell.go
│   │           └── utils.go
│   └── pkg
│       ├── reporter
│       │   ├── generator.go
│       │   ├── notifier.go
│       │   ├── telegram.go
│       │   └── telegram_test.go
│       ├── scanner
│       │   ├── cli
│       │   │   ├── cachex.go
│       │   │   ├── checker.go
│       │   │   ├── cli.go
│       │   │   ├── errors.go
│       │   │   ├── exec_test.go
│       │   │   ├── gau.go
│       │   │   ├── httpx.go
│       │   │   ├── katana.go
│       │   │   ├── nuclei.go
│       │   │   ├── nuclei_test.go
│       │   │   ├── subfinder.go
│       │   │   ├── subfinder_test.go
│       │   │   ├── uro.go
│       │   │   └── utils.go
│       │   ├── core
│       │   │   ├── cachex_step.go
│       │   │   ├── core.go
│       │   │   ├── executor.go
│       │   │   ├── gau_step.go
│       │   │   ├── httpx_step.go
│       │   │   ├── katana_step.go
│       │   │   ├── nuclei_dast_step.go
│       │   │   ├── nuclei_generic_step.go
│       │   │   ├── pipeline.go
│       │   │   ├── subfinder_step.go
│       │   │   ├── types.go
│       │   │   ├── urlsutils.go
│       │   │   ├── uro_step.go
│       │   │   └── utils.go
│       │   └── db
│       │       ├── chunk.go
│       │       ├── db.go
│       │       ├── domain_repository.go
│       │       ├── repository.go
│       │       ├── scan_tracker_repo.go
│       │       └── tool_repositories.go
│       └── shell
│           ├── banner.go
│           ├── context.go
│           ├── handlers.go
│           ├── shell.go
│           ├── shell_test.go
│           └── ui.go
└── pkg
    ├── config
    │   ├── config.go
    │   ├── default.go
    │   ├── loader.go
    │   └── loader_test.go
    ├── logger
    │   ├── colors.go
    │   ├── logger.go
    │   └── logger_test.go
    ├── spike
    │   ├── errors.go
    │   ├── options.go
    │   ├── spike.go
    │   └── spike_test.go
    └── version
        ├── version.go
        └── version_test.go

20 directories, 76 files

```
