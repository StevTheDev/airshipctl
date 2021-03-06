## airshipctl cluster rotate-sa-token

Rotate tokens of Service Accounts

### Synopsis

Use to reset/rotate the Service Account(SA) tokens and additionally restart the
corresponding pods to get the latest token data reflected in the pod spec

Secret-namespace is a mandatory field and secret-name is optional. If secret-
name is not given, all the SA tokens in that particular namespace is considered,
else only that particular input secret-name

```
airshipctl cluster rotate-sa-token [flags]
```

### Examples

```

# To rotate a particular SA token
airshipctl cluster rotate-sa-token -n cert-manager -s cert-manager-token-vvn9p

# To rotate all the SA tokens in cert-manager namespace
airshipctl cluster rotate-sa-token -n cert-manager

```

### Options

```
  -h, --help                      help for rotate-sa-token
      --kubeconfig string         Path to kubeconfig associated with cluster being managed
  -s, --secret-name string        name of the secret containing Service Account Token
  -n, --secret-namespace string   namespace of the Service Account Token
```

### Options inherited from parent commands

```
      --airshipconf string   Path to file for airshipctl configuration. (default "$HOME/.airship/config")
      --debug                enable verbose output
```

### SEE ALSO

* [airshipctl cluster](airshipctl_cluster.md)	 - Manage Kubernetes clusters

