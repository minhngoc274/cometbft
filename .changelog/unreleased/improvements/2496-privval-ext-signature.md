- `[privval]` DO NOT require extension signature from privval if vote extension
  is disabled [\#2496](https://github.com/cometbft/cometbft/pull/2496)

  Remote signers must ONLY sign the extension if `sign_extension` flag in
  `SignVoteRequest` is true.
