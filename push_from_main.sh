#!/usr/bin/bash

# This is a shell script created to push from the local repo to the remote repo (aka GitHub repo)
# easier

git checkout main # Checkout 'main'
git merge feature main # Merge 'main' with 'feature'
git push # Push from 'main' to remote
git push origin feature # Ditto with 'feature'
