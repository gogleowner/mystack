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
