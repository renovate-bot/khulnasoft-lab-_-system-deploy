package builtin

import (
	// Import all built-in actions
	_ "github.com/khulnasoft-lab/system-deploy/pkg/actions/builtin/copy"
	_ "github.com/khulnasoft-lab/system-deploy/pkg/actions/builtin/editfile"
	_ "github.com/khulnasoft-lab/system-deploy/pkg/actions/builtin/exec"
	_ "github.com/khulnasoft-lab/system-deploy/pkg/actions/builtin/onchange"
	_ "github.com/khulnasoft-lab/system-deploy/pkg/actions/builtin/platform"
	_ "github.com/khulnasoft-lab/system-deploy/pkg/actions/builtin/systemd"
)
