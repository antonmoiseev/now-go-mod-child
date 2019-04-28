# now-go-mod-child

A sample project that demos the issue reported here - https://github.com/zeit/now-builders/issues/431.
Here is the full error log:

```
2019-04-28T09:07:44.799Z  TaskID f71815e8-1acc-41b6-a4ea-a777006d8a4d
2019-04-28T09:07:45.648Z  @zeit/now-dcs-info: using https://dcs.now.systems as the upstream
2019-04-28T09:07:46.113Z  running yarn info for @now/build-utils...
2019-04-28T09:07:46.117Z  running yarn info for @now/go@canary...
2019-04-28T09:07:47.507Z  running yarn for @now/build-utils...
2019-04-28T09:07:47.752Z  yarn add v1.13.0
2019-04-28T09:07:47.773Z  warning package.json: No license field
2019-04-28T09:07:47.783Z  info No lockfile found.
2019-04-28T09:07:47.823Z  warning No license field
2019-04-28T09:07:47.825Z  [1/4] Resolving packages...
2019-04-28T09:07:48.396Z  [2/4] Fetching packages...
2019-04-28T09:07:49.137Z  [3/4] Linking dependencies...
2019-04-28T09:07:49.354Z  [4/4] Building fresh packages...
2019-04-28T09:07:49.367Z  success Saved lockfile.
2019-04-28T09:07:49.369Z  warning No license field
2019-04-28T09:07:49.379Z  success Saved 36 new dependencies.
2019-04-28T09:07:49.379Z  info Direct dependencies
2019-04-28T09:07:49.380Z  └─ @now/build-utils@0.5.0
2019-04-28T09:07:49.380Z  info All dependencies
2019-04-28T09:07:49.381Z  ├─ @now/build-utils@0.5.0
2019-04-28T09:07:49.381Z  ├─ async-retry@1.2.3
                          ├─ async-sema@2.1.4
                          ├─ balanced-match@1.0.0
                          ├─ brace-expansion@1.1.11
                          ├─ buffer-crc32@0.2.13
                          ├─ concat-map@0.0.1
                          ├─ core-util-is@1.0.2
2019-04-28T09:07:49.381Z  ├─ double-ended-queue@2.1.0-0
                          ├─ end-of-stream@1.4.1
                          ├─ errno@0.1.7
                          ├─ from2@2.3.0
                          ├─ fs-extra@7.0.0
2019-04-28T09:07:49.381Z  ├─ fs.realpath@1.0.0
                          ├─ glob@7.1.3
                          ├─ graceful-fs@4.1.15
                          ├─ inflight@1.0.6
                          ├─ inherits@2.0.3
                          ├─ into-stream@5.0.0
2019-04-28T09:07:49.381Z  ├─ isarray@1.0.0
                          ├─ jsonfile@4.0.0
                          ├─ memory-fs@0.4.1
                          ├─ minimatch@3.0.4
                          ├─ multistream@2.1.1
2019-04-28T09:07:49.381Z  ├─ node-fetch@2.2.0
                          ├─ p-is-promise@2.1.0
                          ├─ path-is-absolute@1.0.1
                          ├─ process-nextick-args@2.0.0
                          ├─ prr@1.0.1
                          ├─ readable-stream@2.3.6
2019-04-28T09:07:49.382Z  ├─ retry@0.12.0
                          ├─ safe-buffer@5.1.2
                          ├─ string_decoder@1.1.1
                          ├─ universalify@0.1.2
2019-04-28T09:07:49.382Z  ├─ util-deprecate@1.0.2
                          └─ yazl@2.4.3
2019-04-28T09:07:49.383Z  Done in 1.64s.
2019-04-28T09:07:49.395Z  running yarn for @now/go@canary...
2019-04-28T09:07:49.641Z  yarn add v1.13.0
2019-04-28T09:07:49.662Z  warning package.json: No license field
2019-04-28T09:07:49.673Z  info No lockfile found.
2019-04-28T09:07:49.713Z  warning No license field
2019-04-28T09:07:49.715Z  [1/4] Resolving packages...
2019-04-28T09:07:50.198Z  [2/4] Fetching packages...
2019-04-28T09:07:50.817Z  [3/4] Linking dependencies...
2019-04-28T09:07:51.020Z  [4/4] Building fresh packages...
2019-04-28T09:08:00.479Z  success Saved lockfile.
2019-04-28T09:08:00.482Z  warning No license field
2019-04-28T09:08:00.492Z  success Saved 34 new dependencies.
2019-04-28T09:08:00.492Z  info Direct dependencies
2019-04-28T09:08:00.493Z  └─ @now/go@0.4.1-canary.0
2019-04-28T09:08:00.493Z  info All dependencies
2019-04-28T09:08:00.493Z  ├─ @now/go@0.4.1-canary.0
2019-04-28T09:08:00.493Z  ├─ chownr@1.1.1
                          ├─ cross-spawn@6.0.5
                          ├─ debug@4.1.1
                          ├─ end-of-stream@1.4.1
                          ├─ execa@1.0.0
2019-04-28T09:08:00.494Z  ├─ fs-extra@7.0.1
                          ├─ fs-minipass@1.2.5
                          ├─ get-stream@4.1.0
                          ├─ graceful-fs@4.1.15
                          ├─ is-stream@1.1.0
2019-04-28T09:08:00.494Z  ├─ isexe@2.0.0
                          ├─ jsonfile@4.0.0
                          ├─ minimist@0.0.8
                          ├─ minizlib@1.2.1
                          ├─ mkdirp@0.5.1
2019-04-28T09:08:00.494Z  ├─ ms@2.1.1
                          ├─ nice-try@1.0.5
                          ├─ node-fetch@2.4.1
                          ├─ npm-run-path@2.0.2
                          ├─ once@1.4.0
2019-04-28T09:08:00.494Z  ├─ p-finally@1.0.0
                          ├─ path-key@2.0.1
                          ├─ pump@3.0.0
                          ├─ semver@5.7.0
                          ├─ shebang-command@1.2.0
2019-04-28T09:08:00.494Z  ├─ shebang-regex@1.0.0
                          ├─ signal-exit@3.0.2
                          ├─ strip-eof@1.0.0
                          ├─ tar@4.4.6
2019-04-28T09:08:00.495Z  ├─ universalify@0.1.2
                          ├─ which@1.3.1
                          ├─ wrappy@1.0.2
                          └─ yallist@3.0.3
2019-04-28T09:08:00.496Z  Done in 10.86s.
2019-04-28T09:08:00.662Z  running builder.exports.build...
2019-04-28T09:08:00.663Z  Downloading user files...
2019-04-28T09:08:00.990Z  Parsing AST for "api/handlers/with_child.go"
2019-04-28T09:08:01.005Z  Found exported function "Child" in "api/handlers/with_child.go"
2019-04-28T09:08:01.044Z  Tidy `go.mod` file...
2019-04-28T09:08:01.412Z  go: finding github.com/zeit/now-builders/utils/go/bridge latest
2019-04-28T09:08:06.109Z  go: finding github.com/zeit/now-builders/utils/go latest
2019-04-28T09:08:06.111Z  go: finding github.com/zeit/now-builders/utils latest
2019-04-28T09:08:06.113Z  go: finding github.com/zeit/now-builders latest
2019-04-28T09:08:06.123Z  go: downloading github.com/zeit/now-builders v0.0.0-20190425211131-3373cbca4e67
2019-04-28T09:08:07.496Z  go: extracting github.com/zeit/now-builders v0.0.0-20190425211131-3373cbca4e67
2019-04-28T09:08:07.965Z  go: finding github.com/aws/aws-lambda-go/lambda latest
2019-04-28T09:08:07.965Z  go: finding github.com/aws/aws-lambda-go/events latest
2019-04-28T09:08:08.820Z  go: finding github.com/aws/aws-lambda-go v1.10.0
2019-04-28T09:08:08.828Z  go: downloading github.com/aws/aws-lambda-go v1.10.0
2019-04-28T09:08:08.897Z  go: extracting github.com/aws/aws-lambda-go v1.10.0
2019-04-28T09:08:09.204Z  go: finding github.com/stretchr/testify/assert latest
2019-04-28T09:08:09.962Z  go: finding github.com/stretchr/testify v1.3.0
2019-04-28T09:08:09.968Z  go: downloading github.com/stretchr/testify v1.3.0
2019-04-28T09:08:10.030Z  go: extracting github.com/stretchr/testify v1.3.0
2019-04-28T09:08:10.050Z  go: finding github.com/pmezard/go-difflib v1.0.0
2019-04-28T09:08:10.052Z  go: finding github.com/stretchr/objx v0.1.0
2019-04-28T09:08:10.054Z  go: finding github.com/davecgh/go-spew v1.1.0
2019-04-28T09:08:10.911Z  go: downloading github.com/davecgh/go-spew v1.1.0
2019-04-28T09:08:10.920Z  go: downloading github.com/pmezard/go-difflib v1.0.0
2019-04-28T09:08:10.950Z  go: extracting github.com/pmezard/go-difflib v1.0.0
2019-04-28T09:08:10.971Z  go: extracting github.com/davecgh/go-spew v1.1.0
2019-04-28T09:08:10.985Z  Running `go build`...
2019-04-28T09:08:11.115Z  build command-line-arguments: cannot load now-go-mod-child/handlers/child: cannot find module providing package now-go-mod-child/handlers/child
2019-04-28T09:08:11.117Z  failed to `go build`
2019-04-28T09:08:11.120Z  { Error: Command failed: go build -o /tmp/3bf95f77/handler /tmp/5f2020ac/src/lambda/api/handlers/main__mod__.go
                              at makeError (/tmp/utils/build-module/node_modules/execa/index.js:174:9)
                              at Promise.all.then.arr (/tmp/utils/build-module/node_modules/execa/index.js:278:16)
                              at <anonymous>
                              at process._tickDomainCallback (internal/process/next_tick.js:228:7)
                            code: 1,
                            stdout: null,
                            stderr: null,
                            failed: true,
                            signal: null,
                            cmd: 'go build -o /tmp/3bf95f77/handler /tmp/5f2020ac/src/lambda/api/handlers/main__mod__.go',
                            timedOut: false,
                            killed: false }
```
