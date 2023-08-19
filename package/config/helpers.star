load("@ytt:assert", "assert")
load("@ytt:data", "data")

def config_writer():
  if data.values.gitops.strategy == "pull_request":
    return "tekton-write-config-and-pr-template"
  end
  return "tekton-write-config-template"
end
