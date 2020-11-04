# Hikeshi

> Hikeshi is a security incident response application that keeps documenting incidents simple, so you can focus on fighting fires. 

## Table of Contents

- [About Hikeshi](#About-Hikeshi)
- [Pre-Requisites](#Requirements)
- [Database Setup](#Database-Setup)
    * [Database Migrations](#Create-and-Populate-Your-Databases)
- [Starting the application](#Starting-the-Application)
- [Editing the source](#Editing-the-source)

---

![Hikeshi Gif](/assets/images/hikeshi.gif)


## About Hikeshi

The name comes from the firefighting system established in Japan during the Edo period. Incident response involves a lot of "firefighting" in the sense that you're only called upon when something has gone very wrong. In high stress situations, it's nice to have a simple tool that assists you while you fight those fires.

- No-clutter UI. 
- Dark mode support by default.
- Designed specifically for incident response.

It isn't feature packed, but it does exactly what it needs to while you handle the incidents at hand.


## Requirements

1. `brew install buffalo`
2. `brew install postgres`

## Database Setup

> Hikeshi uses a Postgres database to store its information. 

1. Start your postgres database
1. Edit the `database.yml` file to use the correct usernames, passwords, and hosts 

### Create and Populate Your Databases

> Ok, so you've edited the "database.yml" file and started your database, now Buffalo can create the databases in that file for you:

```shell
buffalo db create -a
buffalo task db:seed:users
buffalo task db:seed:incidents
```

## Starting the Application

> Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

```shell
buffalo dev
```

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a Login page.

You now have your Hikeshi application up and running.

## Editing the source

If you'd like to make some additions to the project, go checkout [gobuffalo.io](http://gobuffalo.io) to get started with the Buffalo framework.

[Powered by Buffalo](http://gobuffalo.io)
