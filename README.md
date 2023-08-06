# NCGLPM
Golang binary to consolidate helper tools for managing Proxmox. Also intended as a mechanism to learn Golang and CICD (probably Github Actions).



### Release

1. Update the `CHANGELOG.md` file - this should be human readable text.
2. Upversion the `version` value in the `root.go`
3. Merge release branch into main
4. Tag main with `release/<version>` to trigger the release CICD.