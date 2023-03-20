#!/bin/bash

set -eu
# don't set any options, otherwise it will fail arbitrarily
# set -euo pipefail

work_dir=$1
last_commit_tag=$2
repo_name=$3
branch=$4
origin_repo_url=$5
remote_repo_url=$6

set +e
test -d $work_dir || mkdir -p $work_dir
set -e

cd $work_dir

if [ -d $repo_name ]; then
    cd $repo_name

    git fetch origin
else
    git clone -q $origin_repo_url

    cd $repo_name
fi

set +e
git rev-parse --verify $branch 2>/dev/null
has=$?
set -e

if [ $has -ne 0 ]; then
    git checkout -b $branch
fi

git merge origin/$branch

git push -f $remote_repo_url $branch

last_commit=$(git log --format="%H" -n 1)

echo "${last_commit_tag}${last_commit}"
