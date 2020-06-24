# Github Enterprise token should be set in shell via export GITHUB_TOKEN.
ifeq ($(wildcard .BuildHarnessConfig.global),)
    ifndef GITHUB_TOKEN
        $(error GITHUB_TOKEN is not set in environment variables)
    endif
endif

# Include global config from CICD repo.
ifeq ($(wildcard .BuildHarnessConfig.global),)
    -include $(shell curl -sSL -H "Authorization: token ${GITHUB_TOKEN}" -o .BuildHarnessConfig.global "https://raw.github.ibm.com/cp4mcm-devops/configurations/master/build-harness-config/BuildHarnessConfig.global"; echo .BuildHarnessConfig.global)
else
    -include .BuildHarnessConfig.global
endif

# Bootstrap the build harness and extensions.
-include $(shell $(BUILD_HARNESS_BOOTSTRAP); echo .build-harness-bootstrap)

# Include local config and overrides.
-include BuildHarnessConfig.local
