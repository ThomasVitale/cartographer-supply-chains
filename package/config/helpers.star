load("@ytt:data", "data")
load("@ytt:assert", "assert")

def config_writer():
  if data.values.gitops.strategy == "pull_request":
    return "tekton-config-writer-and-pull-requester-template"
  end
  return "tekton-config-writer-template"
end
