# Gotcha 1

If you push the backend before specify the `routes` in the manifest, it will create the default one, and not remove it when you later add the explicit `routes`. Deleting the app and re-pushing seemed to fix this.

# Gotcha 2

If you don't set a policy, you'll get `Error talking to backend: Get http://backend-cni-test.apps.internal:8080: dial tcp 10.x: connect: connection refused`

Fix by adding a policy:

```bash
cf add-network-policy frontend-cni-test --destination-app backend-cni-test
```

# Gotcha 3

Note the policy maps apps by GUIDs, so it won't survive a `zero-downtime-push`.

Watch this space.
