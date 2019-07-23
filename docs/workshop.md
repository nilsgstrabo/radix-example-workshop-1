 <!-- markdownlint-disable MD014 MD007 MD034-->

# 1. Omnia Radix Workshop 1

Purpose

The purpose of the workshop is to give a general and hands-on introduction to Radix. We do this by "forking" and example application, get familiar with the application, establish a CI environment locally, move the application into Radix and then establish the full CI/CD DevOps cycle.

<!-- TOC -->

- [1. Omnia Radix Workshop 1](#1-omnia-radix-workshop-1)
    - [1.1. Part 1](#11-part-1)
        - [1.1.1. Pre-requisites](#111-pre-requisites)
        - [1.1.2. Getting started](#112-getting-started)
        - [1.1.3. Exploring the Echo app](#113-exploring-the-echo-app)
        - [1.1.4. Exploring the WWW app](#114-exploring-the-www-app)
        - [1.1.5. Preparing for Radix](#115-preparing-for-radix)
        - [1.1.6. Explore radixconfig.yaml](#116-explore-radixconfigyaml)
        - [1.1.7. Creating the application on Radix](#117-creating-the-application-on-radix)
        - [1.1.8. Using multiple branches - multiple environments](#118-using-multiple-branches---multiple-environments)
        - [1.1.9. Monitoring & Metrics](#119-monitoring--metrics)
    - [1.2. Typical questions](#12-typical-questions)
    - [1.3. Where to get started, get help, log issues or feature requests](#13-where-to-get-started-get-help-log-issues-or-feature-requests)
        - [1.3.1. Getting help](#131-getting-help)
        - [1.3.2. Getting started](#132-getting-started)
        - [1.3.3. Log issues & feature requests](#133-log-issues--feature-requests)
    - [1.4. Part 2](#14-part-2)
        - [1.4.1. Next steps](#141-next-steps)

<!-- /TOC -->

## 1.1. Part 1

### 1.1.1. Pre-requisites

- Have access to the Radix Playground
  - Apply for access to the Radix Playground in AccessIT (seach for ````Radix Playground Users````)
  -  Find Radix at [https://console.playground.radix.equinor.com](https://console.playground.radix.equinor.com)
- Account on github.com
- Git installed and working locally against github.com
- Docker running locally
- Local dev. environment (IDE++)
- Node js eco system installed and running ([Downnload Nodejs](https://nodejs.org/en/download/))
- Laptop that "works" on “Statoil Approved” WiFi. If not - know how to handle proxy fun for both Docker and local development environment

### 1.1.2. Getting started

1. Fork repository to your home on github. Consider choosing an alternative name for the repository
2. Clone your newly forked repository down to your developer laptop

### 1.1.3. Exploring the Echo app

1. Move into the [echo](../echo/) folder and explore how to develop the Echo app using ```Node.js``` as well as Dockerizing the application.

### 1.1.4. Exploring the WWW app

1. Move into the [www](../www/) folder and explore how to develop the WWW app using ```Node.js``` as well as Dockerizing the application.

Remember that the Echo app needs to run somewhere to get proper response.

### 1.1.5. Preparing for Radix

- The Radix cluster we use for the workshop is available at https://console.playground.radix.equinor.com/
- Radix documentation is available at https://www.radix.equinor.com/

Important to know:

1. The difference between ```platform user``` and ```application user```
2. Important terminology: ```application```, ```environments```,```components```, and ```replicas``` [Important Radix Concepts](https://www.radix.equinor.com/docs/topic-concepts/)
3. ```radixconfig.yaml``` - lives on the master branch and is your infrastrucure as code - drive your app in Radix.

### 1.1.6. Explore radixconfig.yaml

1. Reading the [docs](https://github.com/equinor/radix-operator/blob/master/docs/radixconfig.md) [docs app](https://www.radix.equinor.com/docs/reference-radix-config/)
2. Exploring the config file for the example app [./radixconfig.yaml](../radixconfig.yaml)

### 1.1.7. Creating the application on Radix

1. Update the name of ```your instance``` of the application in radixconfig.yaml
2. Follow the getting started guide (www.radix.equinor.com) or "just do it!"
3. Do a change to trigger the initial build (or use the "New job" feature in the jobs/environment section). Examine web-hooks and reponse in Radix
4. Verify that the app work on the public end-point it has been given.

### 1.1.8. Using multiple branches - multiple environments

Radix support connecting a branch to a specific environment. Let's explore this.

1. Update the radixconfig.yaml file in ```master```, commit, push and explore what's happening in Radix. (Copy the file ./radixconfigs/radixconfig-feature1.yaml to ./radixconfig.yaml. Remember to update app name)
2. Check out the "new" branch (feature1)
3. Examine code - the new feature  (getting a new env variable from Echo (node_env))
    - Alternatively:
    - Checkout a new branch
    - Update the tests for Echo
    - Run tests - which fails
    - Add feature to Echo
    - Run tests - hopefully with success
    - Build Docker image and verify
    - Update views for echo.ejs in WWW
    - Run npm start and verify
    - Commit changes and push
4. (Do a change, commit, push and) explore what's happening in Radix.

### 1.1.9. Monitoring & Metrics

The Echo component is exposing metrics on the /metrics endpoint. These metrics are scraped by [Prometheus](https://prometheus.io/docs/introduction/overview/) and made available in [Grafana](https://grafana.com/). Consult the docs for Prometheus and Grafana for how to work with metrics and monitoring.

## 1.2. Typical questions

(Status as of January 2019)

- Storage - databases
- Authentication
- Logging
- Metrics - Monitoring
- Radix CLI (Api)
- Backup & Disaster recovery
- Own domain names / urls for apps

## 1.3. Where to get started, get help, log issues or feature requests

### 1.3.1. Getting help

- Radix on #Slack ([#omnia_radix](https://equinor.slack.com/messages/C8U7XGGAJ), [#omnia_radix_support](https://equinor.slack.com/messages/CBKM6N2JY))
- [Raidx front app](https://www.radix.equinor.com/)
- Radix on [github.com](https://github.com/equinor/radix-platform)

### 1.3.2. Getting started

- Radix Getting Started - https://www.radix.equinor.com/

### 1.3.3. Log issues & feature requests

- https://github.com/equinor/radix-platform/issues

It makes sense to examing existing issues and perhaps discuss on Slack prior to logging a new one

## 1.4. Part 2

### 1.4.1. Next steps

- Move your own apps into Radix
