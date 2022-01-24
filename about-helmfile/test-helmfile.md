# The test-helmfile cli

$ test-helmfile
Usage: test-helmfile [OPTIONS]

  Wrapper around helmfile commands to support debugging.

Options:
  --app TEXT                  Which app should we use during rendering.  [required]
  --branch TEXT               Which branch should we use during rendering.  [required]
  --env TEXT                  Which env should we use during rendering.  [required]
  --tag TEXT                  Which tag should we use during rendering.
  --debug                     Pass to enable helm debug flag.
  --temp_dir TEXT             Where to save rendered templates.
  --helmfile TEXT             Which helmfile to use for rendering.
  --build / --no-build        Should we run helmfile build?
  --write / --no-write        Should we run helmfile write-values?
  --template / --no-template  Should we run helmfile template?
  --kubeval / --no-kubeval    Should we run kubeval?
  --help                      Show this message and exit.

  # The test-helmfile cli