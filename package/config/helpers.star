load("@ytt:data", "data")
load("@ytt:assert", "assert")

def config_writer():
  if data.values.gitops.strategy == "pull_request":
    return "tekton-write-config-and-pr-template"
  end
  return "tekton-write-config-template"
end
