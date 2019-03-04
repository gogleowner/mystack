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

    ```
8b47041 (HEAD -> master, origin/master, origin/HEAD) add docker commaned ; stop or remove all container
dd852df add docker cp, resource constraint
f32cafa add nginx trouble shooting
7ea6701 add reference link about root user id
    ```

    - `$ git rebase -i -p 7ea6701` 어느 커밋 시점부터 author을 변경할지 선택
    - author를 변경할 커밋을 `edit` 으로 변경

        ```
edit f32cafa add nginx trouble shooting
edit dd852df add docker cp, resource constraint
edit 8b47041 add docker commaned ; stop or remove all container
        ```

    - 해당 커밋에 대한 author 를 수정하여 커밋, rebase를 진행한다.
        - `$ git commit --amend --author="gogleowner <gogleowner@gmail.com>" --no-edit`
        - `$ git rebase --continue`


