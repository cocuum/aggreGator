# Gator

We're going to build an RSS feed aggregator in Go! We'll call it "Gator", you know, because aggreGATOR 🐊. Anyhow, it's a CLI tool that allows users to:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Prerequisites

This program requires the following to run successfully:

- GO
- SQL (Postgres)

## Install

This CLI program can be installed using go install

## Configuration

This program use .gatorcongif.json to db and current user (logged in) information

Here is a list of command usable in the CLI:

- register
- login
- reset
- users
- agg
- addfeed
- feeds
- follow
- following
- unfollow
- browse
