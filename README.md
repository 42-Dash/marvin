## TODO

#### Flushing does not work for the loggers.

#### The output of commands like `git clone` and `git push` should be logged instead of printed to the console.

# Dash

This repository contains a script to manage GitHub repositories for a competition. The script allows you to create repositories, add collaborators with specific permissions, and restrict access to repositories once the competition is over.

## Usage

You can use the script in three ways:

1. **Restrict permissions**:
   ```sh
   go run main.go restrict
   ```
2. **Delete repositories**:
   ```sh
   go run main.go delete
   ```
3. Create repositories and add collaborators:
   ```sh
   go run main.go create
   ```

Each option allows you to either set the permissions of the participants to read-only, delete all the repositories, or create all the repositories and add collaborators with permission to push.

## Prerequisites

1. **.env file**: You should have a `.env` file with the following properties:
   - `GITHUB_ACCESS`: Your GitHub access token with permissions to delete, create, and set collaborator permissions.
   - `GITHUB_ORGANISATION`: The name of your GitHub organization, which you will use as the turn-in organization for repositories.

2. **participants.json file**: You should have a participants.json file formatted in the following way:
   ```json
    {
        "teams": [
            {
                "name": "The-Avengers",
                "members": ["Ironman", "Thor"]
            },
            {
                "name": "Pirates-of-the-Caribbean",
                "members": ["CaptainSparrow", "barbossa007"]
            },
            {
                "name": "The-Justice-League",
                "members": ["Superman3310", "tatman_is_here"]
            }
        ]
    }
   ```

### Properties in participants.json
   - `teams`: An array of team objects.
      - `name`: The name of the team.
      - `members`: An array of GitHub usernames of the team members.

### Example

Here is an example of how to use the programm:
1. Create the `.env` file with your GitHub access token and organization name:
```sh
echo 'GITHUB_ACCESS=<your_github_access_token>' > .env
echo 'GITHUB_ORGANISATION=<your_github_organisation>' >> .env
```

2. Create the `participants.json` file with the team and member information:
```json
{
    "teams": [
        {
            "name": "The-Avengers",
            "members": ["Ironman", "Thor"]
        },
        {
            "name": "Pirates-of-the-Caribbean",
            "members": ["CaptainSparrow", "barbossa007"]
        },
        {
            "name": "The-Justice-League",
            "members": ["Superman3310", "tatman_is_here"]
        }
    ]
}
```

3. Run the script to create repositories and add collaborators:
```sh
go run main.go create
```

4. Run the script to restrict permissions to read-only once the competition is over:
```sh
go run main.go restrict
```

5. Run the script to delete all repositories:
```sh
go run main.go delete
```
