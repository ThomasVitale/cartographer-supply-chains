# Verifying the Knative Serving Package Release

This package is published as an OCI artifact, signed with Sigstore [Cosign](https://docs.sigstore.dev/cosign/overview), and associated with a [SLSA Provenance](https://slsa.dev/provenance) attestation.

Using `cosign`, you can display the supply chain security related artifacts for the `ghcr.io/kadras-io/cartographer-supply-chains` images. Use the specific digest you'd like to verify.

```shell
cosign tree ghcr.io/kadras-io/cartographer-supply-chains
```

The result:

```shell
📦 Supply Chain Security Related artifacts for an image: ghcr.io/kadras-io/cartographer-supply-chains
└── 💾 Attestations for an image tag: ghcr.io/kadras-io/cartographer-supply-chains:sha256-75d932ea4d326c7e104ae1403b2999d964f71282c11d521b98e06094dd0317f3.att
   └── 🍒 sha256:e3ace2d7101a4af2f93921a6e939081f5eb984e73f00b96afbb9ddde3c535b58
└── 🔐 Signatures for an image tag: ghcr.io/kadras-io/cartographer-supply-chains:sha256-75d932ea4d326c7e104ae1403b2999d964f71282c11d521b98e06094dd0317f3.sig
   └── 🍒 sha256:f25ac25a6a21cd4ade2753e39a4430ca48159da9fd789e5ba175ce779faaf88e
```

You can verify the signature and its claims:

```shell
cosign verify \
   --certificate-identity-regexp https://github.com/kadras-io \
   --certificate-oidc-issuer https://token.actions.githubusercontent.com \
   ghcr.io/kadras-io/cartographer-supply-chains | jq
```

You can also verify the SLSA Provenance attestation associated with the image.

```shell
cosign verify-attestation --type slsaprovenance \
   --certificate-identity-regexp https://github.com/slsa-framework \
   --certificate-oidc-issuer https://token.actions.githubusercontent.com \
   ghcr.io/kadras-io/cartographer-supply-chains | jq .payload -r | base64 --decode | jq
```
