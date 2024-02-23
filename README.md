## Table of Contents
- [Table of Contents](#table-of-contents)
- [Current Scope](#current-scope)
- [Project Status](#project-status)
- [Installation and Setup Instructions](#installation-and-setup-instructions)
- [Requirements](#requirements)
  - [Getting The Codebase:](#getting-the-codebase)
  - [Installation and Configuration:](#installation-and-configuration)
- [License](#license)

## Current Scope

A basic API collection for Habit Tracker App implementing JWT Auth & CRUD operations having the following functionalities:
* `Register`, `Login`, & Get User Info.
* User can Add new `Habit`.
* Update an existing `Habit`.
* Delete an existing `Habit`.
* View `Habit` information.
* List all `Habit`.

## Project Status

This project is currently under development. Users now can do the above functionalities, As per requirements of the assignment.

## Installation and Setup Instructions

You need the following requirements installed globally on your machine.

## Requirements
- `GoLang` >= v1.22
- `PostgreSQL` >= v16.2
- Proxy Web Server: `Apache2` or preferably `NGINX` in case of creating a virtual host instead of specifying the port in URL rather than the default port `80`.

<b><span style="background-color:yellow; color:black">Note:</span></b> The requirements listed above are needed in case you'll ignore the docker setup. Otherwise, if you find it convenient for you to work with `Docker` example files are available in the root directory of the project for convenience.

### Getting The Codebase:

The simplest way to obtain the code is using the github .zip feature. Click [here](https://github.com/ahmadSaeedGoda/habit_tracking_app/archive/refs/heads/master.zip) to get the latest stable version as a .zip compressed file.

The recommended way is using `GIT`. You'll need to make sure `git version ~2.34.1` is installed on your machine. Use a terminal or Power Shell to visit the directory where you'd like to have the source code placed, then type in:
```sh
$ git clone https://github.com/ahmadSaeedGoda/habit_tracking_app.git
```
Feel free to switch the URL to use `SSH` protocol instead!

### Installation and Configuration:
- <b>Step 1:</b> Get the code. "As explained [above](#getting-the-codebase)".

- <b>Step 2:</b> Set the Environment Variables. Find the file named `.env.habit_tracker.example` in the root directory of the project. Duplicate the file in the same path/location, then rename the new one `.env.habit_tracker` then set the values of the environment variables listed within the file according to your environment respectively.

Replicate the same step for `docker-compose.yml.example`, `Dockerfile.dev.example`, `Dockerfile.postgres.example`. Just remove the suffix `.examples` after duplicating these files and you're good to go!

Same thing/procedures can be conducted for `.air.toml.example` to be used for configuring the Dev server during development. Consult the [Air's documentation]([URL](https://github.com/cosmtrek/air)) for more details/instructions in this regard.

- <b>Step 3:</b> Navigate to the root directory of the project you cloned or downloaded via CLI, then run the following command to get it up and running!
```sh
$ docker compose up
```
Easy peasy lemon squeasy, my friend! That's how we roll.<br>
Voila! Just like that, but with more pizzazz.

<b><span style="background-color:yellow; color:black">Note:</span></b> `docker compose up` in the console after making sure the prompt points to the root directory of the project. So that you can have an up & running <strong><span style="background-color:white; color:blue">dev server</strong></span> on the default port 8080 with `Air` for Auto-Restart/Hot-Reload whenever you make any modifications to the code and save the files.

><span style="color:lightgreen">Specify the file named `Dockerfile.staging.uat.example` after duplicating it & removal of `.example` suffix in your `docker-compose.yml`.
This way, when running the file in the console after making sure the prompt points to the root directory of the project you can have  a production similar running server.
</span>

<br>

Whatever convenient for you to run the app for visiting the endpoints in an API client such as `postman`, `insomnia` or even `curl`.

<br><b><span style="background-color:yellow; color:black">Note:</span></b> A shared Postman collection is included/shared within the source code root directory, this can be imported and ready to use after changing the `base_url` variable as per your env. "In case you'd like to change the default"

For documentation & Usage see `Postman Collection` shared collection with Examples included in file: `Insomnia_2024-02-22.json`.

Import the above file in `Insomnia` or `Postman`!

Now you can Register & Login, View your info, Create new `Habit`, then once a new one is created via the available request/endpoint .. Update an existing one, Deleting an existing one, Get a specified `Habit`, or Get all `Habits` for the authenticated user can be achieved.

## License
This is a free software distributed under the terms of the WTFPL license along with MIT license as dual-licensed, You can choose whatever works for you.<br/><br/>
Review the attached License file within the source code for mor details.
