# Understanding Setup Scripts

```
        k3d cluster create local                                                                                \
            --registry-use k3d-registry.localhost:5000                                                          \
            --k3s-server-arg '--disable=traefik'                                                                \
            --k3s-server-arg '--disable=metrics-server'                                                         \
            --k3s-server-arg '--kubelet-arg=eviction-hard=imagefs.available<3%,nodefs.available<3%'             \
            --k3s-server-arg '--kubelet-arg=eviction-minimum-reclaim=imagefs.available=1%,nodefs.available=1%'  \
            --port 27017:27017@loadbalancer                                                                     \
            --port 5432:5432@loadbalancer                                                                       \
            --port 80:80@loadbalancer                                                                           \
            --volume "${HOME}/.aws:/mnt/.aws"                                                                   \
            --volume "${DIR}/airflow:/mnt/airflow"                                                              \
            --volume "${DIR}/cathedral:/mnt/cathedral"                                                          \
            --volume "${DIR}/test-bot:/mnt/test-bot"                                                            \
            --volume "${DIR}/ice:/mnt/ice"                                                                      \
            --volume "${DIR}/mist:/mnt/mist"                                                                    \
            --volume "${DIR}/frazil:/mnt/frazil"                                                                \
            --volume "${DIR}/styleguide:/mnt/styleguide"                                                        \
            --volume "${DIR}/local/mongo/data/volume:/mnt/mongo"                                                \
            --volume "${DIR}/local/postgres/data/volume:/mnt/postgres"                                          \
            --volume "${DIR}/config/k3d/d-coredns-patch.yaml:/var/lib/rancher/k3s/server/manifests/d-coredns-patch.yaml
```