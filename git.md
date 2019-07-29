# git commands

## move to a fork remote

1. Fork repository to your account on github
2. add new remote to your repository

    ```
$ git remote add upstream git@github.com/{your_account}/{your_repository}
    ```

3. Fetch & Push
        
    ```
// fetch upstream
$ git fetch upstream
// push upstream
$ git push upstream
    ```

## git config
- `$ git config --list` : print global, local repository settings
- `$ git config --edit` : open editor to edit local repository config
- `$ git config --unset {some config}` : unset config in local repository config
- `$ git config --global --unset {some config}` : unset config in global repository config

## Editing the Author of Past Commits
- Change latest commit using `--amend`
    - `$ git commit --amend --author="gogleowner <gogleowner@gmail.com>"`
- Change previous commits using `rebase`
    - `$ git rebase -i -p 7ea6701` 어느 커밋 시점부터 author을 변경할지 선택
    - author를 변경할 커밋을 `edit` 으로 변경
    - 해당 커밋에 대한 `author` 를 수정하여 커밋, rebase를 진행한다.
        - `$ git commit --amend --author="gogleowner <gogleowner@gmail.com>" --no-edit`
        - `$ git rebase --continue`

## git log
- `$ git log --stat` : 커밋 이력에 변경된 파일 목록 확인

## git stash
- `$ git stash clear` : 모든 stash 상태 제거
  - https://mirrors.edge.kernel.org/pub/software/scm/git/docs/git-stash.html

# reference
- https://www.ocpsoft.org/tutorials/git/reset-and-sync-local-respository-with-remote-branch/

